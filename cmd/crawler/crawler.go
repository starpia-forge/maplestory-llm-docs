package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"maplestory-world-llms-txt/internal/crawler"
)

var (
	targets = map[string]string{
		/* Reference docs */
		"https://maplestoryworlds-creators.nexon.com/ko/docs/?postId=472": "docs/kr/reference.md",
		"https://maplestoryworlds-creators.nexon.com/en/docs/?postId=472": "docs/en/reference.md",
		/* API docs */
		"https://maplestoryworlds-creators.nexon.com/ko/apiReference/How-to-use-API-Reference": "docs/kr/api.md",
		"https://maplestoryworlds-creators.nexon.com/en/apiReference/How-to-use-API-Reference": "docs/en/api.md",
	}
)

func main() {
	var (
		head    bool
		delay   time.Duration
		limit   int
		timeout time.Duration
	)

	flag.BoolVar(&head, "headless", true, "run headless Chrome")
	flag.DurationVar(&delay, "delay", 150*time.Millisecond, "delay between clicks")
	flag.IntVar(&limit, "limit", 0, "max number of documents to crawl (0 = no limit)")
	flag.DurationVar(&timeout, "timeout", 120*time.Second, "overall timeout for crawling")
	flag.Parse()

	// Configure default slog logger (text to stderr, Info level)
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo})
	slog.SetDefault(slog.New(handler))

	c := crawler.NewCrawler(
		crawler.WithClickDelay(delay),
		crawler.WithLimit(limit),
		crawler.WithOverallTimeout(timeout),
		crawler.WithHeadless(head),
	)

	for targetURL, outFileName := range targets {
		docs, err := c.Run(targetURL)
		if err != nil {
			log.Fatalf("crawler error: %v", err)
		}
		log.Printf("crawled %d documents from %q", len(docs), targetURL)

		// Create temp dir to store individual HTML doc files
		tmpDir, err := os.MkdirTemp("", "crawler_docs_*")
		if err != nil {
			log.Fatalf("create temp dir: %v", err)
		}
		defer os.RemoveAll(tmpDir)

		paths, err := crawler.SaveDocumentFile(docs, tmpDir)
		if err != nil {
			log.Fatalf("SaveDocumentFile error: %v", err)
		}

		// Convert each HTML fragment to Markdown individually via mdream
		mdTmpDir, err := os.MkdirTemp("", "md_parts_*")
		if err != nil {
			log.Fatalf("create temp md dir: %v", err)
		}
		defer os.RemoveAll(mdTmpDir)

		mdParts := make([]string, 0, len(paths))
		for i, p := range paths {
			partOut := filepath.Join(mdTmpDir, fmt.Sprintf("%03d_%s", i, filepath.Base(outFileName)))
			if err := mdream(p, partOut); err != nil {
				log.Fatalf("mdream error for %s: %v", p, err)
			}
			mdParts = append(mdParts, partOut)
		}

		// Concatenate all generated Markdown parts into the final output file
		outF, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
		if err != nil {
			log.Fatalf("open %s: %v", outFileName, err)
		}
		for i, p := range mdParts {
			inF, err := os.Open(p)
			if err != nil {
				_ = outF.Close()
				log.Fatalf("open %s: %v", p, err)
			}
			if _, err := io.Copy(outF, inF); err != nil {
				_ = inF.Close()
				_ = outF.Close()
				log.Fatalf("copy from %s: %v", p, err)
			}
			_ = inF.Close()
			// Separate documents with a newline to preserve previous behavior
			if i < len(mdParts)-1 {
				if _, err := outF.WriteString("\n"); err != nil {
					_ = outF.Close()
					log.Fatalf("write newline: %v", err)
				}
			}
		}
		if err := outF.Close(); err != nil {
			log.Fatalf("close %s: %v", outFileName, err)
		}
		log.Printf("wrote concatenated markdown to %s (from %d parts in %s)", outFileName, len(mdParts), filepath.Base(mdTmpDir))
	}
}
func mdream(inputFileName, outFileName string) error {
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	cmd := exec.Command("npx", "mdream", "--preset", "minimal")
	cmd.Stdin = inputFile
	cmd.Stdout = io.Writer(outputFile)
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
