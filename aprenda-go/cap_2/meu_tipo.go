package main

import "fmt"

type dinheiro float64

var b dinheiro = 10.5

func main() {

	a := 10.5
	fmt.Printf("%v, %T \n", b, b)
	fmt.Printf("%v, %T", a, a)

	a = float64(b)

}
