package main

import (
	"fmt"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Go web request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type is: %T", response)

	response.Body.Close()	// user's responsibility to close the connection.
}