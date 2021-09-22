package main

import "fmt"
import "math/rand"
import "math"
import "time"

func tablorandom(longueur int, minimum int, maximum int) []float64 {
	var tablo[]float64
	rand.Seed(time.Now().UnixNano())
	for i:=0; i < longueur; i++ {
		tablo = append(tablo, float64(rand.Intn(maximum - minimum) + minimum) + float64(rand.Intn(4))/4)
	}
	return tablo
}


func main() {
	
	var notes[]float64 = tablorandom(20, 12, 20)
	var moyenne float64
	for i:=0; i < len(notes); i++ {
		moyenne += notes[i]
	}
	fmt.Println(notes)
	fmt.Println(math.Floor(moyenne/float64(len(notes))*100)/100)  
}
