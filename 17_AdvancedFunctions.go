package main

import "fmt"

func test1(x int){
	fmt.Println("Hello", x)
}

func main() {
	x := test1
	x(34)

	// function within function
	fun := func(x int){
		fmt.Println(x)
	}
	fun(56)

	// passing argument to function while defining it.
	test2 := func(x int) int{
		return x * -1
	}(78)
	fmt.Println(test2)
}