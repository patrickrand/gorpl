package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type REPL struct {
	*bufio.Scanner
	ticker int
	output io.Writer
}

func NewREPL() *REPL {
	return &REPL{
		ticker:  0,
		Scanner: bufio.NewScanner(os.Stdin),
		output:  os.Stdout,
	}
}

func (r *REPL) Prompt() {
	t := strconv.Itoa(r.ticker)
	for n := len(t); n < 4; n++ {
		t = "0" + t
	}
	fmt.Fprintf(r.output, `gorpl:%s> `, t)
	r.tick()
}

func (r *REPL) Reply(output string) {
	fmt.Fprintln(r.output, `=>`, output)
}

func (r *REPL) Exit() {
	fmt.Fprintln(r.output, "exiting gorpl")
	os.Exit(0)
}

func (r *REPL) tick() {
	if r.ticker++; r.ticker == 10000 {
		r.ticker = 0
	}
}

var TmplMain = `
import (
	{{}}
)

func main() {

}
`
