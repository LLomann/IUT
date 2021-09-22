package main

import "fmt"

func f(tab[] int, n int){
	if len(tab) > (n + 1){
		fmt.Println(true, tab[n])
	}else{
		fmt.Println(false,0)
	}
}

func main() {
	var t []int = []int{1,2,3,45,6,7,8}
	var c = 10
	f(t, c)
}
