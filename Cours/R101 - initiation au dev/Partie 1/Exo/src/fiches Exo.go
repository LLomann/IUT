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

func multiple_de(n,x,y int){	// ex 4.1
	for i := x; i <= y; i = i + n {
		fmt.Println("les multiples de", n,"entre", x,"et", y,"son:", i)
	}
}

func tab(tab[]int){				// exercice 4.1 suiteBis
	for i:=0; i<len(tab); i++{
		fmt.Println(tab[i])
	}
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

	multiple_de(9,0,100) // exercice 4.1
	
	for i:=0; i<10; i++{ // exercice 4.1 suite
		fmt.Println("bonjour")
	}

	var tableau1 []int = []int{0,1,2,3,4,5,6,7,8,9}		// exercice 4.1 suiteBis
	var tableau2 []int = []int{0,1,2,3,4,5,6,7,8,9,10,11}
	var tableau3 []int
	tab(tableau1)
	tab(tableau2)
	tab(tableau3)

	var compte int  // exercice 4.2
	var tour int = 2
	for i:=0; i < 10; i = i + tour {
		compte = compte + tour
	}
	fmt.Println("le nb de tour est de:", compte)


	var m float64 = 0	// exercice 4.3
  	var tableau []float64 = []float64 {5.4, 9.4, 7.5, 16.5}
  	for i := 0; i < len(tableau); i = i + 1 {
    	m = m + tableau[i]
    }
  	m = m / float64(len(tableau))
  	fmt.Println(m)

}