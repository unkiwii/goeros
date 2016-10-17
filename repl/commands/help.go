package commands

import (
	"fmt"
	"sort"
)

func init() {
	Add(".help", "shows this message", Help)
}

func Help() error {
	fmt.Print("\nAvailable commands:\n")
	maxNameLength := 0
	allSorted := All()
	sort.Sort(allSorted)
	for _, entry := range allSorted {
		l := len(entry.name)
		if l > maxNameLength {
			maxNameLength = l
		}
	}
	for _, entry := range allSorted {
		fmt.Printf("\n   %*s   %s", maxNameLength, entry.name, entry.desc)
	}
	fmt.Println("\n")

	return nil
}
