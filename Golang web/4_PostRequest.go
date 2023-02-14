package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func DoPostRequest(myurl string) {

	// making JSON data
	requestBody := strings.NewReader(`
		{
			"title": "Golang for intermidiate",
			"description": "Its Golang course for intermidiate learners."
		}
	`)

	// making JSON data
	data := url.Values{}
	data.Add("title", "Golang for Advanced")
	data.Add("description", "Its Golang course for advanced learners")

	response, err := http.Post(myurl, "application/json", requestBody)
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
