// #10
package main

import "fmt"

func main() {
	x := 1
	for x < 5 {
		fmt.Printf("%d ", x)
		x++
	}

	for x := 1; x <= 5; x++ {
		fmt.Printf("%d ", x)
	}
}
