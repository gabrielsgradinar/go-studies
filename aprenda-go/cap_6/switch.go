package main

import (
	"fmt"
)

var z interface{}

func main() {
	x := 5
	switch {
	case x <= 5:
		fmt.Println("menor ou igual a 5")
		fallthrough
	case x == 5:
		fmt.Println("igual a 5")
	case x > 5:
		fmt.Println("maior que 5")
	default:
		fmt.Println("eita")
	}

	// case composto
	y := 5
	switch y {
	case 2, 3:
		fmt.Println("2 ou 3")
	case 4, 5:
		fmt.Println("4 ou 5")
	}

	// chegando o tipo da variavel
	z = 123
	switch z.(type) {
	case int:
		fmt.Println("é um int")
	case bool:
		fmt.Println("é um bool")
	case string:
		fmt.Println("é uma string")
	case float64:
		fmt.Println("é um float")

	}

}
