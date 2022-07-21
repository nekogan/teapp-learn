package db

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"gopkg.in/yaml.v3"
)

type db struct {
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

func config() string {
	d := db{}
	file, err := ioutil.ReadFile("db/config.yml")
	if err != nil {
		log.Println(err)
	}

	err = yaml.Unmarshal(file, &d)
	if err != nil {
		log.Println(err)
	}
	return fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		d.Host, d.Port, d.User, d.Password, d.Dbname)
}

// TODO:Допилить функцию для получения id пользователя по никнейму
func Connection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), config())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	log.Println("CONNECTED TO POSTGRE")
	return conn
}
