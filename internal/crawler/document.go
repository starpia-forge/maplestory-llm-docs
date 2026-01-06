package crawler

import (
	"fmt"
	"os"
	"path/filepath"
)

// Document represents a crawled document item.
type Document struct {
	Title     string `json:"title"`
	URL       string `json:"url"`
	InnerHTML string `json:"innerHTML"`
	Content   string `json:"content"`
}

// SaveDocumentFile writes each Document.InnerHTML into a separate file under outFileDir.
// Files are named using the slice index: "<index>_doc.html". It returns the full paths
// of the created files in the same order as docs.
func SaveDocumentFile(docs []Document, outFileDir string) ([]string, error) {
	paths := make([]string, 0, len(docs))
	for i := range docs {
		name := fmt.Sprintf("%d_doc.html", i)
		p := filepath.Join(outFileDir, name)
		f, err := os.OpenFile(p, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
		if err != nil {
			return nil, fmt.Errorf("open %s: %w", p, err)
		}
		if _, err := f.WriteString(docs[i].InnerHTML); err != nil {
			_ = f.Close()
			return nil, fmt.Errorf("write innerHTML: %w", err)
		}
		if err := f.Close(); err != nil {
			return nil, fmt.Errorf("close file: %w", err)
		}
		paths = append(paths, p)
	}
	return paths, nil
}
