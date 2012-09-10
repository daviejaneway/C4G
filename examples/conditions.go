package main

import "os"
import "fmt"
import "strconv"
import "github.com/daviejaneway/C4G/src"

//Declare an implicit Contract consisting
//of a pre-condition and a post-condition.
var (
  pre = C4G.Condition{"pre-conditon: i must be > 0"}
  post = C4G.Condition{"post-condition: i*i must be < 100"}
)

//Declare some function to test our Contract.
//Here, we simply return i squared. However,
//a runtime error will occur if and when we
//fail to satisfy our Contract.
func foo2(i int) int {
  pre.Execute(i > 0)
  
  j := i * i
  
  //defer helps us define a true post-condition which
  //will not be executed until AFTER we return j
  defer post.Execute(j < 100)
  
  return j
}

func main() {
  s := os.Args[1]
  
  i, err := strconv.Atoi(s)

  if err == nil {
    i = foo2(i)
  
    fmt.Println(i)
  } else {
    fmt.Println("usage: go run conditions.go <int>")
  }
}