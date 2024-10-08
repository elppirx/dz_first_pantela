package main

import (
	"dz_first_pantela/iternal/database"
	"dz_first_pantela/iternal/handlers"
	"dz_first_pantela/iternal/messagesService"
	"dz_first_pantela/iternal/usersService"
	"dz_first_pantela/iternal/web/messages"
	"dz_first_pantela/iternal/web/users"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()
	//database.DB.AutoMigrate(&messagesService.Message{})
	//Messages
	messagesRepo := messagesService.NewMessageRepository(database.DB)
	usersRepo := usersService.NewUserRepository(database.DB)

	messagesService := messagesService.NewService(*messagesRepo)
	usersService := usersService.NewUsersService(*usersRepo)

	//Users

	messagesHandler := handlers.NewHandler(messagesService)
	usersHandler := handlers.NewUserHandler(usersService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	messagesStrictHandler := messages.NewStrictHandler(messagesHandler, nil)
	usersStrictHandler := users.NewStrictHandler(usersHandler, nil)

	messages.RegisterHandlers(e, messagesStrictHandler)
	users.RegisterHandlers(e, usersStrictHandler)

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
