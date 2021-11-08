package main

import "fmt"

type Etudiant struct {
	nombre   string
	apeilldo string
	notes    []float64
}

type Promotion struct {
	Promo          string
	liste_etudiant []Etudiant
}

func (e Etudiant) Moyenne() float64 {
	var somme float64
	var compte float64
	for i := 0; i < len(e.notes); i++ {
		somme = somme + e.notes[i]
		compte++
	}
	return somme / compte

}

func (e Etudiant) AjoutNote() []float64 {
	e.notes = append(e.notes, 18)
	return e.notes
}

func (p Promotion) Classer() []string {
	for i := 0; i<len(p1.liste_etudiant); i++ }

	table = []int{}
	var v int
	var tmp int

	for i := 0; i < len(table); i++ {

		v = tab[i]
		var j int = 0

		for j < len(table) && table[i] < v {

			j++
		}

		for j < len(table) {

			tmp = table[j]
			table[j] = v
			v = tmp
			j++
		}

		table = append(table, v)
	}
	return table

}

func main() {

	var e1 Etudiant = Etudiant{nombre: "Jean", apeilldo: "BOB", notes: []float64{20, 18, 19, 15, 16}}
	var e2 Etudiant = Etudiant{nombre: "Jeff", apeilldo: "Licorne", notes: []float64{20, 19, 0}}
	var e3 Etudiant = Etudiant{nombre: "Sam", apeilldo: "Pauly", notes: []float64{16, 18, 19, 19, 20}}
	var e4 Etudiant = Etudiant{nombre: "Léo", apeilldo: "Tranchet", notes: []float64{20, 18, 19, 15, 16}}
	var e5 Etudiant = Etudiant{nombre: "Ewen", apeilldo: "Bosquet", notes: []float64{0, 0, 0, 0, 0, 0}}
	var e6 Etudiant = Etudiant{nombre: "Lénny", apeilldo: "Barassin", notes: []float64{17, 18, 19, 15, 16}}
	var p1 Promotion = Promotion{Promo: "Pro", liste_etudiant: []Etudiant{e1, e4, e6}}
	var p2 Promotion = Promotion{Promo: "BG_qui_son_BO", liste_etudiant: []Etudiant{e2, e3, e5}}

	fmt.Println(e1, e2, e3, e4, e5, e6)
	fmt.Println(e3.Moyenne())
	fmt.Println(e3.AjoutNote())
	fmt.Println(p1, p2)

}
