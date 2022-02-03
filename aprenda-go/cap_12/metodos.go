package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) bomdia() {
	fmt.Println(p.nome, "diz bom dia !")
}

func main() {

	gabriel := pessoa{"Gabriel", 23}

	gabriel.bomdia()

}
