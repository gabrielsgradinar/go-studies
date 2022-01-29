package main

import "fmt"

func main() {

	idade := 70

	if idade >= 18 && idade < 65 {
		fmt.Println("maior de idade")
	} else if idade > 65 {
		fmt.Println("idoso")
	} else {
		fmt.Println("menor de idade")
	}
}
