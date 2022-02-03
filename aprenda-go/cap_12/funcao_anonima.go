package main

import "fmt"

func main() {
	x := 10
	y := 5
	func(x, y int) {
		fmt.Println(x + y)
	}(x, y)
}
