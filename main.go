package main

import (
	"fmt"
	"os"
)

func main() {
	repl := NewREPL()
	repl.Prompt()
	for repl.scan.Scan() {
		input := repl.scan.Text()
		if input == "exit" {
			fmt.Println("\nexiting gorpl")
			os.Exit(0)
		}
		fmt.Println("=>", input)
		repl.Prompt()
	}
}
