package main

import "fmt"

func tab (tab[] int ) int {
	var n = 0
	for i:= 0; i < len(tab); i++ {
		n += 1
	}
	return n
}

func main(){
	var T1 []int = []int{1,2,3,4,5,6,7,8,9}
	var T2 []int = []int{1,2,3,4,5}
	var T3 []int = []int{1,2}

	fmt.Println(tab(T1))
	fmt.Println(tab(T2))
	fmt.Println(tab(T3))
}