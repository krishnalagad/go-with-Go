package main

import "fmt"

func test(){
	fmt.Println("Test function !!")
}

func getLength(str string) {
	fmt.Println(len(str))
}

func main() {
	test()
	getLength("Krishna")
}