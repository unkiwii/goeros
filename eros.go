package main

import (
	"fmt"
	"github.com/unkiwii/goeros/commands"
	"os"
)

func main() {
	command := "repl"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	args := []string{}
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	if err := commands.Execute(command, args); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
