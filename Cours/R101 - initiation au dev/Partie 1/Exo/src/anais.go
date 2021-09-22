package main

import "fmt"

func main() {
  var epargne int = 100
  var age int = 0
  var depo int = 0
  for age < 40; age = age + 1 {
    depo = 10 * age
    epargne = epargne + depo
    epargne = epargne * 1.035
    }
  fmt.Println( "A l'age de", age, "Anaïs aura", epargne, " sur son compte ( la chance ).")
}
