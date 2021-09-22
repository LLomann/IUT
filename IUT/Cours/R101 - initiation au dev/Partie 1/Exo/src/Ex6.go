package main

import "fmt"


func f(tab[]int) int {
	return tab[2]
}

func main() {
	var T1 []int = []int{1,2,3,4,5,6,7,8}
	var T2 []int = []int{12,54,96,81}
	var T3 []int = []int{0,2,5,4,8}
	fmt.Println(f(T1))
	fmt.Println(f(T2))
	fmt.Println(f(T3))
}


func g(tab[]int, n int) int {
	return tab[n]
}

func main() {
	var T4 []int = []int{1,2,3,4,5,6,7,8}
	var T5 []int = []int{12,54,96,81}
	var T6 []int = []int{0,2,5,4,8}
	fmt.Println(g(T4, 5))
	fmt.Println(g(T5, 2))
	fmt.Println(g(T6, 4))
}
