package commands

import (
	"os"
	"os/exec"
)

func init() {
	Add(".clear", "clears the screen", Clear)
}

func Clear() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return nil
}
