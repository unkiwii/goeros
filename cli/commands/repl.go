package commands

import "github.com/unkiwii/goeros/repl"

func init() {
	add("repl", "enters interactive mode", func(args []string) error {
		return repl.Loop()
	})
}
