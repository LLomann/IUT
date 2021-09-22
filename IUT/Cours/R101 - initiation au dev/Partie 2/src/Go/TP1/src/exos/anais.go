package main

import "fmt"

func ArgentBanque(argent float64) []float64 {

	var banque float64 = 100
	var age float64

	for banque < argent {
		age += 1
		banque += age*10
		banque *= 1.035
	}
	var sortie[]float64 = []float64 {banque, age}

	return sortie
}

func main() {

	fmt.Println(ArgentBanque(99999))
}