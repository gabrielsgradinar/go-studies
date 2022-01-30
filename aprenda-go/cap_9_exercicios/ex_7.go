package main

import "fmt"

func main() {
	multi_slice := [][]string{
		{"gabriel", "gradinar", "games"},
		{"julia", "sofreio", "estudos"},
		{"igor", "gradinar", "nft"},
	}

	for _, value := range multi_slice {
		for _, info := range value {
			fmt.Println(info)
		}
	}

}
