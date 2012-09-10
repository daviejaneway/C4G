C4G - Minimal Contracts for Go
===

C4G is a basic Design By Contract framework for the Go programming language.

## Installation

    go get github.com/daviejaneway/C4G/src
    
## Example

    var pre = Condition{"This is a pre-condition, i must be > 0"}
    var post = Condition{"This is a post-condition, i must be < 10"}
    
    func foo(i int) {
      pre.Execute(i > 0)
      defer post.Execute(i < 10)
      
      //Your code here
    }