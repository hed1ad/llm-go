package scripts

import (
	"fmt"
	"os"
)

func CreateScript(script string) {

	file, err := os.Create("script.py")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(script)
}
