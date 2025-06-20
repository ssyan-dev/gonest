package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ssyan-dev/gonest/internal/commands"
	"github.com/ssyan-dev/gonest/internal/commands/help"
	init_project "github.com/ssyan-dev/gonest/internal/commands/init"
	"github.com/ssyan-dev/gonest/internal/parser"
)



func main() {
	version := "0.0.1"
	log.SetPrefix(fmt.Sprintf("[goNest v%s] >>> ", version))
	log.SetFlags(0)

	cmds := []*commands.Command{
		commands.HelpCommand,
		commands.InitCommand,
	}

	if len(os.Args) < 2 {
		help.HandlePrintHelp(cmds)
		return
	}
	res, err := parser.ParseCommand(cmds)
	if err != nil {
		log.Fatal(err)
	}

	switch res.Command {
	case commands.HelpCommand:
		help.HandlePrintHelp(cmds)
	case commands.InitCommand:
		init_project.HandleInitProject(res.Args)
	}
}