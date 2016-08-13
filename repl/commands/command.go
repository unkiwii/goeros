package commands

import (
	"fmt"
)

type command func() error

type entry struct {
	name string
	desc string
	cmd  command
}

var cmds map[string]entry = make(map[string]entry)

func exists(name string) bool {
	_, exists := cmds[name]
	return exists
}

func All() map[string]string {
	all := make(map[string]string)
	for name, entry := range cmds {
		all[name] = entry.desc
	}
	return all
}

func Add(name string, desc string, cmd command) error {
	if exists(name) {
		return fmt.Errorf("command %s already added", name)
	}
	cmds[name] = entry{name, desc, cmd}
	return nil
}

func Remove(name string) error {
	if !exists(name) {
		return fmt.Errorf("command %s already removed or was never added", name)
	}
	delete(cmds, name)
	return nil
}

func Execute(name string) error {
	c, ok := cmds[name]
	if ok {
		return c.cmd()
	}

	return fmt.Errorf("invalid command '%s'", name)
}
