package crawler

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveDocumentFile_WritesFilesPerIndexAndReturnsPaths(t *testing.T) {
	docs := []Document{
		{InnerHTML: "<div>first</div>"},
		{InnerHTML: "<p>second</p>"},
		{InnerHTML: "<span>third</span>"},
	}
	dir := t.TempDir()
	paths, err := SaveDocumentFile(docs, dir)
	if err != nil {
		t.Fatalf("SaveDocumentFile error: %v", err)
	}
	if len(paths) != len(docs) {
		t.Fatalf("expected %d paths, got %d", len(docs), len(paths))
	}
	// Check concrete names and contents
	expectNames := []string{
		filepath.Join(dir, "0_doc.html"),
		filepath.Join(dir, "1_doc.html"),
		filepath.Join(dir, "2_doc.html"),
	}
	for i, p := range paths {
		if p != expectNames[i] {
			t.Fatalf("unexpected path at %d: want %q got %q", i, expectNames[i], p)
		}
		b, err := os.ReadFile(p)
		if err != nil {
			t.Fatalf("read file %s: %v", p, err)
		}
		if string(b) != docs[i].InnerHTML {
			t.Fatalf("unexpected content for %s: want %q got %q", p, docs[i].InnerHTML, string(b))
		}
	}
}
