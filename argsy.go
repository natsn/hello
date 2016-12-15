package main
// This program prints its command line arguments

import (
  "fmt"
  "os"
)

func main() {
  // Print command line arguments
  for i := 1; i < len(os.Args); i++ {
    fmt.Println(os.Args[i])
  }
}