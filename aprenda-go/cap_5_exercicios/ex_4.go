package main

import "fmt"

func main() {
	n := 10
	fmt.Printf("%d - %b - %#x \n", n, n, n)

	x := n << 1
	fmt.Printf("%d - %b - %#x", x, x, x)

}
