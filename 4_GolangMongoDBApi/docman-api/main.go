package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/krishnalagad/docman-api/router"
)

func main() {

	fmt.Println("Project started")
	fmt.Println("Server is getting started...")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))

	fmt.Println("Listening to port 4000")

}
