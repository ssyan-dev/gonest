package utils

import (
	"os"
	"os/exec"
	"runtime"
	"text/template"
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

func CreateFolderIfNotExists(dir string) bool {
	_, err := os.Stat(dir)
	if err == nil {
		return false
	}
	if os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
		return true
	}

	return false
}

func CreateFile(tmplPath, path string) error {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, nil)
}