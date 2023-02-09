package main

import "fmt"

func main() {
	// Mutable: change in y will change in x and vice versa.
	var x[] int = []int { 1, 2, 3, 4}
	y := x
	y[0] = 100
	fmt.Println(x, y)

	var mp map[string]int = map[string]int{
		"apple": 3,
	}
	raw := mp
	raw["banana"] = 5
	fmt.Println(mp, raw)

	// Immutable: change in one cannot change other
	var arr [2]int = [2]int{2,3}
	arr2 := arr
	arr2[0] = 100
	fmt.Println(arr, arr2)
}