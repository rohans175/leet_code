package main

import (
	"fmt"
)

func main() {

	a := []int{1, 2, 3, 4, 5, 6}
	for index, each_a := range a {
		fmt.Println("index", index, "eachA", each_a)
	}
}
