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

// controller
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

	// get movie data from request
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
	json.NewEncoder(res).Encode(movie)
	return
}

func UpdateMovie(res http.ResponseWriter, req *http.Request) {
	fmt.Println("PUT API - UpdateMovie")
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req) // getting all parameters

	for index, movie := range movies {
		if movie.MovieId == params["id"] { // getting query string parameter with key as 'id'
			movies = append(movies[:index], movies[index+1:]...) // deleting movie with 'id'
			var movie Movie
			_ = json.NewDecoder(req.Body).Decode(&movie) // fetching data from request body into variable.
			movie.MovieId = params["id"]                 // setting id of existing movie
			movies = append(movies, movie)               // adding updated data into fake DB.
			json.NewEncoder(res).Encode(movies)
			return
		}
	}
	json.NewEncoder(res).Encode("No data with given ID")
	return
}

func GetMovie(res http.ResponseWriter, req *http.Request) {
	fmt.Println("GET API - Get Movie")
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for _, movie := range movies {
		if movie.MovieId == params["id"] { // compare id from query string to the id in fakr db
			json.NewEncoder(res).Encode(movie) // if found, send movie in responce
			return
		}
	}
	json.NewEncoder(res).Encode("No movie found with id " + params["id"])
	return
}

func GetAllMovies(res http.ResponseWriter, req *http.Request) {
	fmt.Println("GET API - Get all movies")
	res.Header().Set("Content-Type", "application/json")

	// check if db is nil
	if movies == nil {
		json.NewEncoder(res).Encode("No data available in DB.")
		return
	}

	// send all data in response
	json.NewEncoder(res).Encode(movies)
	return
}

func DeleteMovie(res http.ResponseWriter, req *http.Request) {
	fmt.Println("DELETE API - Delete movie")
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	var responce string

	for index, movie := range movies {
		if movie.MovieId == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			responce = "Movie deleted successfully."
			break
		} else {
			responce = "Movie not found with id: " + params["id"]
		}
	}
	json.NewEncoder(res).Encode(responce)
	return
}

func main() {

	fmt.Println("Project started - Movie review API's")

	r := mux.NewRouter()

	// routing
	r.HandleFunc("/", ServeHome).Methods("GET", "POST")
	r.HandleFunc("/movie", AddMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", GetMovie).Methods("GET")
	r.HandleFunc("/movies", GetAllMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", DeleteMovie).Methods("DELETE")

	// setting port
	log.Fatal(http.ListenAndServe(":4000", r))

}
