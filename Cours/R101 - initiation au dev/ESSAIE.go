package main

import (
	"fmt"
	"os"
	"log"
)
func main() {
	var myFile *os.File
	var err error
	myFile, err = os.Create("Ouch")
	if err != nil{
		log.Fatal(err)
	}

	_, err = fmt.Fprintln(myFile, "Ouch parce que t nul")
	if err != nil{
		log.Fatal(err)
	}

	err = myFile.Close()
	if err != nil{
		log.Fatal(err)
	}
}