package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"teapp/models"

	"github.com/jackc/pgx/v4"
)

func SaveToDB(u uint, p models.Post, conn *pgx.Conn) error {
	sqlStatement := fmt.Sprintf(`INSERT INTO post (user_id, post_title, post_classification, post_text, post_rating) 
		VALUES ('%d', '%s', '%s', '%s', '%d') RETURNING post_id;`,
		u, p.Title, p.Classification, p.Text, p.Rating)
	var post_id uint
	err := conn.QueryRow(context.Background(), sqlStatement).Scan(&post_id)
	if err != nil {
		return err
	}
	return nil
}

func GetUserID(user string, conn *pgx.Conn) (uint, error) {
	sql := fmt.Sprintf(`select user_id from users where user_name='%s'`, user)
	var usrID uint
	err := conn.QueryRow(context.Background(), sql).Scan(&usrID)
	if err != nil {
		return 0, errors.New("пользователь не найден")
	}
	return usrID, nil
}

func GetUser(usrID uint, conn *pgx.Conn) (models.User, error) {
	sql := fmt.Sprintf(`select user_name, user_pass, user_avatar, user_firstname, user_secondname from users where user_id='%d'`,
		usrID)
	var u models.User
	err := conn.QueryRow(context.Background(), sql).Scan(&u.Username, &u.Pass, &u.Avatar, &u.FirstName, &u.SecondName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Пользователь не найден: %v\n", err)
		return models.User{}, err
	}

	return u, nil
}

func GetUserPosts(usrID uint, conn *pgx.Conn) []models.Post {
	sqlStatement := fmt.Sprintf(`select post_id, post_title, post_classification, post_text, post_rating from post where user_id='%d'`,
		usrID)
	rows, err := conn.Query(context.Background(), sqlStatement)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Нет записей: %v\n", err)
	}

	var rowSlice []models.Post
	for rows.Next() {
		var p models.Post
		err := rows.Scan(&p.ID, &p.Title, &p.Classification, &p.Text, &p.Rating)
		if err != nil {
			log.Fatal(err)
		}
		rowSlice = append(rowSlice, p)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return rowSlice
}
