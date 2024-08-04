package main

import (
	"github.com/grandleemon/golang-todo-list"
	"github.com/grandleemon/golang-todo-list/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error starting server: %s", err.Error())
	}
}
