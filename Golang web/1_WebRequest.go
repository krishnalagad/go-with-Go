package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://bit.ly/3x2D8oj"

func main() {
	fmt.Println("Go web request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type is: %T", response)

	defer response.Body.Close() // user's responsibility to close the connection.

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(string(content))
}
