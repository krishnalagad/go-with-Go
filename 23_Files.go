package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File Handling in Golang")

	content := "This content need to go in file - Krishna Lagad is my name."
	file, err := os.Create("./myfile.docx")

	if err != nil {
		panic(err)
	}
	len, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("Length: ", len)
	defer file.Close()
}
