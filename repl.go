package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type REPL struct {
	scan   *bufio.Scanner
	ticker int
	output io.Writer
}

func NewREPL() *REPL {
	return &REPL{
		ticker: 0,
		scan:   bufio.NewScanner(os.Stdin),
		output: os.Stdout,
	}
}

func (r *REPL) Read() string {
	return r.scan.Text()
}

func (r *REPL) Eval(s string) string {
	if s == "exit" {
		exit(r)
	}
	return s
}

func (r *REPL) Print(s string) {
	fmt.Fprintln(r.output, `=>`, s)
}

func (r *REPL) Loop() bool {
	prompt(r)
	return r.scan.Scan()
}

func prompt(r *REPL) {
	t := strconv.Itoa(r.ticker)
	for n := len(t); n < 4; n++ {
		t = "0" + t
	}
	fmt.Fprintf(r.output, `gorpl:%s> `, t)
	if r.ticker++; r.ticker == 10000 {
		r.ticker = 0
	}
}

func exit(r *REPL) {
	fmt.Fprintln(r.output, "exiting gorpl")
	os.Exit(0)
}

var TmplMain = `
import (
	{{}}
)

func main() {

}
`
