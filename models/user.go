package models

type User struct {
	ID              uint   `json:"id"`
	Username        string `json:"usrnm"`
	Pass            string `json:"pass"`
	ProfileImageURL string `json:"imageURL"`
	FirstName       string `json:"firstName"`
	SecondName      string `json:"secondName"`
}

func NewUser(u, p, piurl, fn, sn string) *User {
	return &User{
		ID:              1,
		Username:        u,
		Pass:            p,
		ProfileImageURL: piurl,
		FirstName:       fn,
		SecondName:      sn,
	}
}
