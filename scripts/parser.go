package scripts

import (
	"regexp"
	"strings"
)

func ExtractScript(response string) string {
	re := regexp.MustCompile("(?s)```(?:python)?\n?(.*?)```")
	matches := re.FindStringSubmatch(response)

	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}

	return ""
}
