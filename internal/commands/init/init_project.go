package init_project

import (
	"log"

	"github.com/ssyan-dev/gonest/internal/utils"
)

func HandleInitProject(args []string) {
	project := ""
	if len(args) > 0 {
		project = args[0]
	}

	log.Printf("Initializing project \"%s\"...", project)

	if err := utils.CreateFolderIfNotExists(project); err == nil {
		log.Fatal("ERROR: project already exists")
	}

	if checkIsGitExists() {
		cmdStr := "git init"
		if err := utils.ExecuteCommand(cmdStr, cmdStr, project); err != nil {
			log.Fatal(err)
		}
	}

	initGoStr := "go mod init " + project
	if err := utils.ExecuteCommand(initGoStr, initGoStr, project); err != nil {
		log.Fatal(err)
	}

	utils.ClearTerminal()
	log.Printf("Project \"%s\" initialized.", project)
}

func checkIsGitExists() bool {
	cmdStr := "git --version"
	if err := utils.ExecuteCommand(cmdStr, cmdStr, ""); err == nil {
		return true
	}
	return false
}
