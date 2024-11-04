package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/train-do/Golang-Web/model"
	"github.com/train-do/Golang-Web/service"
)

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := model.User{}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			badResponse := model.Response{
				StatusCode: http.StatusBadRequest,
				Message:    "Bad Request",
				Data:       nil,
			}
			json.NewEncoder(w).Encode(badResponse)
			return
		}
		serviceUser := service.ServiceUser{Db: db}
		err = serviceUser.Login(&user)
		if err != nil {
			unauthorized := model.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Login Failed",
				Data:       nil,
			}
			json.NewEncoder(w).Encode(unauthorized)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		response := model.Response{
			StatusCode: http.StatusOK,
			Message:    "Login Success",
			Data:       nil,
		}
		cookie := http.Cookie{
			Name:   "access_token",
			Value:  user.Id,
			Path:   "/",
			Domain: "localhost",
		}
		fmt.Println(cookie.Name, cookie.Value, "<<<<<<")
		http.SetCookie(w, &cookie)
		json.NewEncoder(w).Encode(response)
	}
}
func Register(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := model.User{}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			badResponse := model.Response{
				StatusCode: http.StatusBadRequest,
				Message:    "Bad Request",
				Data:       nil,
			}
			json.NewEncoder(w).Encode(badResponse)
			return
		}
		serviceUser := service.ServiceUser{Db: db}
		err = serviceUser.CreateUser(&user)
		if err != nil {
			badResponse := model.Response{
				StatusCode: http.StatusBadRequest,
				Message:    "Bad Request",
				Data:       nil,
			}
			json.NewEncoder(w).Encode(badResponse)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		response := model.Response{
			StatusCode: http.StatusCreated,
			Message:    "User registered",
			Data:       nil,
		}
		json.NewEncoder(w).Encode(response)
	}
}
