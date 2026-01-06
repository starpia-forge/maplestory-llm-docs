package crawler

import (
	"bufio"
	"fmt"
	"os"
)

// Document represents a crawled document item.
type Document struct {
	Title     string `json:"title"`
	URL       string `json:"url"`
	InnerHTML string `json:"innerHTML"`
	Content   string `json:"content"`
}

// SaveDocumentFile creates/overwrites outFile and writes each Document.InnerHTML
// followed by a newline, streaming via a buffered writer to avoid excessive memory use.
func SaveDocumentFile(docs []Document, outFile string) error {
	f, err := os.OpenFile(outFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return fmt.Errorf("open %s: %w", outFile, err)
	}
	defer func() { _ = f.Close() }()

	// Use a reasonably large buffer to reduce syscalls but keep memory bounded.
	w := bufio.NewWriterSize(f, 64*1024)
	for i := range docs {
		if _, err := w.WriteString(docs[i].InnerHTML); err != nil {
			return fmt.Errorf("write innerHTML: %w", err)
		}
		if err := w.WriteByte('\n'); err != nil {
			return fmt.Errorf("write newline: %w", err)
		}
	}
	if err := w.Flush(); err != nil {
		return fmt.Errorf("flush: %w", err)
	}
	return nil
}
