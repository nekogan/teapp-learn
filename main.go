package main

import (
	"context"
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
	router.POST("/registration", Registration)
	router.GET("/login", Auth)
	router.GET("/user/:user", UserPage)
	log.Println("STARTING SERVER ON PORT :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Добро пожаловать в Teapp")
}

func Registration(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	usrName := r.Header.Get("Username")
	usrPass, err := db.HashPassword(r.Header.Get("Password"))
	if err != nil {
		log.Println(err)
	}
	newUser := models.NewUser(usrName, usrPass, "Avatar", "Dima", "Koval")
	if err := db.CreateNewUser(newUser, db.Connection()); err != nil {
		fmt.Fprintf(w, "Ошибка: %v", err)
		return
	}

	fmt.Fprintf(w, "Вы зарегистрировались! Добро пожаловать!")
}

func Auth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user, password, _ := r.BasicAuth()

	match := db.UserAuth(user, password, db.Connection())
	if !match {
		fmt.Fprintf(w, "Введены не верные данные для входа")
		w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	w.Header().Set("WWW-Authenticate", "Basic realm=Access to the staging site")
	http.Error(w, http.StatusText(http.StatusAccepted), http.StatusAccepted)
	
	fmt.Fprintf(w, "%s, Добро пожаловать!", user)
}

func UserPage(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	conn := db.Connection()
	defer conn.Close(context.Background())
	userId, err := db.GetUserID(ps.ByName("user"), conn)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	user := db.GetUser(userId, conn)
	posts := db.GetUserPosts(userId, conn)
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
