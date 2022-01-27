package main

import "fmt"

func main() {
	igual := 5 == 5
	menor_igual := 2 <= 5
	menor := 10 < 5
	maior_igual := 2 >= 3
	maior := 3 > 5
	diferente := 2 > 1

	fmt.Println(igual)
	fmt.Println(menor_igual)
	fmt.Println(menor)
	fmt.Println(maior_igual)
	fmt.Println(maior)
	fmt.Println(diferente)
}
