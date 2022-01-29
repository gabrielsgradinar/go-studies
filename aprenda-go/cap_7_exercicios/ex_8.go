package main

import (
	"fmt"
)

func main() {
	idade := 23

	switch {
	case idade >= 18:
		fmt.Println("maior de idade")
	case idade < 18:
		fmt.Println("menor de idade")
	}

}
