package main

import "fmt"

func main() {
	//         0   1   2   3   4   5   6   7   8   9
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// [42, 43, 44, 48, 49, 50, 51]
	novo_slice := append(x[:3], x[6:]...)

	fmt.Println(novo_slice)
}
