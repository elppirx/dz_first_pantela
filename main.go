package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	ids := mux.Vars(r)
	id, err := strconv.Atoi(ids["id"])
	if err != nil {
		fmt.Fprintln(w, "Ошибка чтения id")
	}
	var message Message
	DB.First(&message, id)

	decoder := json.NewDecoder(r.Body)
	var body Message
	err = decoder.Decode(&body)
	if err != nil {
		http.Error(w, "Ошибка чтения json", http.StatusBadRequest)
		return
	}

	message.Text = body.Text
	DB.Save(&message)
	fmt.Fprintf(w, "Сообщение с ID %v обновлено на %v", id, body.Text)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	ids := mux.Vars(r)
	id, err := strconv.Atoi(ids["id"])
	if err != nil {
		fmt.Fprintln(w, "Ошибка чтения id")
	}

	var message Message
	DB.First(&message, id)
	DB.Delete(&message)
	fmt.Fprintf(w, "Сообщение с ID %v удалено", id)
}

func main() {
	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()

	router.HandleFunc("/api/message", GetMessage).Methods("GET")
	router.HandleFunc("/api/message", AddMessage).Methods("POST")
	router.HandleFunc("/api/message/{id}", UpdateMessage).Methods("PUT")
	router.HandleFunc("/api/message/{id}", DeleteMessage).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
