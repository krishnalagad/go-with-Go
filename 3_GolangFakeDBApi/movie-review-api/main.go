package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Model for Moview
type Movie struct {
	MovieId         string   `json:"movieid"`
	MovieName       string   `json:"moviename"`
	MovieBudget     int32    `json:"budget"`
	MovieCollection int32    `json:"collection"`
	Reviews         []Review `json:"reviews"`
}

// Model for Review
type Review struct {
	ReviewerName  string `json:"name"`
	ReviewContent string `json:"review"`
	RatingStars   int8   `json:"rating"`
}

// fake DB
var movies []Movie

// controller
func serveHome(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("<h1>Welcome to Movie review API by Krishna</h1>"))
}

func main() {

	fmt.Println("Movie review API's")

	r := mux.NewRouter()

	// routing
	r.HandleFunc("/", serveHome).Methods("GET", "POST")

	// setting port
	log.Fatal(http.ListenAndServe(":4000", r))

}
