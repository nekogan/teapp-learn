package models

type Post struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Text     string `json:"text"`
}

func NewPost(t, c, txt string) Post {
	return Post{
		Title:    t,
		Category: c,
		Text:     txt,
	}
}
