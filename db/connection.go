package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "nekogan"
	password = "root"
	dbname   = "teapp_db"
)

var DATABASE_URL = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func Connection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	log.Println("CONNECTED TO POSTGRE")
	return conn
}