package commands

import (
	"fmt"
	"github.com/unkiwii/goeros/info"
)

func init() {
	add(".license", "shows license", func() error {
		fmt.Printf(`
       DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
                   Version 2, December 2004

Copyright (C) %s %s <%s>

Everyone is permitted to copy and distribute verbatim or modified
copies of this license document, and changing it is allowed as long
as the name is changed.

           DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

0. You just DO WHAT THE FUCK YOU WANT TO.

`, info.YearCompiled, info.Author, info.AuthorMail)

		return nil
	})
}
