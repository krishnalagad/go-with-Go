// #5
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type something: ")
	scanner.Scan()
	input := scanner.Text()

	fmt.Printf("You typed: %q", input)
	fmt.Println()

	fmt.Printf("Type the year you were born: ")
	scanner.Scan()                                      // By default Scan() takes input in string format.
	year, _ := strconv.ParseInt(scanner.Text(), 10, 64) // type conversion from string to number
	fmt.Printf("Your current age is %d years", 2023-year)
}
