package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	for _, value := range array {
		fmt.Print(" ", value)
	}

	fmt.Printf("\n %T", array)
}
