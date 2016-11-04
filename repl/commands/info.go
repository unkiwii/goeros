package commands

import (
	"fmt"
	"github.com/unkiwii/goeros/info"
)

func init() {
	add(".info", "shows build information", func() error {
		fmt.Printf("\n%s repl, build: %s (%s)\n\n", info.Name, info.Version, info.DateCompiled())
		return nil
	})
}
