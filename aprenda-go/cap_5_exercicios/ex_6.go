package main

import "fmt"

const (
	_ = iota + 2022
	a
	b
	c
	d
)

func main() {
	fmt.Printf("%v - %v - %v - %v", a, b, c, d)

	fmt.Printf("%#x", 31337)
}
