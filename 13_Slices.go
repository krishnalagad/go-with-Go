package main

import "fmt"

func main() {
	a := [5]int{3, 4, 5, 6, 7}
	b := a[:]
	fmt.Println(b, len(b), cap(b))

	c := []int{1, 2, 3, 4, 5}
	d := append(c, 6)
	fmt.Println(d, len(d), cap(d))

	// using make() method we can create slice
	e := make([]int, 5)
	fmt.Printf("%T", e)
	fmt.Println(e)
}
