package main

import (
	"dz_first_pantela/iternal/database"
	"dz_first_pantela/iternal/handlers"
	"dz_first_pantela/iternal/messagesService"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()
	//database.DB.AutoMigrate(&messagesService.Message{})

	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(*repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/message", handler.GetAllMessages).Methods("GET")
	router.HandleFunc("/api/message", handler.CreateMessage).Methods("POST")
	router.HandleFunc("/api/message/{id}", handler.UpdateMessageById).Methods("PUT")
	router.HandleFunc("/api/message/{id}", handler.DeleteMessageById).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
