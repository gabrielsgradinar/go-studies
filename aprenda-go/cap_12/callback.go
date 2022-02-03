package main

import "fmt"

func main() {
	total := somenteImpares(soma, []int{50, 51, 52, 5, 7, 8, 9, 10, 68, 59, 70}...)

	fmt.Println(total)
}

func soma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func somenteImpares(f func(x ...int) int, y ...int) int {
	var slice []int

	for _, v := range y {
		if v%2 != 0 {
			slice = append(slice, v)
		}
	}

	total := f(slice...)

	return total
}
