package main

import (
	"fmt"
)

var x int = 6
var y int = 5

var n *int = &x
var m *int = &y

func multiple(x *int, y *int) {
	*x = *x + y
}

func main() {

	x = 9

	fmt.Println(*n, *m)

	multiple(&n, &m)

}
