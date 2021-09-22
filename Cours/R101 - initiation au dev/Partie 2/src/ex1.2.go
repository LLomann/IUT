package main

import "fmt"

func main() {
	var n int = 10
	var a *int = &n
	*a = 12
	fmt.Println(n, *a)

	n = 6
	fmt.Println(n, *a)

	var m int = 7
	a = &m
	fmt.Println(n, *a, m)
}
