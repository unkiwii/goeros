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

var cmds map[string]entry = make(map[string]entry)

func exists(name string) bool {
	_, exists := cmds[name]
	return exists
}

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
	if exists(name) {
		return fmt.Errorf("command %s already added", name)
	}
	cmds[name] = entry{name, desc, cmd}
	return nil
}

func remove(name string) error {
	if !exists(name) {
		return fmt.Errorf("command %s already removed or was never added", name)
	}
	delete(cmds, name)
	return nil
}

func IsValid(name string) bool {
	return len(name) > 1 && name[:1] == "."
}

func Execute(name string) error {
	c, ok := cmds[name]
	if ok {
		return c.cmd()
	}

	return fmt.Errorf("invalid command '%s'", name)
}
