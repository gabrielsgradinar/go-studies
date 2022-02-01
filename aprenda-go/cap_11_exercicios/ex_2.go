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

	pessoas := map[string]pessoa{
		pessoa1.sobrenome: pessoa1,
		pessoa2.sobrenome: pessoa2,
	}

	for sobrenome, pessoa := range pessoas {
		fmt.Println(pessoa.nome, sobrenome)
		for _, sabor := range pessoa.saboresfavoritos {
			fmt.Printf("\t %v \n", sabor)
		}
	}

}
