package main

import "fmt"

func main() {

	valor := soma(2, 2)

	fmt.Println(valor)

	fmt.Println(basica())

	resultado, quantidade := multiplicacao(2, 2, 2)

	fmt.Println(resultado, quantidade)
}

// n1 int, n2 int
func soma(n1, n2 int) int {
	return n1 + n2
}

func basica() string {
	return "Olá Mundo!"
}

func multiplicacao(x ...int) (int, int) {
	mult := 1

	for _, v := range x {
		mult = mult * v
	}

	return mult, len(x)
}
