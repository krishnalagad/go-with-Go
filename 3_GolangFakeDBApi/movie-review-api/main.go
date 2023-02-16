package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for Moview
type Movie struct {
	MovieId         string    `json:"movieid"`
	MovieName       string    `json:"moviename"`
	MovieBudget     int32     `json:"budget"`
	MovieCollection int32     `json:"collection"`
	Reviews         *[]Review `json:"reviews"`
}

// Model for Review
type Review struct {
	ReviewerName  string `json:"name"`
	ReviewContent string `json:"review"`
	RatingStars   int8   `json:"rating"`
}

// fake DB
var movies []Movie

// helper methods
func (m *Movie) isNameEmpty() bool {
	return m.MovieName == ""
}
func (m *Movie) isBudgetEmpty() bool {
	return m.MovieBudget == 0
}
func (m *Movie) isCollectionEmpty() bool {
	return m.MovieName == ""
}

func ServeHome(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("<h1>Welcome to Movie review API by Krishna</h1>"))
}

func AddMovie(res http.ResponseWriter, req *http.Request) {
	fmt.Println("POST API - Add movie")
	res.Header().Set("Content-Type", "application/json")
	var movie Movie

	// data validation before submission.
	// checking if body is empty
	if req.Body == nil {
		json.NewEncoder(res).Encode("Please send some data")
	}

	_ = json.NewDecoder(req.Body).Decode(&movie)

	// checking if data is empty
	if movie.isNameEmpty() { // check if movie name is empty
		json.NewEncoder(res).Encode("Movie name is empty.")
		return
	} else if movie.isBudgetEmpty() { // check is movie budget is empty
		json.NewEncoder(res).Encode("Movie budget is empty")
		return
	} else if movie.isCollectionEmpty() { // check if movie collection is empty
		json.NewEncoder(res).Encode("Movie collection is empty")
		return
	}

	// generating unique id for movie and converting it into string
	rand.Seed(time.Now().UnixNano())
	movie.MovieId = strconv.Itoa(rand.Intn(100)) // assigning new generated number to movie id.

	// adding new movie into fake DB.
	movies = append(movies, movie)
	json.NewEncoder(res).Encode(movies)
	return
}

func main() {

	fmt.Println("Movie review API's")

	r := mux.NewRouter()

	// routing
	r.HandleFunc("/", ServeHome).Methods("GET", "POST")
	r.HandleFunc("/movie", AddMovie).Methods("POST")

	// setting port
	log.Fatal(http.ListenAndServe(":4000", r))

}
