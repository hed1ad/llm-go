package main

import (
	"hed1ad/llm_golang/llm"
	"hed1ad/llm_golang/scripts"
)

func main() {
	resp, err := llm.CallLLM("Напиши скрипт на python который выводит hello world!")
	if err != nil {
		panic(err)
	}

	fin := scripts.ExtractScript(resp)
	scripts.CreateScript(fin)
	scripts.RunScript()
}
