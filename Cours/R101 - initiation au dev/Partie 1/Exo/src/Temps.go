package main

import "fmt"

func main() {
	var sec int = 3800
	var min int = 0
	var heu int = 0

	min = sec / 60
	sec = sec % 60
	heu = min / 60
	min = ( min % 60 )

	fmt.Println("heure:",heu, "min:",min, "secode:",sec )
}
