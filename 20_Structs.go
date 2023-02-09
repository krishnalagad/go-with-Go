package main

import "fmt"

type Point struct {
	x int32
	y int32
}

func main() {
	p1 := Point{1, 2}
	p2 := Point{x: 3}
	fmt.Println(p1, p2)
}
