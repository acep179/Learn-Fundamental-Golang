package routes

// Import gorilla/mux package here ...
import (
	"github.com/gorilla/mux"
)

// Create RouteInit function and Call TodoRoutes function here ...

func RouteInit(r *mux.Router) {
	TodoRoutes(r)
}

//. Route init adalah utk menginisialisasi route-nya. Route init itu diambil dari mux.Router. Adapun mux.Router akan menerima nilai routing yg kita kirimkan dari main.go (PathPrefix). Sehingga, seluruh routing kita akan kita grup ke dalam /api/v1 diikuti dengan /todos, /users, /products, dll.

//. Penulisan function RouteInit harus diawali dengan huruf kapital.

//. Karena Function RouteInit akan dimasukkan atau digunakan ke dalam file main.go. Maka, kita harus memastikan bahwa func RouteInit merupakan sebuah func global. Adapun cara memastikannya adalah, kita harus menuliskan kata awalnya dengan menggunakan huruf kapital: RouteInit, bukan huruf kecil: routeInit

//. Jika RouteInit diawali dengan huruf kecil, maka golang akan menganggap bahwa routeInit adalah sebuah function local. Sehingga, routeInit tdk akan dapat digunakakn di package lain.
