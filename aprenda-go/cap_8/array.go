package main

import "fmt"

var x [5]int // declaração do array

func main() {
	x[0] = 1
	x[1] = 10
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Printf("%T \n", x)
	fmt.Println(len(x))

	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)

}
