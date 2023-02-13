package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println(string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("File Handling in Golang")

	content := "This content need to go in file - Krishna Lagad is my name."
	file, err := os.Create("./myfile.txt")

	checkNilError(err)
	len, err := io.WriteString(file, content)

	checkNilError(err)
	fmt.Println("Length: ", len)
	defer file.Close()
	readFile("./myfile.txt")
}
