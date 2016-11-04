package repl

import (
	"bufio"
	"fmt"
	"github.com/unkiwii/goeros/repl/commands"
	"os"
	"strings"
)

const (
	infoCommand = ".info"
	prompt      = "> "
)

func Loop() (err error) {
	if commands.Has(infoCommand) {
		if err = commands.Execute(infoCommand); err != nil {
			panic(fmt.Sprintf("%s command not found"))
		}
	}
	fmt.Println("Type .help for more information\n")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(prompt)
	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())
		if len(t) > 0 {
			if commands.Has(t) {
				err = commands.Execute(t)
			} else {
				err = eval(t)
			}
			if err != nil {
				fmt.Println("error:", err)
			}
		}
		fmt.Printf(prompt)
	}

	return nil
}

func eval(t string) error {
	fmt.Printf("eval: %s\n", t)
	return nil
}
