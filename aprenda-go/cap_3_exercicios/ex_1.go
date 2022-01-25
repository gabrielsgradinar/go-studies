package main

import "fmt"

func main() {

	x := 42
	y := "James Bond"
	z := true

	fmt.Println("única declaração")
	fmt.Printf("X: %v \n", x)
	fmt.Printf("Y: %v \n", y)
	fmt.Printf("Z: %v \n", z)

	fmt.Println("múltiplas declarações")
	fmt.Println(x, y, z)
}
