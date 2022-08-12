package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/users", h.FindUsers).Methods("GET")
	r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
	// Create "/user" route with CreateUser handler here ...
	r.HandleFunc("/user", h.CreateUser).Methods("POST")
	//todo kita buatkan endpoint-nya dengan function h.CreateUser dan method-nya adalah POST
}
