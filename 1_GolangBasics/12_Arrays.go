// #12
package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 12
	arr[4] = 21
	fmt.Println(arr, len(arr))

	arr1 := []int{12, 13, 14, 15, 16, 17, 18}
	fmt.Println(arr1, len(arr1))

	sum := 0
	for i := 0; i < len(arr1); i++ {
		sum += arr1[i]
	}
	fmt.Println(sum)

	// 2D arrays
	arr2D := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2D)

}
