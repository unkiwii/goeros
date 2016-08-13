package commands

import (
	"fmt"
)

func init() {
	Add(".help", "shows this message", Help)
}

func Help() error {
	fmt.Print("\nAvailable commands:\n")
	for name, desc := range All() {
		fmt.Print("\n\t")
		fmt.Print(name)
		fmt.Print("\t")
		fmt.Print(desc)
	}
	fmt.Println("\n")

	return nil
}
