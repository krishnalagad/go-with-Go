// #6
package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 float64 = 90
	var num2 int = 8
	ans := num1 / float64(num2)
	fmt.Printf("%.3g", math.Sqrt(ans))

	var num3 float64 = 90
	var num4 int = 8
	res := num3 / float64(num4)
	fmt.Printf("\n%.3g", math.Abs(res))
}
