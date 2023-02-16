package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func DoGetRequest(url string) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	// fmt.Println("Content length: ", response.ContentLength)

	var responseString strings.Builder
	databyte, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(databyte)

	fmt.Println("Byte count: ", byteCount)
	fmt.Println(responseString.String())
	// fmt.Println(string(databyte))
}

func main() {
	fmt.Println("Golang GET request")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type Id: ")
	scanner.Scan()
	input := scanner.Text()
	DoGetRequest("http://localhost:8080/api/courses/" + input)
}
