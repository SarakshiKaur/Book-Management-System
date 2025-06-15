package routes

import (
	"github.com/SarakshiKaur/Book-Management-System/pkg/controllers"
	"github.com/gorilla/mux"
)

// this is how we register routes
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", controllers.HandleRoot).Methods("GET")
	r.HandleFunc("/books", controllers.getBooks).Methods("GET")
	r.HandleFunc("/book", controllers.createBook).Methods("GET")
	r.HandleFunc("/book/{id}", controllers.getBookById).Methods("GET")
}
