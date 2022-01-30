package main

import "fmt"

func main() {
	x := make([]int, 5, 10)

	x[0], x[1], x[2], x[3], x[4] = 1, 2, 3, 4, 5

	x = append(x, 6)
	x = append(x, 7)
	x = append(x, 8)
	x = append(x, 9)
	x = append(x, 10)

	fmt.Println(x, len(x), cap(x))

	x = append(x, 11)

	fmt.Println(x, len(x), cap(x))

}
