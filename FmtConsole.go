// #4
package main

import "fmt"

func main() {
	fmt.Printf("KRISHNA %T %v", 10, 10)
	var result string = fmt.Sprintf("KRISHNA %T %v", 10, 10)
	fmt.Println(result)

	fmt.Println()

	fmt.Printf("Number binary: %b ", 345)
	fmt.Printf("Number octal: %o ", 345)
	fmt.Printf("Number decimal: %d ", 345)
	fmt.Printf("Number hexodecimal: %x ", 3435)

	fmt.Println()

	fmt.Printf("String: %s\n ", "Krishna")
	fmt.Printf("String: %q ", "Krishna")
}