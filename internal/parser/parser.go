package parser

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/ssyan-dev/gonest/internal/commands"
)

type Result struct {
	Command *commands.Command
	Args    []string
}

func ParseCommand(cmds []*commands.Command) (*Result, error) {
	input := strings.ToLower(os.Args[1])

	for _, cmd := range cmds {
		if slices.Contains(cmd.Aliases, input) {
			args := os.Args[2:]
			if len(args) < len(cmd.Args) {
				return nil, fmt.Errorf("ERROR: missing arguments.\n  example usage: %s", cmd.Example)
			}

			return &Result{
				Command: cmd,
				Args:    args,
			}, nil
		}
	}

	return nil, fmt.Errorf("ERROR: unknown command '%s'", input)
}