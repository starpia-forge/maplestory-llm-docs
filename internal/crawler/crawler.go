package crawler

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
	"maplestory-llm-docs/internal/logger"
)

// StartURL is the initial page to begin crawling.
const StartURL = "https://maplestoryworlds-creators.nexon.com/ko/docs/?postId=472"

const (
	navContainerSel     = "#App > main > div.contents_wrap > div.tree_view_container"
	contentContainerSel = "#App > main > div.contents_wrap > div.renderContent > div.text_content_container > div.text_content"
	titleSel            = "#App > main > div.contents_wrap > div.renderContent h1"
)

// Run executes the crawl with the provided configuration and writes results to outPath.
func Run(headless bool, outPath, format string, clickDelay time.Duration, limit int, overallTimeout time.Duration) error {
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

	if overallTimeout > 0 {
		var toCancel context.CancelFunc
		ctx, toCancel = context.WithTimeout(ctx, overallTimeout)
		defer toCancel()
	}

	// Navigate to start URL
	if err := chromedp.Run(ctx, chromedp.Navigate(StartURL)); err != nil {
		return err
	}
	if err := waitVisible(ctx, navContainerSel, 30*time.Second); err != nil {
		return fmt.Errorf("navigation container not visible: %w", err)
	}

	visited := make(map[string]bool)
	var docs []Doc
	backoff := NewBackoff(500*time.Millisecond, 20*time.Second, 2.0, 0.2)

	// Scanning loop: continue until no new docs are found in a pass
	for {
		// Scroll to ensure all lazy items are present
		_ = scrollMenuToEnd(ctx)

		var nodes []*cdp.Node
		if err := chromedp.Run(ctx, chromedp.Nodes(navContainerSel+" *", &nodes, chromedp.ByQueryAll)); err != nil {
			return fmt.Errorf("query nodes: %w", err)
		}

		foundNew := false

		for _, n := range nodes {
			// Scroll node into view if needed and click
			_ = chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
				return dom.ScrollIntoViewIfNeeded().WithNodeID(n.NodeID).Do(ctx)
			}))
			if err := chromedp.Run(ctx, chromedp.MouseClickNode(n)); err != nil {
				// ignore elements that cannot be clicked
				continue
			}

			// politeness delay
			time.Sleep(clickDelay)

			// After click, check if we are on a doc page (via postId in URL)
			var curURL string
			_ = chromedp.Run(ctx, chromedp.Location(&curURL))
			postID, ok := ExtractPostIDFromURL(curURL)
			if !ok {
				// could be an expander; continue
				continue
			}
			if visited[postID] {
				continue
			}

			// Wait for content and extract, with retry/backoff for transient failures
			var title, html string
			err := withRetry(backoff, 5, func() error {
				if err := waitVisible(ctx, contentContainerSel, 30*time.Second); err != nil {
					return err
				}
				// ensure title present too
				if err := waitVisible(ctx, titleSel, 10*time.Second); err != nil {
					return err
				}
				var t, h string
				if err := chromedp.Run(ctx,
					chromedp.Text(titleSel, &t, chromedp.NodeVisible, chromedp.ByQuery),
					chromedp.InnerHTML(contentContainerSel, &h, chromedp.ByQuery),
				); err != nil {
					return err
				}
				title, html = strings.TrimSpace(t), h
				return nil
			})
			if err != nil {
				// If content couldn't be retrieved, skip this item
				continue
			}

			// verify URL stays within docs domain
			if !isAllowedDocURL(curURL) {
				// navigate back to start to avoid getting stuck on external pages
				_ = chromedp.Run(ctx, chromedp.Navigate(StartURL))
				_ = waitVisible(ctx, navContainerSel, 15*time.Second)
				continue
			}

			doc := Doc{PostID: postID, Title: title, URL: curURL, Content: html}
			docs = append(docs, doc)
			visited[postID] = true
			foundNew = true

			// info log for each newly parsed document
			logger.LogParsedDoc(nil, doc.PostID, doc.Title, doc.URL)

			if limit > 0 && len(visited) >= limit {
				break
			}
		}

		if limit > 0 && len(visited) >= limit {
			break
		}
		if !foundNew {
			break
		}
	}

	// Save result
	if err := saveOutput(outPath, format, docs); err != nil {
		return err
	}
	return nil
}

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

func saveOutput(path, format string, docs []Doc) error {
	switch format {
	case "json":
		return SaveJSON(path, docs)
	case "csv":
		return SaveCSV(path, docs)
	default:
		return fmt.Errorf("unknown format: %s", format)
	}
}
