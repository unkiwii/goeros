package repl

import (
	"bufio"
	"fmt"
	"github.com/unkiwii/goeros/lexer"
	"github.com/unkiwii/goeros/repl/commands"
	"os"
	"strings"
)

const (
	INFO_COMMAND = ".info"
	PROMPT       = "> "
)

// Loop start and executes the main loop of the repl
func Loop() (err error) {
	if commands.Has(INFO_COMMAND) {
		if err = commands.Execute(INFO_COMMAND); err != nil {
			panic(fmt.Sprintf("%s command not found", INFO_COMMAND))
		}
	}

	fmt.Println("Type .help for more information")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(PROMPT)
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
		fmt.Printf(PROMPT)
	}

	return nil
}

func eval(t string) error {
	l := lexer.New("eval")
	c := l.Lex(t)
	for {
		select {
		case r, ok := <-c:
			if ok {
				fmt.Printf("%s\n", r)
			} else {
				return nil
			}
		}
	}
}
