package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=jkadfhsd56"

func main() {
	fmt.Println("Golang URLs")
	fmt.Println(myurl)

	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Port())
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("\nType of quesry param: %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println(val)
	}

	// build url
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "localhost",
		Path:     "/learn",
		RawQuery: "user=krishna",
	}
	fmt.Println(partsOfUrl.String())
}
