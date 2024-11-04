package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/train-do/Golang-Web/database"
	"github.com/train-do/Golang-Web/handler"
	"github.com/train-do/Golang-Web/middleware"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	userMux := http.NewServeMux()
	userMux.HandleFunc("POST /register", handler.Register(db))
	userMux.HandleFunc("POST /login", handler.Login(db))

	todoMux := http.NewServeMux()
	todoMux.HandleFunc("GET /", handler.GetTodo(db))
	todoMux.HandleFunc("POST /", handler.CreateTodo(db))
	// todoMux.HandleFunc("PUT /{id}", handler.EditTodo(db))
	// todoMux.HandleFunc("DELETE /{id}", handler.DeleteTodo(db))

	serverMux := http.NewServeMux()
	serverMux.Handle("/", userMux)
	serverMux.Handle("/todo", middleware.Authentication(db, todoMux))
	fmt.Println("server started on port 8080")
	http.ListenAndServe(":8080", serverMux)
}
