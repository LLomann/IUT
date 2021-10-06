package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
)

const numGen = 1000
const probaError = 200

var firstLine = []string{"Prénom", "Note", "Bonus", "Total"}

var names = []string{"Jean", "Léa", "Louis", "Lucie", "Marc", "Claire", "Antoine", "Anne", "Rémi", "Maria", "Luc", "Sylvie", "René", "Anne-Claire", "Serge", "Louise", "Jacques", "Sonia", "Gérard", "Charlotte", "Vincent", "Éléonore", "Jean-François", "Christine", "Christian", "Véronique"}

func genLine() []string {
	name := names[rand.Intn(len(names))]
	note := rand.Intn(15) + 6
	bonus := rand.Intn(2)
	total := note + bonus
	if rand.Intn(probaError) == 0 {
		total--
		fmt.Println("Added one error")
	}
	return []string{name, fmt.Sprint(note), fmt.Sprint(bonus), fmt.Sprint(total)}
}

func main() {

	file, err := os.Create("notes.csv")
	if err != nil {
		log.Fatal(err)
	}

	w := csv.NewWriter(file)

	if err := w.Write(firstLine); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < numGen; i++ {
		if err := w.Write(genLine()); err != nil {
			log.Fatal(err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}

}
