package main

import "fmt"

type pessoa struct {
	nome             string
	sobrenome        string
	saboresfavoritos []string
}

func main() {

	pessoa1 := pessoa{
		nome:             "Gabriel",
		sobrenome:        "Gradinar",
		saboresfavoritos: []string{"Maracuja", "Limão", "Pistache"},
	}

	pessoa2 := pessoa{
		nome:             "Julia",
		sobrenome:        "Sofreio",
		saboresfavoritos: []string{"Pistache", "Pistache", "Pistache"},
	}

	fmt.Printf("Pessoa: %v, Sobrenome: %v \n", pessoa1.nome, pessoa1.sobrenome)

	for _, sabor := range pessoa1.saboresfavoritos {
		fmt.Printf("Sabor: %v \n", sabor)
	}

	fmt.Printf("Pessoa: %v, Sobrenome: %v \n", pessoa2.nome, pessoa2.sobrenome)

	for _, sabor := range pessoa2.saboresfavoritos {
		fmt.Printf("Sabor: %v \n", sabor)
	}

}
