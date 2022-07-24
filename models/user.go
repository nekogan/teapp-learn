package models

type User struct {
	ID         uint   `json:"user_id"`
	Username   string `json:"username"`
	Pass       string `json:"user_pass"`
	Avatar     string `json:"user_avatar"`
	FirstName  string `json:"user_firstName"`
	SecondName string `json:"user_secondName"`
}

func NewUser(u, p, aurl, fn, sn string) *User {
	return &User{
		ID:         1,
		Username:   u,
		Pass:       p,
		Avatar:     aurl,
		FirstName:  fn,
		SecondName: sn,
	}
}
