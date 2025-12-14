package crawler

// Doc represents a crawled document item.
type Doc struct {
	PostID  string `json:"postId"`
	Title   string `json:"title"`
	URL     string `json:"url"`
	Content string `json:"content"`
}
