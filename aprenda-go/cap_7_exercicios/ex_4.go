package main

import "fmt"

func main() {
	anoAtual := 2022
	nascimento := 1998
	for {
		fmt.Println(nascimento)
		nascimento++
		if nascimento > anoAtual {
			break
		}
	}
}
