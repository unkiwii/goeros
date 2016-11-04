package commands

import (
	"fmt"
	"strings"
)

type command func(args []string) error

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

var cmds map[string]entry = make(map[string]entry)

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

func Has(name string) bool {
	_, has := cmds[name]
	return has
}

func Execute(name string, args []string) error {
	c, found := cmds[name]
	if !found {
		c, found = cmds["help"]
		if !found {
			panic("can't find help command")
		}
	}

	return c.cmd(args)
}
