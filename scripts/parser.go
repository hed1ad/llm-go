package scripts

import (
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

func ExtractScript(response string) string {
	log.Debug("Обрабатываем скрипт")
	re := regexp.MustCompile("(?s)```(?:python)?\n?(.*?)```")
	matches := re.FindStringSubmatch(response)

	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}
	log.Debug("Результат обработки: ", strings.TrimSpace(matches[1]))

	return ""
}
