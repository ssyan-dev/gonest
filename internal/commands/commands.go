package commands

import (
	"fmt"
)

type Command struct {
	Aliases []string
	Args []string
	Description string
	Example string
}

var (
	HelpCommand = newCommand([]string{"-help", "-h"}, []string{}, "Show commands", "-help")
	InitCommand = newCommand([]string{"-init", "-i"}, []string{"github.com/user/repo"}, "Init new project", "-init github.com/user/repo")
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
	return "[" + stringJoin(args, " ") + "]"
}

func stringJoin(strs []string, sep string) string {
	result := ""
	for i, s := range strs {
		result += s
		if i != len(strs) - 1 {
			result += sep
		}
	}
	return result
}