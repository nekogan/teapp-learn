package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"teapp/models"

	"github.com/jackc/pgx/v4"
)

func SaveToDB(u *models.User, p models.Post, conn *pgx.Conn) {
	defer conn.Close(context.Background())
	sqlStatement := fmt.Sprintf(`INSERT INTO post (user_id, post_title, post_classification, post_text, post_rating) 
		VALUES ('%d', '%s', '%s', '%s', '%d') RETURNING post_id;`,
		u.ID, p.Title, p.Classification, p.Text, p.Rating)
	post_id := 0
	err := conn.QueryRow(context.Background(), sqlStatement).Scan(&post_id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("New record ID is:", post_id)
}

func GetUser(usrID uint, conn *pgx.Conn) models.User {
	defer conn.Close(context.Background())
	sql := fmt.Sprintf(`select user_name, user_avatar, user_firstname, user_secondname from users where user_id='%d'`,
		usrID)
	var u models.User
	err := conn.QueryRow(context.Background(), sql).Scan(&u.Username, &u.Avatar, &u.FirstName, &u.SecondName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	return u
}

func GetUserPosts(usrID uint, conn *pgx.Conn) []models.Post {
	defer conn.Close(context.Background())
	sqlStatement := fmt.Sprintf(`select post_id, post_title, post_classification, post_text, post_rating from post where user_id='%d'`,
		usrID)
	rows, err := conn.Query(context.Background(), sqlStatement)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
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
