package main

import (
	"fmt"
)

func main() {
	amigos := map[string]int{
		"joao":  123456,
		"maria": 9786134,
	}

	fmt.Println(amigos["joao"])
	fmt.Println(amigos["maria"])

	amigos["gopher"] = 45678

	fmt.Println(amigos)

	// comma ok idiom
	if sera, ok := amigos["fantasma"]; !ok {
		fmt.Println("não tem!")
	} else {
		fmt.Println(sera)
	}
}
