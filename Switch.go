//#11
package main

import "fmt"

func main() {
	ans := 19
	switch ans {
	case 1:
		fmt.Println("One")
		break
	case 2:
		fmt.Println("Two")
		break
	default:
		fmt.Println("not a case")
	}
}
