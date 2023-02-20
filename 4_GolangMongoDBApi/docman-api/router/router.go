package router

import (
	"github.com/gorilla/mux"
	"github.com/krishnalagad/docman-api/controller"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", controller.ServerHome).Methods("GET", "POST")
	router.HandleFunc("/api/doc", controller.CreateDocument).Methods("POST")
	router.HandleFunc("/api/doc/{id}", controller.UpdateOneDocument).Methods("PUT")
	router.HandleFunc("/api/doc/{id}", controller.GetOneDocument).Methods("GET")
	router.HandleFunc("/api/docs", controller.GetAllDocuments).Methods("GET")
	router.HandleFunc("/api/doc/{id}", controller.DeleteOneDocument).Methods("DELETE")
	router.HandleFunc("/api/docs", controller.DeleteAllDocuments).Methods("DELETE")

	return router
}
