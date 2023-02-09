package main

import "fmt"

func test1(x int){
	fmt.Println("Hello", x)
}

func main() {
	x := test1
	x(34)
}