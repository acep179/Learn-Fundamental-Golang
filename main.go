package main

import (
	"dumbmerch/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	//. Utk pendeklarasian routernya masih sama, yaitu dengan menggunakan mux.NewRouter(). Kemudain, kita panggil RouteInit dari package routes dengan mengetikkan routes.RouteInit sebagaiman adi bawah ini

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	// . Jika kita lihat ke file routes.go. Maka dapat dilihat bahwa penulisan function RouteInit-nya adalah sebagaimana berikut: func RouteInit(r *mux.Router). Jika dilihat lagi maka terlihat bahwa function RouteInit membutuhkan parameter. Adapun parameternya adalah mux.Router. Maka dari itu, dalam pemanggilan function RouteInit di sini pun diberikan parameternya. yaitu, segala sesuatu yg berada di dlm tanda kurung RouteInit().

	//ctt Jika dijabarkan tanpa memisahkan file dan folder:
	// mux.NewRouter().PathPrefix("/api/v1").Subrouter().HandleFunc("/todos", handlers.FindTodos).Methods("GET")
	// mux.NewRouter().PathPrefix("/api/v1").Subrouter().HandleFunc("/todo", handlers.GetTodo).Methods("GET")

	//todo Untuk pemanggilan function dari package lain, kita harus mengetikkan nama package-nya terlebih dahulu barulah nama function-nya

	//todo Perhatikan penamaan package dan function, pastikan sudah sama / sudah sesuai dengan sumbernya!

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
