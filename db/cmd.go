package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"teapp/models"
)

func SaveToDB(u *models.User, p models.Post) {
	conn := Connection()
	defer conn.Close(context.Background())
	sqlStatement := fmt.Sprintf(`INSERT INTO post (user_id, post_title, post_classification, post_text, post_rating) 
		VALUES ('%d', '%s', '%s', '%s', '%d') RETURNING post_id;`,
		u.ID, p.Title, p.Classification, p.Text, p.Rating)
	fmt.Println(sqlStatement)
	post_id := 0
	err := conn.QueryRow(context.Background(), sqlStatement).Scan(&post_id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	fmt.Println("New record ID is:", post_id)
}

func GetUserPosts(usrID int) []models.Post {
	conn := Connection()
	defer conn.Close(context.Background())

	sqlStatement := fmt.Sprintf(`select post_id, post_title, post_classification, post_text, post_rating from post where user_id='%d'`, usrID)
	rows, err := conn.Query(context.Background(), sqlStatement)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

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
