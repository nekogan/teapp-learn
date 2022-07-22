package db

import (
	"context"
	"fmt"
	"os"
	"teapp/models"

	"github.com/jackc/pgx/v4"
)

func CreateNewUser(u *models.User, conn *pgx.Conn) {
	defer conn.Close(context.Background())
	sqlStatement := fmt.Sprintf(`INSERT INTO users (user_name, user_pass, user_avatar, user_firstname, user_secondname) 
		VALUES ('%s', '%s', '%s', '%s', '%s') RETURNING user_id;`,
		u.Username, u.Pass, u.Avatar, u.FirstName, u.SecondName)
	user_id := 0
	err := conn.QueryRow(context.Background(), sqlStatement).Scan(&user_id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("New record ID is:", user_id)
}

func UserAuth(user, pass string, conn *pgx.Conn) bool {
	defer conn.Close(context.Background())
	sql := fmt.Sprintf(`select user_pass from users where user_name='%s'`,
		user)
	var password string
	err := conn.QueryRow(context.Background(), sql).Scan(&password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return models.CheckPasswordHash(pass, password)
}
