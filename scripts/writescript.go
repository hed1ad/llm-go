package scripts

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func CreateScript(script string) {

	log.Debug("Создаем файл скрипта")
	file, err := os.Create("script.py")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(script)
}
