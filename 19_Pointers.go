package main

import "fmt"

func changeValue1(str *string) {
	*str = "changed"
}

func changeValue2(str string) {
	str = "changed"
}

func main() {
	toChange := "hello"
	changeValue1(&toChange)
	fmt.Println(toChange)
}
