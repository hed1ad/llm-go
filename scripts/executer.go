package scripts

import (
	"os"
	"os/exec"
)

func RunScript() error {
	cmd := exec.Command("python3", "script.py")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
