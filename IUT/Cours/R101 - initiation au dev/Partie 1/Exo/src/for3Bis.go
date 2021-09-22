package main

import "fmt"

func main() {
	var var1 int = 5
	var var2 int = 10
	var r = 0
	var co int = 0
	for co < var2 {
		r = var1 + r
		co++
	}
	fmt.Println(r)
}
