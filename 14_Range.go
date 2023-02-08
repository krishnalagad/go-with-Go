package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 5, 66, 7, 65}

	for i, element := range arr {
		fmt.Printf("%d: %d\n", i, element)
	}
}
