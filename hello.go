package main

import (
  "fmt"
)

func main(){
  
  r := []rune("Hello, universe.\n") // make string into an array?
  x := int len(r)
  fmt.Printf(string(x))
  for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
    fmt.Printf(string(r[j]))
    r[i], r[j] = r[j], r[i]
  }

}