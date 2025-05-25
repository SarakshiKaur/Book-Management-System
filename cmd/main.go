package main

import (
	"fmt"
	"net/http"

	"github.com/SarakshiKaur/Book-Management-System/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize mux
	r := mux.NewRouter()

	// registering routes here from routes folder
	routes.RegisterRoutes(r)

	if err := http.ListenAndServe(":3000", r); err != nil {
		fmt.Printf("err : %v", err)
	}
}
