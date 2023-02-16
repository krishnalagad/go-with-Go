package main

import "net/http"

// import (
// 	"encoding/json"
// 	"fmt"
// 	"math/rand"
// 	"net/http"
// 	"strconv"
// 	"time"
// )

// func ServeHome(res http.ResponseWriter, req *http.Request) {
// 	res.Write([]byte("<h1>Welcome to Movie review API by Krishna</h1>"))
// }

// API to add movie
// func AddMovie(res http.ResponseWriter, req *http.Request) {
// 	fmt.Println("POST API - Add movie")
// 	res.Header().Set("Content-Type", "application/json")
// 	var movie Movie

// 	// data validation before submission.
// 	// checking if body is empty
// 	if req.Body == nil {
// 		json.NewEncoder(res).Encode("Please send some data")
// 	}

// 	_ = json.NewDecoder(req.Body).Decode(&movie)

// 	// checking if data is empty
// 	if movie.isNameEmpty() { // check if movie name is empty
// 		json.NewEncoder(res).Encode("Movie name is empty.")
// 		return
// 	} else if movie.isBudgetEmpty() { // check is movie budget is empty
// 		json.NewEncoder(res).Encode("Movie budget is empty")
// 		return
// 	} else if movie.isCollectionEmpty() { // check if movie collection is empty
// 		json.NewEncoder(res).Encode("Movie collection is empty")
// 		return
// 	}

// 	// generating unique id for movie and converting it into string
// 	rand.Seed(time.Now().UnixNano())
// 	movie.MovieId = strconv.Itoa(rand.Intn(100)) // assigning new generated number to movie id.

// 	// adding new movie into fake DB.
// 	movies = append(movies, movie)
// 	json.NewEncoder(res).Encode(movie)
// 	return
// }

func updateMovie(res http.ResponseWriter, req *http.Request) {

}

func getMovie(res http.ResponseWriter, req *http.Request) {

}

func getAllMovies(res http.ResponseWriter, req *http.Request) {

}

func deleteMovie(res http.ResponseWriter, req *http.Request) {

}
