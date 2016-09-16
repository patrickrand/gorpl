# gorpl
`gorpl` (rhymes with "purple") is a REPL based off a subset of the Go syntax.

## Installation

```bash
$ go get github.com/patrickrand/gorpl
```

## Usage

```go
$ gorpl 
gorpl:0000> import "fmt"
=> true
gorpl:0001> fmt.Println("Hello, world!")
=> "Hello, world!"
gorpl:0002> x := 123
=> 123
gorpl:0003> s := fmt.Sprintf("ABC is easy as %d", x)
=> "ABC is easy as 123"
gorpl:0004> exit

```