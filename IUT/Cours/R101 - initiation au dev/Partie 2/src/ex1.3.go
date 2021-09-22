package main

import "fmt"

func f(x *int, y *int, b bool) {
	if b {
		*x += 1 // ajoute 1 à la val dans 0x001 -> n
		*y -= 2 // enleve 1 à la val dans 0x001 -> n
	} else {
		*x, *y = *y, *x
	}
}

func main() {
	var n int = 0 // 0x001 numéro adresse
	var m int = 10 // 0x002 numéro adresse
	fmt.Println(&n)
	fmt.Println("Au début n vaut", n, "et m vaut", m)
	for i := 0; i < 10; i++ {
		f(&n, &m, i%2 == 0) // 0x001 0x002 bool
		fmt.Println("Après l'itération", i, "n vaut", n, "et m vaut", m)
	}
	fmt.Println("À la fin n vaut", n, "et m vaut", m)
}
