package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"teapp/models"

	"github.com/jackc/pgx/v4"
)

func SaveToDB(u models.User, p models.Post) {
	conn, err := pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	log.Println("CONNECTED TO POSTGRE")
	defer conn.Close(context.Background())

	sqlStatement := fmt.Sprintf(`INSERT INTO post (user_id, post_title, post_classification, post_text, post_rating) VALUES ('%d', '%s', '%s', '%s', '%d') RETURNING post_id;`, u.ID, p.Title, p.Classification, p.Text, p.Rating)
	fmt.Println(sqlStatement)
	post_id := 0
	err = conn.QueryRow(context.Background(), sqlStatement).Scan(&post_id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	fmt.Println("New record ID is:", post_id)
}
