package main

import "fmt"

func main() {

	x := "Oi bom dia\n como vai ?\n \"topper\"." // string interpretada
	fmt.Println(x)

	y := `"Oi bom dia\n 
	
	como vai ?\n            \"topper\"."` // string literal
	fmt.Println(y) // printa tudo na tela mas com um espaço no final

	b := "Tudo bem ?"
	a := fmt.Sprint("Oi, bom dia!", b) // pega os valores passados e retorna uma string
	fmt.Print(a)                       // print mais sem um espaço no final

}
