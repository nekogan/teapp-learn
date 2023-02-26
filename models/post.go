package models

type Post struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Text     string `json:"text"`
	Tags     string `json:"tags"`
}

func NewPost(title, cat, txt, tags string) Post {
	return Post{
		Title:    title,
		Category: cat,
		Text:     txt,
		Tags:     tags,
	}
}
