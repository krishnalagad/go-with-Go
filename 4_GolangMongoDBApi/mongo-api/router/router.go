package router

import (
	"github.com/gorilla/mux"
	"github.com/krishnalagad/mongo-api/controller"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", controller.ServeHome).Methods("POST", "GET")
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkedAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", controller.DeleteAllMovie).Methods("DELETE")

	return router
}
