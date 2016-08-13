package commands

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/unkiwii/goeros/info"
)

type command func(args []string) error

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

func Execute(name string, args []string) error {
	c, ok := cmds[name]
	if ok {
		return c.cmd(args)
	} else {
		return usage()
	}
}

func usage() error {
	var buffer bytes.Buffer

	buffer.WriteString(info.NAME)
	buffer.WriteString(" is a tool for managing eros source code\n\n")
	buffer.WriteString("Usage:\n\n\t")
	buffer.WriteString(info.NAME)
	buffer.WriteString(" command [arguments]\n\n")
	buffer.WriteString("The commands are:\n")

	for _, cmd := range cmds {
		buffer.WriteString("\n\t")
		buffer.WriteString(cmd.name)
		buffer.WriteString("\t")
		buffer.WriteString(cmd.desc)
	}

	buffer.WriteString("\n\nUse \"")
	buffer.WriteString(info.NAME)
	buffer.WriteString(" help [command]\" for more information about a command")

	return errors.New(buffer.String())
}
