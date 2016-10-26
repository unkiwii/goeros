package repl

import (
	"bufio"
	"fmt"
	"github.com/unkiwii/goeros/info"
	"github.com/unkiwii/goeros/repl/commands"
	"os"
	"strings"
)

const (
	REPL_HELP   = "Type '.help' for more information"
	REPL_PROMPT = "> "
)

func Loop() (err error) {
	fmt.Printf("\n%s repl, build: %s (%s)\n%s\n\n", info.NAME, info.VERSION, info.DATE, REPL_HELP)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(REPL_PROMPT)
	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())
		if len(t) > 0 {
			if commands.IsValid(t) {
				err = commands.Execute(t)
			} else {
				err = eval(t)
			}
			if err != nil {
				fmt.Println("error:", err)
			}
		}
		fmt.Printf(REPL_PROMPT)
	}

	return nil
}

func eval(t string) error {
	fmt.Printf("eval: %s\n", t)
	return nil
}
