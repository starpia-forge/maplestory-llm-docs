package crawler

import (
	"net/url"
)

// ExtractPostIDFromURL parses the given URL string and returns the value of the
// "postId" query parameter if present. The second return value indicates
// whether a non-empty postId was found.
func ExtractPostIDFromURL(raw string) (string, bool) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", false
	}
	v := u.Query().Get("postId")
	if v == "" {
		return "", false
	}
	return v, true
}
