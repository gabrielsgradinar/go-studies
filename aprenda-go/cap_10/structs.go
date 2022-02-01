package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

func main() {

	pessoa1 := pessoa{
		nome:  "Gabriel",
		idade: 30,
	}

	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "Maria",
			idade: 24,
		},
		titulo:  "Jornalista",
		salario: 3000,
	}

	pessoa3 := pessoa{"Julia", 22}
	pessoa4 := profissional{pessoa{"Joao", 50}, "Politico", 20000}

	fmt.Println(pessoa1.nome)
	fmt.Println(pessoa2.nome)
	fmt.Println(pessoa3.nome)
	fmt.Println(pessoa4)

	// struct anonimo

	x := struct {
		nome  string
		idade int
	}{
		nome:  "Igor",
		idade: 20,
	}

	fmt.Println(x)
}
