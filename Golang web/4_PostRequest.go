package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func DoPostRequest(url string) {

	// making JSON data
	requestBody := strings.NewReader(`
		{
			"title": "Golang for intermidiate",
			"description": "Its Golang course for intermidiate learners."
		}
	`)

	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func main() {
	DoPostRequest("http://localhost:8080/api/courses/")
}
