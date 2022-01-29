package main

import "fmt"

var x [5]int // declaração do array

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	slice := []int{1, 2, 3, 4, 5} // declaração do slice
	fmt.Println(slice)

	slice2 := append(slice, 6)
	fmt.Println(slice2)

	nomes := []string{"Gabriel", "Julia", "Igor", "Tiago"}

	// for indice, valor := range nomes {
	// 	fmt.Println("No índice", indice, "temos o valor", valor)
	// }

	nomes = append(nomes, "Nicolly")

	for indice, valor := range nomes {
		fmt.Println("No índice", indice, "temos o valor", valor)
	}
}
