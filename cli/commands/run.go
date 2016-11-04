package commands

import "fmt"

func init() {
	add("run", "compile and run Eros program", func(args []string) error {
		fmt.Printf("run %v\n", args)
		return nil
	})
}
