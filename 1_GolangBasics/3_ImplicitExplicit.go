// #3
package main

import "fmt"

func main() {
	fmt.Println("Implicit and explicit")

	var name = "Krishna Lagad" // implicit datatype
	fmt.Println(name)

	var mob = 9834765434
	fmt.Println(mob)
	fmt.Printf("%T", mob)

	city := "Pune" // another way of declaring variables.
	fmt.Println(city)
	fmt.Printf("%T", city)

	fmt.Println()

	fmt.Println("Default values of variables...")

	var number uint
	fmt.Println(number)

	var country string
	fmt.Println(country)

	var isTrue bool
	fmt.Println(isTrue)

	var percent float64
	fmt.Println(percent)
}
