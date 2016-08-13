package commands

import (
	"os"
)

func init() {
	Add(".exit", "exit to the os", Exit)
}

func Exit() error {
	os.Exit(0)
	return nil
}
