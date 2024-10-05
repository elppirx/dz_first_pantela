package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func AddMessage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body Message
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, "Ошибка чтения json", http.StatusBadRequest)
		return
	}
	DB.Create(&body)
	fmt.Fprintf(w, "Сообщение %v добавлено", body.Text)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	var getResult []Message
	DB.Find(&getResult)
	fmt.Fprintln(w, getResult)
}

func main() {
	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()

	router.HandleFunc("/api/message", GetMessage).Methods("GET")
	router.HandleFunc("/api/message", AddMessage).Methods("POST")

	http.ListenAndServe(":8080", router)
}
