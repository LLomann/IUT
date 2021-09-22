package main

import "fmt"

func main() {
	var n int = 3
	var m int = 3
	var tab [][]int = make([][]int, n)
	for i:=0; i<n; i++{
		var t []int = make([]int, m)
		tab[i] = t
		fmt.Println(tab)
		for j:=0; j<m; j++{
			tab[i][j] = i*j
		}
	}	
}
