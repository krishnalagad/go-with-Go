package main

import "fmt"

func test1(x int){
	fmt.Println("Hello", x)
}

// taking function as a parameter
func test3(myFunc func(x int) int){
	res := myFunc(378)
	fmt.Println(res)
}

func main() {
	x := test1
	x(34)

	// function within function
	fun := func(x int) int{
		return x * -1
	}
	fun(56)

	// passing argument to function while defining it.
	test2 := func(x int) int{
		return x * -1
	}(78)
	fmt.Println(test2)

	// passing function as a argument to a function.
	test3(fun)
}