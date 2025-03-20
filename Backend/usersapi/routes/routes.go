package routes

import (
	"net/http"

	"userapi/handlers"

	"github.com/gorilla/mux"
)

func NewRouter(userHandler *handlers.UserHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", userHandler.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users", userHandler.GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/{id:[0-9]+}", userHandler.GetUserByID).Methods(http.MethodGet)
	r.HandleFunc("/users/{id:[0-9]+}", userHandler.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/users/{id:[0-9]+}", userHandler.DeleteUser).Methods(http.MethodDelete)

	return r
}
