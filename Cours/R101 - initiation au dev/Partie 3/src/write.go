package main

import (
	"fmt"
	"log"
	"os"
)


func main() {
	var myFile *os.File
	var err error
	myFile, err = os.Create("test")
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintln(myFile, "Bonjour")
	if err != nil {
		log.Fatal(err)
	}

	var n int = 5
	for i:=0; i<10; i++{
		fmt.Fprintln(myFile, i,"x", n,"=", i*n)
	}
 
	_, err = fmt.Fprintln(myFile, "Au revoir")
	if err != nil {
		log.Fatal(err)
	}

	err = myFile.Close()
	if err != nil {
		log.Fatal(err)
	}
}