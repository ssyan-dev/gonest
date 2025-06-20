package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearTerminal() error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cls")
	default:
		cmd = exec.Command("clear")
	}
	
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
