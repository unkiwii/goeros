package commands

import (
	"fmt"
)

func init() {
	Add("run", "compile and run Eros program", Run)
}

func Run(args []string) error {
	fmt.Printf("run %v\n", args)
	return nil
}
