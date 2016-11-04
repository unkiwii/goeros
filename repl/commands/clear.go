package commands

import (
	"os"
	"os/exec"
)

func init() {
	add(".clear", "clears the screen", func() error {
		cmd := exec.Command(getClearCommand())
		cmd.Stdout = os.Stdout
		cmd.Run()
		return nil
	})
}

func getClearCommand() string {
	//TODO: check for OS and run a different clear command
	return "clear"
}
