package main

func main() {
	resp, err := callLLM("Напиши скрипт на python который выводит hello world!")
	if err != nil {
		panic(err)
	}

	fin := extractScript(resp)
	createScript(fin)
	runScript()
}
