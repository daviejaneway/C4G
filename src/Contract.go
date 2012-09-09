package C4G

import "fmt"
import "errors"

type Contract struct {
  ident string
  pre bool
  post bool
}

func (c Contract) Pre () error {
  if !c.pre {
    return errors.New(fmt.Sprintf("Pre condition not satisfied for: %s", c.ident))
  }
  
  return nil
}

func (c Contract) Post () error {
  if !c.post {
    return errors.New(fmt.Sprintf("Post condition not satisfied for: %s", c.ident))
  }
  
  return nil
}

func NewContract(ident string, pre bool, post bool) bool {
  c := Contract{ident:ident, pre:pre, post:post}
  
  if err := c.Pre(); err != nil {
    panic(err)
    return false
  }
  
  if err := c.Post(); err != nil {
    panic(err)
    return false
  }
  
  return true
}