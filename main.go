package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	db "teapp/db"
	m "teapp/models"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/:user/", UserPosts)
	log.Println("STARTING SERVER ON PORT :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	user := m.NewUser("Nekogan", "password", "ImageURL", "Dima", "Koval")
	post := m.NewPost("Дракон", "Красный", "Самый лучший чай", 10)
	db.SaveToDB(user, post)
}

func UserPosts(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	posts := db.GetUserPosts(1)
	data, err := json.MarshalIndent(posts, "", "   ")
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, "All posts %+v", string(data))
}
