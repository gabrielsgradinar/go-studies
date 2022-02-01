package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	tracaoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {

	caminhao := caminhonete{veiculo{2, "preto"}, true}

	sedanzin := sedan{veiculo{4, "prata"}, false}

	fmt.Println(caminhao)
	fmt.Println(sedanzin)

	fmt.Println(sedanzin.portas)
	fmt.Println(caminhao.tracaoNasQuatro)

}
