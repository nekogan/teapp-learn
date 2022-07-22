package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println(err)
	}
	return err == nil
}
