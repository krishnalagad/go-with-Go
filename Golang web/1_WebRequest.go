package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const url = "http://bit.ly/3x2D8oj"

func main() {
	fmt.Println("Golang web request")

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

	// dumping the response data in the file.
	file, err := os.Create("./response.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	len, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("\nData dumped successfully : ", len)
}
