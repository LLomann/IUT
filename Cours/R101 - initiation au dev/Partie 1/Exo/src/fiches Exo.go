package main

import "fmt"

func Temps(){ 				// ex 2.5 suite
	var sec int = 3800
	var min int = 0
	var heu int = 0

	min = sec / 60
	sec = sec % 60
	heu = min / 60
	min = min % 60

	fmt.Println("le temps est de", heu, "heure(s),", min, "minute(s) et", sec, "seconde(s)")
}



func main(){
	fmt.Println("Bonjour") // ex 1

	var a string = "Hello " // ex 2.4 et 2.5
	fmt.Println(a)

	var s string = "world!" // 2.5 
	s = a + s
	fmt.Println(s)

	var cal int = 5 // ex 2.5 suite
	cal = ((6 * cal) + 4 )/ 2
	fmt.Println(cal)

	var n int = 3
	fmt.Println(5 < n)

	var sec int = 3800 // ex 2.5 suite
	var min int = 0
	var heu int = 0

	min = sec / 60
	sec = sec % 60
	heu = min / 60
	min = min % 60

	fmt.Println("le temps est de", heu, "heure(s),", min, "minute(s) et", sec, "seconde(s)")

}

	



