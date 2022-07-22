package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	db "teapp/db"
	"teapp/models"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/login", Auth)
	// router.GET("/:user", UserPage)
	log.Println("STARTING SERVER ON PORT :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	usrName := r.Header.Get("Username")
	usrPass, err := models.HashPassword(r.Header.Get("Password"))
	if err != nil {
		log.Println(err)
	}
	newUser := models.NewUser(usrName, usrPass, "Avatar", "Dima", "Koval")
	db.CreateNewUser(newUser, db.Connection())
}

func Auth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	usrName := r.Header.Get("Username")
	usrPass := r.Header.Get("Password")

	match := db.UserAuth(usrName, usrPass, db.Connection())
	if !match {
		fmt.Fprintf(w, "Введены не верные данные для входа")
		log.Println(match)
		return
	}
	fmt.Fprintf(w, "%s, Добро пожаловать!", usrName)
}

func UserPage(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	userId, err := db.GetUserID(ps.ByName("user"), db.Connection())
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	user := db.GetUser(userId, db.Connection())
	posts := db.GetUserPosts(userId, db.Connection())
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
