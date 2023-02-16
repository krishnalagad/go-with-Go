package main

import "fmt"

type Point struct {
	x int32
	y int32
}

type Circle struct {
	radius float32
	center Point
}

func main() {
	p1 := Point{1, 2}
	p2 := Point{x: 3}
	fmt.Println(p1, p2)

	c1 := Circle{4.56, p1}
	fmt.Println(c1)
}
