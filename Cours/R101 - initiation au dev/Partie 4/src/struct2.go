package main

import "fmt"

type persones struct {
	nombre   string
	apeilldo string
	anos     int
}

func main() {
	var c1 persones = persones{nombre: "Zame", apeilldo: "poly", anos: 71}
	var c2 persones = persones{nombre: "Jeff", anos: 2480}
	var c3 persones = persones{nombre: "ET"}
	fmt.Println(c1, c2, c3)
}
