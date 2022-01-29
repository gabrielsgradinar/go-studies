package main

import "fmt"

var esporteFavorito string

func main() {

	esporteFavorito = "Futebol"

	switch esporteFavorito {
	case "Futebol":
		fmt.Println("Gosta de jogar bola")
	case "Natação":
		fmt.Println("Gosta de nadar")
	default:
		fmt.Println("Gosta de fazer nada")
	}
}
