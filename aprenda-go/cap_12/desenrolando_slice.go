package main

import "fmt"

func main() {

	valores := []int{2, 2, 2}

	//parametro
	resultado := multiplicacao(valores...)

	fmt.Println(resultado)

	zero := multiplicacao()
	fmt.Println(zero)

}

// argumento
func multiplicacao(x ...int) int {
	mult := 1
	for _, v := range x {
		mult = mult * v
	}
	return mult
}
