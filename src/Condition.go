package C4G

import "fmt"

type Conditional interface {
  execute(b bool)
}

type Condition struct {
  ident string
}

func (c Condition) Execute(b bool) {
  if ! b {
    panic(fmt.Sprintf("Condition not met for: '%s'", c.ident))
  }
}