package main

import "fmt"
import "os"
import "strconv"
import "github.com/daviejaneway/C4G/src"

func foo(i int) int {
  j := i * i
  
  defer C4G.NewContract("Contract over 'i'", i > 0, j < 100)
  
  return j
}


func main() {
  s := os.Args[1]
  
  i, err := strconv.Atoi(s)
  
  if err != nil {
    panic("usage: go basic_contract.go <int>")
  }
  
  fmt.Println(foo(i))
}
