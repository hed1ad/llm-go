package scripts

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func RunScript() error {
	log.Debug("Запускаем скрипт")
	cmd := exec.Command("python3", "script.py")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
