package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"apple":  3,
		"guava":  4,
		"orange": 7,
	}
	fmt.Println(mp)
	// mp1 := map[string]int{

	// }
	mp["apple"] = 10
	mp["banana"] = 6
	fmt.Println(mp)

	delete(mp, "guava")
	fmt.Println(mp)
}
