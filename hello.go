package main

import (
  "fmt"
  "math"
  "strconv"
)

func toFloat(num float64) string {
    return strconv.FormatFloat(num, 'f', 6, 64) // what does this do?
}

func main(){
  r := []rune("Hello, universe.") // make string into an array?
  x := len(r) // := is syntactic sugar for type inference
  fmt.Printf(string(x))
  for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
    r[i], r[j] = r[j], r[i]
  } // reverse a string in place.

  fmt.Println(string(r))
  fmt.Println("%d",float64(math.E))
  fmt.Println(toFloat(math.Pi))
}
