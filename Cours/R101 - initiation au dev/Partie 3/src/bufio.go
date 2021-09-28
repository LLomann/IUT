package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	// Ouverture du fichier
	var filePath string = "test"
	var myFile *os.File
	var err error
	myFile, err = os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Préparation de la lecture
	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(myFile)

	// Lecture des lignes du fichier
	var compte int
	for scanner.Scan() {
		if compte == 5{
			break
		}
		log.Print("Je viens de lire : ", scanner.Text())
		compte += 1
	}

	log.Print(" nb de ligne:", compte)

	// Vérification que tout s'est bien passé
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	// Fermeture du fichier
	err = myFile.Close()
	if err != nil {
		log.Fatal(err)
	}
}
