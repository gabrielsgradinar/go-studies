package main

import "fmt"

func main() {
	mapa := map[string][]string{
		"gradinar_gabriel": {"games", "livros", "código"},
		"sofreio_julia":    {"maquiagem", "estudos", "harry potter"},
	}

	mapa["gradinar_igor"] = []string{
		"games", "design", "nft",
	}

	for chave, valor := range mapa {
		fmt.Println(chave)
		for _, v := range valor {
			fmt.Printf("\t %v \n", v)
		}
	}

}
