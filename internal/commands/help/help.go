package help

import (
	"log"

	"github.com/ssyan-dev/gonest/internal/commands"
	"github.com/ssyan-dev/gonest/internal/utils"
)

func HandlePrintHelp(cmds []*commands.Command) {
	utils.ClearTerminal()
	log.Println("Available commands:")

	for _, cmd := range cmds {
		cmd.Help()
	}
}