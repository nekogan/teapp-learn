package models

type user struct {
	Info         base_user `json:"userInfo"`
	ProfileImage string    `json:"image"`
	FirstName    string    `json:"firstName"`
	SecondName   string    `json:"secondName"`
	Posts        []*post   `json:"posts"`
}

type base_user struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type User interface {
	AddPost(*post)
	AddInfo(string, string, string)
}

func (u *user) AddPost(post *post) {
	u.Posts = append(u.Posts, post)
}

func (u *user) AddInfo(fn, sn, pi string) {
	u.FirstName = fn
	u.SecondName = sn
	u.ProfileImage = pi
}

func NewBaseUser(usrn string, pass string) *base_user {
	return &base_user{
		ID:       1,
		Username: usrn,
		Password: pass,
	}
}

func NewUser(info *base_user) User {
	return &user{
		Info: *info,
	}
}
