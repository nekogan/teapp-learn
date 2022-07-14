package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	m "teapp/models"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := m.NewTea(1, "Дракон", "Шу Пуэр")
	post := m.NewPost(1, t, "Очевидно самый вкусный чай", 10)
	data, err := json.MarshalIndent(post, "", "   ")
	if err != nil {
		log.Println(err)
	}
	fmt.Fprint(w, string(data))
}
