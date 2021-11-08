package main

import "fmt"

type cuple struct {
	first  int
	second int
}
type persones struct {
	nombre   string
	apeilldo string
	anos     int
}

func main() {
	var c cuple = cuple{first: 1, second: 2}
	fmt.Println(c, c.first, c.second)

	c.first = 3
	fmt.Println(c, c.first, c.second)

	var c1 persones = persones{nombre: "Zame", apeilldo: "poly", anos: 71}
	var c2 persones = persones{nombre: "Jeff", anos: 2480}
	var c3 persones = persones{nombre: "ET"}
	fmt.Println(c1, c2, c3)

	c1.nombre = "Sam"
	fmt.Println(c1)

}
