package main

import "fmt"

func main() {
	var s int
	var tableau []int = []int {1,1,1,1,1,1,1,1,1,1,1,1}
	for i := 0; i < len(tableau); i = i + 1 {
		s = s + tableau[i]
		fmt.Println(tableau[i])
	}
	fmt.Println(s)
}
