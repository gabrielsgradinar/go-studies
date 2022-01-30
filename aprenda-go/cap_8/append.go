package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := []int{9, 10, 11, 12}
	fmt.Printf("Slice original : %v \n", x)

	x = append(x, 5, 6, 7, 8)
	fmt.Printf("Slice com append de ints : %v \n", x)

	x = append(x, y...)
	fmt.Printf("Slice com append de um slice de ints : %v \n", x)

}
