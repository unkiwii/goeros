package commands

import (
	"fmt"
	"sort"
)

func init() {
	add(".help", "shows this message", help)
}

func help() error {
	fmt.Print("\nAvailable commands:\n")
	maxNameLength := 0
	commands := getAllCommands()
	sort.Sort(commands)
	for _, entry := range commands {
		l := len(entry.name)
		if l > maxNameLength {
			maxNameLength = l
		}
	}
	for _, entry := range commands {
		fmt.Printf("\n   %*s   %s", maxNameLength, entry.name, entry.desc)
	}
	fmt.Println("\n")

	return nil
}
