package commands

import (
	"fmt"

	"github.com/ssyan-dev/gonest/internal/utils"
)

type Command struct {
	Aliases []string
	Args []string
	Description string
	Example string
}

var (
	HelpCommand = newCommand([]string{"help", "h"}, []string{}, "Show commands", "help")
	InitCommand = newCommand([]string{"init", "i", "n", "new"}, []string{"github.com/user/repo"}, "Init new project", "init github.com/user/repo")
	AddCommand = newCommand([]string{"add"}, []string{"air"}, "Add plugins", "add air")
)

func newCommand(aliases, args []string, desc, example string) *Command {
	return &Command {
		Aliases: aliases,
		Args: args,
		Description: desc,
		Example: example,
	}
}

func (c *Command) Help() {
	aliasStr := ""
	if len(c.Aliases) > 0 {
		aliasStr = c.Aliases[0]
		for _, a := range c.Aliases[1:] {
			aliasStr += ", " + a
		}
	}

	argsStr := ""
	if len(c.Args) > 0 {
		argsStr = " " + joinArgs(c.Args)
	}

	fmt.Printf("  %s %s | %s | example: %s\n", aliasStr, argsStr, c.Description, c.Example)
}

func joinArgs(args []string) string {
	return "[" + utils.StringJoin(args, " ") + "]"
}