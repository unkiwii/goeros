package commands

import (
	"fmt"
)

func init() {
	Add("build", "compile files and dependencies", Build)
}

func Build(args []string) error {
	fmt.Printf("build %v\n", args)
	return nil
}
