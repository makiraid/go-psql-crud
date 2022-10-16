package router

import (
	"go-psql/middleware"

	"github.com/gorilla/mux"
)

func User(mux *mux.Router) {
	mux.HandleFunc("/api/user/{id}", middleware.GetUser).Methods("GET", "OPTIONS")
	mux.HandleFunc("/api/user", middleware.GetAllUser).Methods("GET", "OPTIONS")
	mux.HandleFunc("/api/createuser", middleware.CreateUser).Methods("POST", "OPTIONS")
	mux.HandleFunc("/api/user/{id}", middleware.UpdateUser).Methods("PUT", "OPTIONS")
	mux.HandleFunc("/api/user/{id}", middleware.DeleteUser).Methods("DELETE", "OPTIONS")
}
