package models

type post struct {
	ID     uint   `json:"id"`
	Tea    tea    `json:"tea"`
	Text   string `json:"text"`
	Rating byte   `json:"rating"`
}

func NewPost(id uint, t tea, text string, rating byte) post {
	return post{
		ID:     id,
		Tea:    t,
		Text:   text,
		Rating: rating,
	}
}
