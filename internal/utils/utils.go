package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func ExecuteCommand(windowsCmd, unixCmd, dir string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", windowsCmd)
	default:
		cmd = exec.Command("sh", "-c", unixCmd)
	}
	if dir != "" {
		cmd.Dir = dir
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func ClearTerminal() error {
	return ExecuteCommand("cls", "clear", "")
}

func CreateFolderIfNotExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}
	return nil
}