package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"log"
	"strconv"
)

func main(){

	// Ouverture du fichier
	var filePath string = "notes.csv"
	var myFile *os.File
	var err error
	var compte int = 0
	var ligne int = 2
	var note int
	var bonus int
	var total int
	var nom string = "Jean-François"
	var somme int =0
	var nb int = 0
	
	myFile, err = os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Préparation de la lecture
	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(myFile)


	scanner.Scan() // lis 1er ligne
	
	
	for scanner.Scan(){ // lis a partir de le seconde ligne jusqu'à la fin
		var tab []string = strings.Split(scanner.Text(),",")
		note,err = strconv.Atoi(tab[1])
		bonus,err= strconv.Atoi(tab[2])
		total,err = strconv.Atoi(tab[3])

		if err != nil{
			fmt.Println("la note n'est pas écrit en chiffre à la ligne", ligne)
		}

		if note+bonus != total{ 
			fmt.Println("il y a une faute a la ligne:", ligne)
			compte += 1
		}

		if tab[0] == nom{
			somme = total + somme
			nb += 1
		}
		ligne += 1
	}
	fmt.Println("la moyenne de",nom,"est de", float64(somme)/float64(nb))
	fmt.Println("Il y a en tous", compte,"fautes")
}