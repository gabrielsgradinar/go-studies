package main

import "fmt"

type topper int

var x topper

func main() {

	fmt.Println(x)
	fmt.Printf("%T \n", x)

	x = 42
	fmt.Println(x)
}
