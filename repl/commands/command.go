package commands

import (
	"fmt"
	"strings"
)

type command func() error

type entry struct {
	name string
	desc string
	cmd  command
}

// make entries sortable
type entries []entry

func (a entries) Len() int {
	return len(a)
}

func (a entries) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a entries) Less(i, j int) bool {
	return strings.Compare(a[i].name, a[j].name) < 0
}

var cmds = make(map[string]entry)

func getAllCommands() entries {
	all := make(entries, len(cmds))
	i := 0
	for _, entry := range cmds {
		all[i] = entry
		i++
	}
	return all
}

func add(name string, desc string, cmd command) error {
	if Has(name) {
		return fmt.Errorf("command %s already added", name)
	}
	cmds[name] = entry{name, desc, cmd}
	return nil
}

// Has take a string and returns true when the given string is a valid command
func Has(name string) bool {
	_, has := cmds[name]
	return has
}

// Execute take a string and executes the command with that name
func Execute(name string) error {
	if c, ok := cmds[name]; ok {
		return c.cmd()
	}

	return fmt.Errorf("invalid command '%s'", name)
}
