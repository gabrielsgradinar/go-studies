package main

import "fmt"

func main() {
	x := 0
	for x <= 20 {
		x++
		if x%2 != 0 { // verificar se o numero não é par
			continue
		}
		fmt.Println(x)
	}
}
