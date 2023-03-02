package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krishnalagad/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	fmt.Println("Project started.")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":4000", r))
}
