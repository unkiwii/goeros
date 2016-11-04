package commands

import (
	"os"
)

func init() {
	add(".exit", "exit to the os", func() error {
		os.Exit(0)
		return nil
	})
}
