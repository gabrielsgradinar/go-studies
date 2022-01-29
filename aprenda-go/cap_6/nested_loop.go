package main

import "fmt"

func main() {
	for horas := 0; horas <= 12; horas++ {
		fmt.Println("\n Hora:", horas)
		for x := 0; x <= 60; x++ {
			fmt.Print(" ", x)
		}
	}

	y := 0

	for y < 10 {
		fmt.Println(y)
		y++
	}
}
