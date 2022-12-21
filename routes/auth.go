package routes

import (
	"MisterAladin/handlers"
	"MisterAladin/pkg/middleware"
	"MisterAladin/pkg/mysql"
	"MisterAladin/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(userRepository)

	r.HandleFunc("/Login", h.Login).Methods("POST")
	r.HandleFunc("/GetUsers/{id}", h.GetUsers).Methods("GET")
	r.HandleFunc("/check-auth", middleware.Auth(h.CheckAuth)).Methods("GET")
}
