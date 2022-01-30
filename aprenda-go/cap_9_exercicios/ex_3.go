package main

import "fmt"

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(array[:3])
	fmt.Println(array[4:])
	fmt.Println(array[1:7])
	fmt.Println(array[2:9])
	fmt.Println(array[2 : len(array)-1])
}
