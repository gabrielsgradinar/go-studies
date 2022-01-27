package main

import "fmt"

func main() {
	s := "Hello World"

	for _, s := range s {
		fmt.Printf("%v - %T - %#U - %#x\n ", s, s, s, s)
	}
}
