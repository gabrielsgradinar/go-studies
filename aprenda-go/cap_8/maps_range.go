package main

import "fmt"

func main() {
	amigos := map[string]int{
		"joao":  123456,
		"maria": 9786134,
	}

	for nome, telefone := range amigos {
		fmt.Println(nome, telefone)
	}

	delete(amigos, "joao")

	fmt.Println(amigos)
}
