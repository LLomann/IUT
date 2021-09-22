package main

import "fmt"

func g(tab[] int, n int) int {
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

func f(x int) int {
	var S int = 0
	for i := 1; i < x; i++{
		if x%i == 0 {
			S = S + i
		}
	}
	return S
}

func amicaux(x int, y int) bool {
	return f(x)==y && x==f(y)
}


func main() {
	var n int = 220
	var m int = 284
	fmt.Println(amicaux(n,m))
}
