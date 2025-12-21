package crawler

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"maplestory-llm-docs/internal/logger"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

// StartURL is the default initial page to begin crawling.
const StartURL = "https://maplestoryworlds-creators.nexon.com/ko/docs/?postId=472"

const (
	navContainerSel = "#App > main > div.contents_wrap > div.tree_view_container"
	// contentOuterSel targets the container div used when opening curURL in a separate context
	// per requirement: search for div elements with class text_content_container
	contentOuterSel     = "div.text_content_container"
	contentContainerSel = "#App > main > div.contents_wrap > div.renderContent > div.text_content_container > div.text_content"
	titleSel            = "#App > main > div.contents_wrap > div.renderContent h1"
)

// Crawler는 실행마다 반복 전달하던 설정값을 보관해 재사용하기 위한 구조체입니다.
type Crawler struct {
	ClickDelay     time.Duration
	Limit          int
	OverallTimeout time.Duration
}

// NewCrawler는 주어진 설정으로 Crawler 인스턴스를 생성합니다.
func NewCrawler(clickDelay time.Duration, limit int, overallTimeout time.Duration) *Crawler {
	if limit < 0 {
		limit = 0
	}
	if clickDelay < 0 {
		clickDelay = 0
	}
	if overallTimeout < 0 {
		overallTimeout = 0
	}
	return &Crawler{
		ClickDelay:     clickDelay,
		Limit:          limit,
		OverallTimeout: overallTimeout,
	}
}

// Run은 Crawler에 저장된 설정을 사용하여 크롤링을 수행하고 결과를 outPath에 저장합니다.
func (c *Crawler) Run(headless bool, outPath, format string, startURL string) error {
	allocOpts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", headless),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0 Safari/537.36"),
	)

	ctx, cancelAlloc := chromedp.NewExecAllocator(context.Background(), allocOpts...)
	defer cancelAlloc()

	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	if c.OverallTimeout > 0 {
		var toCancel context.CancelFunc
		ctx, toCancel = context.WithTimeout(ctx, c.OverallTimeout)
		defer toCancel()
	}

	// Navigate to start URL
	if err := chromedp.Run(ctx, chromedp.Navigate(startURL)); err != nil {
		return err
	}
	if err := waitVisible(ctx, navContainerSel, 30*time.Second); err != nil {
		return fmt.Errorf("navigation container not visible: %w", err)
	}

	visited := make(map[string]bool)
	var docs []Document
	backoff := NewBackoff(500*time.Millisecond, 20*time.Second, 2.0, 0.2)

	// 1) 확장 단계: 하위 항목을 가진 모든 닫힌 요소를 더 이상 없을 때까지 클릭
	for {
		_ = scrollMenuToEnd(ctx)

		var nodes []*cdp.Node
		if err := chromedp.Run(ctx, chromedp.Nodes(navContainerSel+" *", &nodes, chromedp.ByQueryAll)); err != nil {
			return fmt.Errorf("query nodes: %w", err)
		}

		expanded := false
		for _, n := range nodes {
			// 조건: span.inactiveDot.isHavingChildren 이고 not .isHavingChildrenAndOpen
			if strings.EqualFold(n.LocalName, "span") && hasAllClasses(n, "inactiveDot", "isHavingChildren") && !hasClass(n, "isHavingChildrenAndOpen") {
				_ = chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
					return dom.ScrollIntoViewIfNeeded().WithNodeID(n.NodeID).Do(ctx)
				}))
				if err := chromedp.Run(ctx, chromedp.MouseClickNode(n)); err != nil {
					continue
				}
				time.Sleep(c.ClickDelay)
				expanded = true
			}
		}
		if !expanded {
			break
		}
	}

	// 2) 수집 단계: div.inactiveDepth 중 텍스트가 한 글자라도 있고, .activeParent 가 아닌 요소만 대상으로 클릭/수집
	_ = scrollMenuToEnd(ctx)
	var leafNodes []*cdp.Node
	if err := chromedp.Run(ctx, chromedp.Nodes(navContainerSel+" div.inactiveDepth", &leafNodes, chromedp.ByQueryAll)); err != nil {
		return fmt.Errorf("query leaf nodes: %w", err)
	}

	for _, n := range leafNodes {
		if hasClass(n, "activeParent") {
			continue
		}
		// 텍스트 존재 여부 확인: outerHTML에서 태그를 제외한 가시 텍스트가 있는지 검사
		hasText := false
		_ = chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
			html, err := dom.GetOuterHTML().WithNodeID(n.NodeID).Do(ctx)
			if err != nil {
				return nil
			}
			if hasAnyTextInHTML(html) {
				hasText = true
			}
			return nil
		}))
		if !hasText {
			continue
		}

		// 클릭
		_ = chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
			return dom.ScrollIntoViewIfNeeded().WithNodeID(n.NodeID).Do(ctx)
		}))
		if err := chromedp.Run(ctx, chromedp.MouseClickNode(n)); err != nil {
			continue
		}
		time.Sleep(c.ClickDelay)

		// 문서 URL 판별
		var curURL string
		_ = chromedp.Run(ctx, chromedp.Location(&curURL))
		postID, ok := ExtractPostIDFromURL(curURL)
		if !ok || visited[postID] {
			continue
		}

		// 내용 수집 (백오프 포함)
		var title string
		err := withRetry(backoff, 5, func() error {
			if err := waitVisible(ctx, contentContainerSel, 30*time.Second); err != nil {
				return err
			}
			if err := waitVisible(ctx, titleSel, 10*time.Second); err != nil {
				return err
			}
			var t string
			if err := chromedp.Run(ctx,
				chromedp.Text(titleSel, &t, chromedp.NodeVisible, chromedp.ByQuery),
			); err != nil {
				return err
			}
			title = strings.TrimSpace(t)
			return nil
		})
		if err != nil {
			continue
		}

		if !isAllowedDocURL(curURL) {
			_ = chromedp.Run(ctx, chromedp.Navigate(startURL))
			_ = waitVisible(ctx, navContainerSel, 15*time.Second)
			continue
		}

		// 별도 컨텍스트에서 InnerHTML 수집 (요구사항)
		var innerHTML string
		_ = withRetry(backoff, 3, func() error {
			ih, e := fetchInnerHTMLWithNewContext(ctx, curURL, 30*time.Second)
			if e != nil {
				return e
			}
			innerHTML = ih
			return nil
		})

		doc := Document{PostID: postID, Title: title, URL: curURL, InnerHTML: innerHTML, Content: ""}
		docs = append(docs, doc)
		visited[postID] = true
		logger.LogParsedDoc(nil, doc.PostID, doc.Title, doc.URL)

		if c.Limit > 0 && len(visited) >= c.Limit {
			break
		}
	}

	// Save result
	if err := saveOutput(outPath, format, docs); err != nil {
		return err
	}
	return nil
}

// Run executes the crawl with the provided configuration and writes results to outPath.
// Deprecated: use (*Crawler).Run with NewCrawler(clickDelay, limit, overallTimeout).
// Note: deprecated Run wrapper has been removed. Use NewCrawler(...).Run(...) instead.

func waitVisible(ctx context.Context, sel string, timeout time.Duration) error {
	c, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	return chromedp.Run(c, chromedp.WaitVisible(sel, chromedp.ByQuery))
}

func scrollMenuToEnd(ctx context.Context) error {
	// Scroll repeatedly until no progress
	const js = `(() => {
        const el = document.querySelector("#App > main > div.contents_wrap > div.tree_view_container");
        if (!el) return {ok:false, top:0, height:0};
        const before = el.scrollTop;
        el.scrollBy(0, 1000);
        return {ok:true, top: el.scrollTop, height: el.scrollHeight};
    })()`
	var lastTop int64 = -1
	for i := 0; i < 20; i++ { // safety bound
		var res struct {
			OK     bool  `json:"ok"`
			Top    int64 `json:"top"`
			Height int64 `json:"height"`
		}
		if err := chromedp.Run(ctx, chromedp.Evaluate(js, &res)); err != nil {
			return err
		}
		if !res.OK {
			return errors.New("nav container not found")
		}
		if res.Top == lastTop {
			break
		}
		lastTop = res.Top
		time.Sleep(50 * time.Millisecond)
	}
	return nil
}

func withRetry(b *Backoff, maxTries int, fn func() error) error {
	b.Reset()
	var err error
	for i := 0; i < maxTries; i++ {
		if err = fn(); err == nil {
			return nil
		}
		time.Sleep(b.Next())
	}
	return err
}

func isAllowedDocURL(raw string) bool {
	u, err := url.Parse(raw)
	if err != nil {
		return false
	}
	if !strings.HasSuffix(u.Hostname(), "nexon.com") {
		return false
	}
	if !strings.Contains(u.Path, "/docs") {
		return false
	}
	return true
}

func saveOutput(path, format string, docs []Document) error {
	switch format {
	case "json":
		return SaveJSON(path, docs)
	case "csv":
		return SaveCSV(path, docs)
	default:
		return fmt.Errorf("unknown format: %s", format)
	}
}
