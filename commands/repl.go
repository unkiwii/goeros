package commands

import (
	"bufio"
	"fmt"
	"github.com/unkiwii/goeros/info"
	"os"
)

const (
	REPL_HELP   = "Type '.help' for more information"
	REPL_PROMPT = "> "
)

func init() {
	Add("repl", "enters interactive mode", Repl)
}

func Repl(args []string) error {
	fmt.Printf("\n%s repl %s (%s)\n%s\n\n", info.NAME, info.VERSION, info.DATE, REPL_HELP)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(REPL_PROMPT)
	for scanner.Scan() {
		fmt.Println("read:", scanner.Text())
		fmt.Printf(REPL_PROMPT)
	}

	return nil
}
