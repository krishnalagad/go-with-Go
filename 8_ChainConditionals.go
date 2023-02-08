// #8
package main

import (
	"fmt"
)

func main() {
	val := true || false && true
	fmt.Printf("%v", val)

	res := (true || false) && !false
	fmt.Printf("%v", res)
}
