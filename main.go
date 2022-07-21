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
	router.POST("/", Index)
	router.GET("/:user", UserPage)
	log.Println("STARTING SERVER ON PORT :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	user := m.NewUser("Nekogan", "password", "ImageURL", "Dima", "Koval")
	post := m.NewPost("Дракон", "Красный", "Самый лучший чай", 10)
	db.SaveToDB(user, post, db.Connection(ps.ByName("user")))
}

func UserPage(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	user := db.GetUser(uint(db.UserID), db.Connection(ps.ByName("user")))
	posts := db.GetUserPosts(uint(db.UserID), db.Connection(ps.ByName("user")))
	userInfo, err := json.MarshalIndent(user, "", "   ")
	if err != nil {
		log.Println(err)
	}

	userPosts, err := json.MarshalIndent(posts, "", "   ")
	if err != nil {
		log.Println(err)
	}

	fmt.Fprintf(w, "User: %+v\n UserPosts: %+v", string(userInfo), string(userPosts))
}
