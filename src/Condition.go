//C4G implements a very basic Design by Contract
//framework. In C4G, a Contract is an implicit
//control structure defined as one of more Conditions.
//The Contract is satisfied is none of its Conitions panic.
package C4G

import "fmt"

//A Conditional takes a bool and panics if
//that bool is false, hence the contract is broken
type Conditional interface {
	execute(b bool)
}

//A Condition is decoupled from its
//Execute method, which means we may
//declare and reuse our Condtion objects
type Condition struct {
	Ident string
}

//This method defines the Contract's
//expectations. The bool param must be
//true for control to pass back to
//enclosing (function) scope
func (c Condition) Execute(b bool) {
	if !b {
		panic(fmt.Sprintf("Condition not met for: '%s'", c.Ident))
	}
}
