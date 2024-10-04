package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var Message string

type requestBody struct {
	Message string `json:"message"`
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body requestBody
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, "Ошибка чтения json", http.StatusBadRequest)
		return
	}
	Message = body.Message
	fmt.Fprintln(w, "Сообщение обновлено")

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %v!", Message)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/message", MessageHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
