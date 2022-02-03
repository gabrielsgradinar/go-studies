package main

import "fmt"

func main() {
	funcao := retornafuncao()

	fmt.Println(funcao(5))

}

func retornafuncao() func(int) int {
	return func(i int) int {
		return i * 10
	}
}
