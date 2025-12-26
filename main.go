package main

import (
	"hed1ad/llm_golang/infra"
	"hed1ad/llm_golang/llm"
	"hed1ad/llm_golang/scripts"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	if os.Getenv("DEBUG") != "" {
		log.SetLevel(log.DebugLevel)
		log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	report := infra.HardInfo()
	log.Debug(report)
	resp, err := llm.CallLLM("Напиши скрипт на python который выводит hello world!.")
	if err != nil {
		panic(err)
	}

	fin := scripts.ExtractScript(resp)
	scripts.CreateScript(fin)
	scripts.RunScript()
	scripts.DeleteScript()
}
