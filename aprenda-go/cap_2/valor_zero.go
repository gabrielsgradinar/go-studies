package main

import "fmt"

var a int     // 0
var b float64 // 0.0
var c string  // string vazia
var d bool    // false

func main() {
	fmt.Printf("%v, %T \n", a, a)
	fmt.Printf("%v, %T \n", b, b)
	fmt.Printf("%v, %T \n", c, c)
	fmt.Printf("%v, %T", d, d)
}
