package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJson() {
	courses := []course{
		{"Golang Bootcamp", 349, "YouTube Live", "abc123", []string{"golang", "api", "web"}},
		{"MERN Bootcamp", 599, "MS Teams", "aqo123", []string{"react", "js", "mongodb", "node", "express"}},
		{"Spring boot Bootcamp", 999, "Google Meet", "asd345", nil},
	}

	// package this data as JSON data
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonData := []byte(`
		{
			"coursename": "Golang Bootcamp",
			"Price": 349,
			"website": "YouTube Live",
			"tags": ["golang","api","web"]
        }
	`)
	var courses course

	isValid := json.Valid(jsonData)

	if isValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonData, &courses)
		fmt.Printf("%#v\n", courses)
	} else {
		fmt.Println("JSON IS NOT VALID")
	}

	// some cases where you just want to add data to key value
	var myData map[string]interface{}
	json.Unmarshal(jsonData, &myData)
	fmt.Printf("%#v\n", myData)

	for k, v := range myData {
		fmt.Printf("Key is %v and value is %v and Type is %T\n", k, v, v)
	}
}

func main() {
	fmt.Println("Golang JSON")
	EncodeJson()
	DecodeJson()
}
