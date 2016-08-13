package commands

import (
	"github.com/unkiwii/goeros/repl"
)

func init() {
	Add("repl", "enters interactive mode", Repl)
}

func Repl(args []string) error {
	return repl.Loop()
}
