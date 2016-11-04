package commands

import "fmt"

func init() {
	add("build", "compile files and dependencies", func(args []string) error {
		fmt.Printf("build %v\n", args)
		return nil
	})
}
