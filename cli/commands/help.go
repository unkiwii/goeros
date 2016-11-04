package commands

import (
	"bytes"
	"fmt"
	"github.com/unkiwii/goeros/info"
	"sort"
)

func init() {
	add("help", "shows this message or more information on a specific command", func(args []string) error {
		var buffer bytes.Buffer

		buffer.WriteString(info.Name)
		buffer.WriteString(" is a tool for managing eros source code\n\n")
		buffer.WriteString("Usage:\n\n\t")
		buffer.WriteString(info.Name)
		buffer.WriteString(" command [arguments]\n\n")
		buffer.WriteString("The commands are:\n")

		commands := getAllCommands()
		sort.Sort(commands)
		for _, entry := range commands {
			buffer.WriteString("\n\t")
			buffer.WriteString(entry.name)
			buffer.WriteString("\t")
			buffer.WriteString(entry.desc)
		}

		buffer.WriteString("\n\nUse \"")
		buffer.WriteString(info.Name)
		buffer.WriteString(" help [command]\" for more information about a command")

		fmt.Println(buffer.String())
		return nil
	})
}
