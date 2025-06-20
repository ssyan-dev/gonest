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

	if ok := utils.CreateFolderIfNotExists(project); !ok {
		log.Fatal("ERROR: already exists")
	}

	if checkIsGitExists() {
		cmdStr := "git init"
		if err := utils.ExecuteCommand(cmdStr, cmdStr, project); err != nil {
			log.Fatalf("ERROR: %v", err)
		}
	}

	initGoStr := "go mod init " + project
	if err := utils.ExecuteCommand(initGoStr, initGoStr, project); err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	if err := createCmdDirMonorepo(project); err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	utils.ClearTerminal()
	log.Printf("Project \"%s\" initialized.", project)
}

func createCmdDirMonorepo(project string) error {
	if ok := utils.CreateFolderIfNotExists(project + "/cmd"); !ok {
		log.Fatal("ERROR: already exists")
	}
	
	if err := utils.CreateFile("templates/monorepo/main.go.tmpl", project + "/cmd/main.go"); err != nil {
		return err
	}
	if err := utils.CreateFile("templates/monorepo/.gitignore.tmpl", project + "/.gitignore"); err != nil {
		return err
	}
	if err := utils.CreateFile("templates/monorepo/Makefile.tmpl", project + "/Makefile"); err != nil {
		return err
	}
	if err := utils.CreateFile("templates/README.md.tmpl", project + "/README.md"); err != nil {
		return err
	}

	return nil
}

func checkIsGitExists() bool {
	cmdStr := "git --version"
	if err := utils.ExecuteCommand(cmdStr, cmdStr, ""); err == nil {
		return true
	}
	return false
}
