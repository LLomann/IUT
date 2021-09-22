package main

import "fmt"

func main() {
	var i int
	var tableau []int = []int {1,2,3,1,2,3,4,5,6,7}
	for i = 0; i < len(tableau); i = i + 1 {
		fmt.Println(tableau[i])
	}
	fmt.Println(i)
}
