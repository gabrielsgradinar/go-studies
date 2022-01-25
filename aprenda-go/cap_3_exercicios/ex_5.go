package main

import "fmt"

type topper int

var x topper

var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T \n", x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T \n", y)
}
