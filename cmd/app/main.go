package main

import (
	"dz_first_pantela/iternal/database"
	"dz_first_pantela/iternal/handlers"
	"dz_first_pantela/iternal/messagesService"
	"dz_first_pantela/iternal/web/messages"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()
	//database.DB.AutoMigrate(&messagesService.Message{})

	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(*repo)

	handler := handlers.NewHandler(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := messages.NewStrictHandler(handler, nil)
	messages.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server: %s", err)
	}

	//router := mux.NewRouter()
	//router.HandleFunc("/api/message", handler.GetAllMessages).Methods("GET")
	//router.HandleFunc("/api/message", handler.CreateMessage).Methods("POST")
	//router.HandleFunc("/api/message/{id}", handler.UpdateMessageById).Methods("PUT")
	//router.HandleFunc("/api/message/{id}", handler.DeleteMessageById).Methods("DELETE")
	//
	//http.ListenAndServe(":8080", router)
}
