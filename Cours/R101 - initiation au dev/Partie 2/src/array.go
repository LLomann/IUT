package main

import "fmt"

func f(t [3]int) {
	for i := 0; i < len(t); i++ {
		t[i] = 0
	}
}

func main() {
	var tab [3]int = [3]int{1, 2, 3}
	fmt.Println(tab)
	f(tab)
	fmt.Println(tab)
}
