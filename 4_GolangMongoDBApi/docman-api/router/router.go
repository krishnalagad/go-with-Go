package router

import (
	"github.com/gorilla/mux"
	"github.com/krishnalagad/docman-api/controller"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", controller.ServerHome).Methods("GET", "POST")
	router.HandleFunc("/api/doc", controller.CreateDocument).Methods("POST")
	// router.HandleFunc("/api/doc", controller.CreateMovie).Methods("PUT")
	// router.HandleFunc("/api/doc/{}", controller.CreateMovie).Methods("POST")

	return router
}
