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
	var ligne int = 1

	myFile, err = os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Préparation de la lecture
	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(myFile)


	for scanner.Scan(){
		var tab []string = strings.Split(scanner.Text(),",")
		
		if tab[1]int + tab[2]int != tab[3]int{
			ligne += 1
			tab[] 
			fmt.Println("il y a une faute a la ligne:"ligne)
		}
	}
}