package main

import "fmt"

func main() {
	var annee int = 2023
	if annee % 4 == 0 || annee % 4 == 0 && annee % 100 != 0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
