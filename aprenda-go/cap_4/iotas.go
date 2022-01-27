package main

import "fmt"

// constantes não tipadas e com valores inteiros
const (
	_ = iota * 2
	y
	z
	a
	b
)

func main() {
	fmt.Println(y, z, a, b)
}
