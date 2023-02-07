// #5
package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type something: ")
	scanner.Scan()
	input := scanner.Text()

	fmt.Printf("You typed: %q", input)
}
