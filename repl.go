package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type REPL struct {
	ticker int
	scan   *bufio.Scanner
	output io.Writer
}

func NewREPL() *REPL {
	return &REPL{
		ticker: 0,
		scan:   bufio.NewScanner(os.Stdin),
		output: os.Stdout,
	}
}

func (r *REPL) SetOutput(w io.Writer) {
	r.output = w
}

func (r *REPL) Prompt() {
	t := strconv.Itoa(r.ticker)
	for n := len(t); n < 4; n++ {
		t = "0" + t
	}
	fmt.Fprintf(r.output, `gorpl:%s> `, t)
	r.tick()
}

func (r *REPL) tick() {
	if r.ticker++; r.ticker == 10000 {
		r.ticker = 0
	}
}
