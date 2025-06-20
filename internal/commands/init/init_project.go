package init_project

import "log"

func HandleInitProject(args []string) {
	project := ""
	if len(args) > 0 {
		project = args[0]
	}

	log.Printf("Initializing project \"%s\"...", project)
}
