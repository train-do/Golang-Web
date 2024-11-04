package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/train-do/Golang-Web/model"
	"github.com/train-do/Golang-Web/service"
)

func Authentication(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("access_token")
		if err != nil {
			if err == http.ErrNoCookie {
				unauthorized := model.Response{
					StatusCode: http.StatusUnauthorized,
					Message:    "No Access Token",
					Data:       nil,
				}
				json.NewEncoder(w).Encode(unauthorized)
				return
			}
			unauthorized := model.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Error Retrieving Cookie",
				Data:       nil,
			}
			json.NewEncoder(w).Encode(unauthorized)
			return
		}
		accessToken := cookie.Value
		serviceUser := service.ServiceUser{Db: db}
		if err = serviceUser.GetById(accessToken); err != nil {
			unauthorized := model.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid Access Token",
				Data:       nil,
			}
			json.NewEncoder(w).Encode(unauthorized)
			return
		}
		fmt.Println(accessToken, "--------- ACCESS TOKEN DARI MIDDLEWARE")
		// accessToken := model.Todo{}
		// json.NewDecoder(r.Body).Decode(&accessToken)
		// fmt.Println(accessToken.UserId, "--------- ACCESS TOKEN DARI MIDDLEWARE")
		next.ServeHTTP(w, r)
	})
}
