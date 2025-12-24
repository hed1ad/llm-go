package main

import (
	"os"
	"os/exec"
)

func runScript() error {
	cmd := exec.Command("python3", "script.py")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
