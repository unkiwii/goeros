package commands

import (
	"os"
)

func init() {
	add(".exit", "exit to the os", exit)
}

func exit() error {
	os.Exit(0)
	return nil
}
