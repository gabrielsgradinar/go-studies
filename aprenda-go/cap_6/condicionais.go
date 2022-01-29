package main

import "fmt"

func main() {
	x := 10
	if x < 100 {
		fmt.Println(x)
	}

	if !(x > 100) { // ! -> inverte a condição
		fmt.Println(x)
	}

	if y := 20; y > 50 {
		fmt.Println("y maior que 50")
	} else if y <= 20 {
		fmt.Println("y menor ou igual a 20")
	} else {
		fmt.Println("y maior que 20 e menor que 50")
	}
}
