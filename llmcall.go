package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

func callLLM(prompt string) (string, error) {
	reqBody := OllamaRequest{
		Model:  "mistral:7b",
		Prompt: prompt,
		Stream: false,
	}

	jsonData, _ := json.Marshal(reqBody)

	resp, err := http.Post(
		"http://192.168.0.54:11434/api/generate",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result OllamaResponse
	json.Unmarshal(body, &result)

	return result.Response, nil
}
