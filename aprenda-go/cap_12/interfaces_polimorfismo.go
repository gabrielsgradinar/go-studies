package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type dev struct {
	pessoa
	linhasdecodigo int
	salario        float64
}

type biomedica struct {
	pessoa
	especializacao string
	salario        float64
}

func (x dev) seapresenta() {
	fmt.Println("Meu nome é", x.nome, "e ja escrevi", x.linhasdecodigo, "linhas de código!")
}

func (x biomedica) seapresenta() {
	fmt.Println("Meu nome é", x.nome, "sou especializada em", x.especializacao)
}

type gente interface {
	seapresenta()
}

func serhumano(g gente) {
	g.seapresenta()
}

func main() {
	developer := dev{pessoa{"Gabriel", "Gradinar", 23}, 5000, 10.500}
	bio := biomedica{pessoa{"Julia", "Sofreio", 21}, "Vigilancia Sanitaria", 10.500}

	serhumano(developer)
	serhumano(bio)
}
