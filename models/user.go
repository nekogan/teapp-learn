package models

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         uint   `json:"user_id"`
	Username   string `json:"username"`
	Pass       string `json:"user_pass"`
	FirstName  string `json:"user_firstName"`
	SecondName string `json:"user_secondName"`
	Avatar     string `json:"user_avatar"`
	Posts      []Post `json:"user_posts"`
}

func NewUser(u, p, fn, sn, aurl string) User {
	return User{
		Username:   u,
		Pass:       p,
		FirstName:  fn,
		SecondName: sn,
		Avatar:     aurl,
	}
}

func checkUsername(username string, conn *pgx.Conn) error {
	sqlStatement := fmt.Sprintf(`select user_id from users where user_name='%s'`, username)
	var user_id uint
	err := conn.QueryRow(context.Background(), sqlStatement).Scan(&user_id)
	if err != nil {
		return nil
	}

	return errors.New("такой пользователь уже существует")
}

func CreateNewUser(u User, conn *pgx.Conn) error {
	defer conn.Close(context.Background())
	err := checkUsername(u.Username, conn)
	if err != nil {
		return err
	}
	sqlStatement := fmt.Sprintf(`INSERT INTO users (user_name, user_pass, user_avatar, user_firstname, user_secondname) 
		VALUES ('%s', '%s', '%s', '%s', '%s') RETURNING user_id;`,
		u.Username, u.Pass, u.Avatar, u.FirstName, u.SecondName)
	var user_id uint
	err = conn.QueryRow(context.Background(), sqlStatement).Scan(&user_id)
	if err != nil {
		return errors.New("не удалось создать пользователя")
	}
	fmt.Println("New record ID is:", user_id)

	return nil
}

func UserAuth(user, pass string, conn *pgx.Conn) bool {
	defer conn.Close(context.Background())
	sql := fmt.Sprintf(`select user_pass from users where user_name='%s'`,
		user)
	var password string
	err := conn.QueryRow(context.Background(), sql).Scan(&password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Пользователь не найден: %v\n", err)
		return false
	}

	return CheckPasswordHash(pass, password)
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
