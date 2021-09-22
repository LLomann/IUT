package main
import "fmt"

func SommeDiviseurs(m int) int {
	var somme int
	for i:=1; i < m; i++ {
		if m % i == 0 {
			somme += i
		}
	}
	return somme
}
func amicaux(m int, n int) bool {
	return SommeDiviseurs(m) == n && SommeDiviseurs(n) == m
}

func main () {
	fmt.Println(amicaux(1184, 1210))
}