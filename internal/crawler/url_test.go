package crawler

import "testing"

func TestExtractPostIDFromURL_ShouldReturnValueWhenPresent(t *testing.T) {
	id, ok := ExtractPostIDFromURL("https://example.com/docs?postId=472&lang=ko")
	if !ok || id != "472" {
		t.Fatalf("expected ok=true and id=472, got ok=%v id=%q", ok, id)
	}
}

func TestExtractPostIDFromURL_ShouldReturnFalseWhenMissing(t *testing.T) {
	if id, ok := ExtractPostIDFromURL("https://example.com/docs?foo=bar"); ok || id != "" {
		t.Fatalf("expected ok=false and id empty, got ok=%v id=%q", ok, id)
	}
}

func TestExtractPostIDFromURL_ShouldReturnFalseOnInvalidURL(t *testing.T) {
	if id, ok := ExtractPostIDFromURL(":nope"); ok || id != "" {
		t.Fatalf("expected ok=false and id empty for invalid url, got ok=%v id=%q", ok, id)
	}
}
