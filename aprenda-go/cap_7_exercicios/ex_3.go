package main

import "fmt"

func main() {
	anoAtual := 2022
	nascimento := 1998
	for nascimento <= anoAtual {
		fmt.Println(nascimento)
		nascimento++
	}
}
