package routes

import (
	"github.com/gorilla/mux"
	"github.com/krishnalagad/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/book", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/api/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/api/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
