package main

import "fmt"

func main() {
	estados_br := make([]string, 26, 50)

	estados_br = []string{"Acre", "Alagoas", "Amapá", "Amazonas",
		"Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso",
		"Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco",
		"Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia",
		"Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins",
	}

	fmt.Println(len(estados_br), cap(estados_br))

	for i := 0; i < len(estados_br); i++ {
		fmt.Print(" ", estados_br[i])
	}

}
