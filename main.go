package main

import (
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, this is a REPL of monkey programming language\n", user.Username)
	fmt.Printf("Fell free to write commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
