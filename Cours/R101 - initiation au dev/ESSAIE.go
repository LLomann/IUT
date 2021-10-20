package main

import "fmt"

func table_de(x *int, y *int){

	for i:=0; i < y; i++{
		fmt.Println(i, "x", x, "=", i*x)
	}
}

func main(){

	table_de(5,10)
	
}