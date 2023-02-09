package main

import "fmt"

func test() {
	fmt.Println("Test function !!")
}

func getLength(str string) {
	fmt.Println(len(str))
}

func getSum(a, b int) { // (a int, b int)
	fmt.Println(a + b)
}

func mulName(name string, no int) {
	for i := 0; i < no; i++ {
		fmt.Printf("%s ", name)
	}
}

func square(no int) int {
	return no * no
}

func mulReturns(a, b int) (r1 int, r2 int) {
	defer fmt.Println("Hello")
	r1 = a + b
	r2 = a - b
	fmt.Println("Before return")
	return
}

func main() {
	test()
	getLength("Krishna")
	getSum(12, 45)
	mulName("Krishna", 10)
	ans1, ans2 := mulReturns(5, 3)
	fmt.Println(ans1, ans2)
	fmt.Println(square(56))
}
