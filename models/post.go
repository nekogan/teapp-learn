package models

type Post struct {
	ID             uint   `json:"id"`
	Title          string `json:"title"`
	Classification string `json:"classification"`
	Text           string `json:"text"`
	Rating         byte   `json:"rating"`
}

func NewPost(t, c, txt string, rat byte) Post {
	return Post{
		Title:          t,
		Classification: c,
		Text:           txt,
		Rating:         rat,
	}
}
