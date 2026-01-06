package crawler

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func sampleDocs() []Document {
	return []Document{
		{Title: "Hello", URL: "https://example.com?a=1", Content: "Line1\nLine2"},
		{Title: "World, CSV", URL: "https://example.com?a=2", Content: "Comma, inside"},
	}
}

func TestEncodeJSON_RoundTrip(t *testing.T) {
	docs := sampleDocs()
	data, err := EncodeJSON(docs)
	if err != nil {
		t.Fatalf("EncodeJSON error: %v", err)
	}
	var got []Document
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(got) != len(docs) || got[1].Title != docs[1].Title || got[0].Content != docs[0].Content {
		t.Fatalf("mismatch: got=%+v", got)
	}
}

func TestSaveJSON_WritesFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.json")
	if err := SaveJSON(path, sampleDocs()); err != nil {
		t.Fatalf("SaveJSON: %v", err)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("expected file written: %v", err)
	}
}

func TestEncodeCSV_AndReadBack(t *testing.T) {
	data, err := EncodeCSV(sampleDocs())
	if err != nil {
		t.Fatalf("EncodeCSV: %v", err)
	}
	r := csv.NewReader(bytesReader(data))
	rows, err := r.ReadAll()
	if err != nil {
		t.Fatalf("read csv: %v", err)
	}
	if len(rows) != 3 {
		t.Fatalf("expected 3 rows (header + 2), got %d", len(rows))
	}
	if rows[0][0] != "title" || rows[0][1] != "url" || rows[0][2] != "content" {
		t.Fatalf("unexpected header: %+v", rows[0])
	}
	if rows[1][0] != "Hello" || rows[2][0] != "World, CSV" {
		t.Fatalf("unexpected data rows: %+v", rows[1:])
	}
}

func TestSaveCSV_WritesFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.csv")
	if err := SaveCSV(path, sampleDocs()); err != nil {
		t.Fatalf("SaveCSV: %v", err)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("expected file written: %v", err)
	}
}

// bytesReader wraps a byte slice as an io.Reader without importing bytes in this file's import block.
type bytesReader []byte

func (b bytesReader) Read(p []byte) (int, error) {
	n := copy(p, b)
	if n == len(b) {
		return n, io.EOF
	}
	return n, nil
}
