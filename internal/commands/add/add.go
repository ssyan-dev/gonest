package add

import (
	"log"
	"slices"

	"github.com/ssyan-dev/gonest/internal/commands"
	"github.com/ssyan-dev/gonest/internal/utils"
)

func HandleAdd(args []string) {
	plugins := args
	
	for _, plugin := range plugins {
		if !slices.Contains(commands.AddCommand.Args, plugin) {
			log.Fatalf("ERROR: unknown plugin '%s'", plugin)
		}

		log.Printf("Adding plugin: %s...", plugin)
		switch plugin {
		case "air":
			if err := addAir(); err != nil {
				log.Fatalf("ERROR: %v", err)
			}
		}
	}
}

func addAir() error {
	cmdStr := "go install github.com/air-verse/air@latest" // >=go1.23
	if err := utils.ExecuteCommand(cmdStr, cmdStr, "."); err != nil {
		return err
	}

	airInitCmdStr := "air init"
	if err := utils.ExecuteCommand(airInitCmdStr, airInitCmdStr, "."); err != nil {
		return err
	}

	log.Println("Plugin \"air\" added")
	return nil
}