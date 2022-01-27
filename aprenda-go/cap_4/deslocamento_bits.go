package main

import "fmt"

// constantes não tipadas e com valores inteiros

func main() {
	x := 2
	y := x >> 1
	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", y)
}
