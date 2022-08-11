package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/gorilla/mux"
)

// Create UserRoutes function here ...
func UserRoutes(r *mux.Router) {
	//ctt Masih sama sebagaimana sebelumnya, kita menggunakan mux.Router. Kemudian dilanjutkan dengan:

	userRepository := repositories.RepositoryUser(mysql.DB)
	//ctt Pendeklarasian userRepository-nya yg mana userRepository itu membutuhkan connection.
	//ctt Karena, kita melakukan 'query'-nya di dlm repositories-nya. Maka dari itu, kita isikan connection-nya (mysql.DB) yg kemudian akan dikirimkan melalui HandlerUser di bawah.

	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/users", h.FindUsers).Methods("GET")
	r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
}
