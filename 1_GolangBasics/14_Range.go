package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 5, 66, 7, 65, 3, 7}

	for i, element := range arr {
		for j := i + 1; j < len(arr); j++ {
			element2 := arr[j]
			if element2 == element {
				fmt.Println(element)
			}
		}
	}
}
