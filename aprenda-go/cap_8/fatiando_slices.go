package main

import "fmt"

var x [5]int // declaração do array

func main() {
	sabores := []string{"portuguesa", "quatro_queijos", "mozzarela", "chocolate"}

	fatia := sabores[:]
	fmt.Println(fatia)

	// for i := 0; i < len(sabores); i++ {
	// 	fmt.Print(" ", sabores[i])
	// }

	sabores = append(sabores[:1], sabores[2:]...)

	fmt.Println(sabores)
}
