package main

import (
	"flag"
	"log"
	"log/slog"
	"os"
	"strings"
	"time"

	"maplestory-llm-docs/internal/crawler"
)

func main() {
	var (
		out      string
		format   string
		head     bool
		delay    time.Duration
		limit    int
		timeout  time.Duration
		startURL string
	)

	flag.StringVar(&out, "out", "docs.json", "output file path")
	flag.StringVar(&format, "format", "json", "output format: json|csv")
	flag.BoolVar(&head, "headless", true, "run headless Chrome")
	flag.DurationVar(&delay, "delay", 150*time.Millisecond, "delay between clicks")
	flag.IntVar(&limit, "limit", 0, "max number of documents to crawl (0 = no limit)")
	flag.DurationVar(&timeout, "timeout", 120*time.Second, "overall timeout for crawling")
	flag.StringVar(&startURL, "start-url", crawler.StartURL, "start URL to begin crawling")
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
	if err := c.Run(out, strings.ToLower(format), startURL); err != nil {
		log.Fatalf("crawler error: %v", err)
	}
}
