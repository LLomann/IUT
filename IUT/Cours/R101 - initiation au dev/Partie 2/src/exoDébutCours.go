package main

import "fmt"

func main(){
	var tab []int = []int{1,2,3,4}
	var Adresse_tab *[]int = &tab

	for i:= 0; i< len(tab); i++ {
		fmt.Println(&tab[i], tab[i])
	}
	fmt.Println(&Adresse_tab, tab)
}