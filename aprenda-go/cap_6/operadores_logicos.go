package main

import "fmt"

func main() {
	x := 2

	if x == 2 || x == 3 {
		fmt.Println("é 2 ou 3")
	}

	if x <= 2 && x < 3 {
		fmt.Println("menor ou igual a 2 e menor que 3 ")
	}
}
