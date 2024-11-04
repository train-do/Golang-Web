package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/train-do/Golang-Web/model"
	"github.com/train-do/Golang-Web/service"
)

func GetTodo(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// todo := model.Todo{}
		// err := json.NewDecoder(r.Body).Decode(&todo)
		// if err != nil {
		// 	fmt.Println("Error Decode:", err)
		// 	badResponse := model.Response{
		// 		StatusCode: http.StatusBadRequest,
		// 		Message:    "Bad Request Decode",
		// 		Data:       nil,
		// 	}
		// 	json.NewEncoder(w).Encode(badResponse)
		// 	return
		// }
		serviceTodo := service.ServiceTodo{Db: db}
		// todos, err := serviceTodo.FindAllTodo(todo.UserId)
		todos, err := serviceTodo.FindAllTodo("7392af7b-cd88-4a0d-9c8d-328da206a562")
		if err != nil {
			ise := model.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    "Internal Server Error",
				Data:       nil,
			}
			json.NewEncoder(w).Encode(ise)
			return
		}
		response := model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error",
			Data:       todos,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}
func CreateTodo(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todo := model.Todo{}
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			fmt.Println("Error Decode:", err)
			badResponse := model.Response{
				StatusCode: http.StatusBadRequest,
				Message:    "Bad Request Decode",
				Data:       nil,
			}
			json.NewEncoder(w).Encode(badResponse)
			return
		}
		serviceTodo := service.ServiceTodo{Db: db}
		err = serviceTodo.InsertTodo(&todo)
		if err != nil {
			unauthorized := model.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Bad Request Service",
				Data:       nil,
			}
			json.NewEncoder(w).Encode(unauthorized)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(todo)
	}
}
