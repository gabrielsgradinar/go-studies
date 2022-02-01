package main

import "fmt"

func main() {

	x := struct {
		listaTelefonica map[string]int
		operadoras      []string
	}{
		map[string]int{
			"joao":  11976541234,
			"maria": 11980651234,
		},
		[]string{"vivo", "tim", "claro"},
	}

	for _, value := range x.listaTelefonica {
		fmt.Println(value)
	}

	for _, value := range x.operadoras {
		fmt.Println(value)
	}

}
