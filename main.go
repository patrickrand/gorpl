package main

func main() {
	repl := NewREPL()
	repl.Prompt()
	for repl.Scan() {
		input := repl.Text()
		if input == "exit" {
			repl.Exit()
		}
		repl.Reply(input)
		repl.Prompt()
	}
}
