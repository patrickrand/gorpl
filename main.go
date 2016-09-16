package main

func main() {
	repl := NewREPL()
	repl.Prompt()
	for repl.scan.Scan() {
		input := repl.scan.Text()
		if input == "exit" {
			repl.Exit()
		}
		repl.Reply(input)
		repl.Prompt()
	}
}
