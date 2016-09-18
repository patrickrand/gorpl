package main

func main() {
	repl := NewREPL()
	for repl.Loop() {
		input := repl.Read()

		output := repl.Eval(input)
		repl.Print(output)
	}
}
