package repl

import (
	"bufio"
	"fmt"
	"github.com/unkiwii/goeros/info"
	"github.com/unkiwii/goeros/repl/commands"
	"os"
)

const (
	REPL_HELP   = "Type '.help' for more information"
	REPL_PROMPT = "> "
)

func Loop() error {
	fmt.Printf("\n%s repl, build: %s (%s)\n%s\n\n", info.NAME, info.VERSION, info.DATE, REPL_HELP)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(REPL_PROMPT)
	for scanner.Scan() {
		t := scanner.Text()
		if t != "" {
			if err := commands.Execute(scanner.Text()); err != nil {
				fmt.Println("error:", err)
			}
		}
		fmt.Printf(REPL_PROMPT)
	}

	return nil
}
