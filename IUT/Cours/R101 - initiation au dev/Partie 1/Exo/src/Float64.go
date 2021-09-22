package main

import "fmt"

func main() {
  var m float64 = 0
  var tableau []float64 = []float64 {5.4, 9.4, 7.5, 16.5}
  for i := 0; i < len(tableau); i = i + 1 {
    m = m + tableau[i]
    }
  m = m / float64(len(tableau))
  fmt.Println(m)
}
