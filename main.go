package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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
	router.POST("/addpost", AddPost)
	log.Println("STARTING SERVER ON PORT :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Добро пожаловать в Teapp")
}

func AddPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	conn := db.Connection()
	defer conn.Close(context.Background())
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	var body struct {
		User_name string      `json:"username"`
		Post      models.Post `json:"post"`
	}
	err = json.Unmarshal(b, &body)
	if err != nil {
		log.Println(err)
		return
	}
	user_id, err := db.GetUserID(body.User_name, conn)
	if err != nil {
		fmt.Fprintf(w, "Ошибка: %v", err)
	}
	newPost := models.NewPost(body.Post.Title, body.Post.Category, body.Post.Text, body.Post.Category)
	err = db.SaveToDB(user_id, newPost, conn)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintf(w, "Запись успешно добавлена!")
}

func Registration(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user, password, _ := r.BasicAuth()
	passHash, err := models.HashPassword(password)
	if err != nil {
		fmt.Fprintf(w, "Ошибка: %v", err)
		return
	}

	newUser := models.NewUser(user, passHash, "Dima", "Koval", "")
	if err := models.CreateNewUser(newUser, db.Connection()); err != nil {
		fmt.Fprintf(w, "Ошибка: %v", err)
		return
	}

	fmt.Fprintf(w, "Вы зарегистрировались! Добро пожаловать!")
}

func Auth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user, password, _ := r.BasicAuth()

	match := models.UserAuth(user, password, db.Connection())
	if !match {
		fmt.Fprintf(w, "Введены не верные данные для входа\n")
		w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		return
	}
	w.Header().Set("WWW-Authenticate", "Basic realm=Access to the staging site")

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

	user, err := db.GetUser(userId, conn)
	log.Println(user)
	if err != nil {
		fmt.Fprintf(w, "Ошибка: %v", err)
		return
	}

	posts := db.GetUserPosts(userId, conn)
	user.Posts = posts
	userInfo, err := json.MarshalIndent(user, "", "   ")
	if err != nil {
		fmt.Fprintf(w, "Ошибка: %v", err)
		return
	}

	fmt.Fprintf(w, "User: %+v\n", string(userInfo))
}
