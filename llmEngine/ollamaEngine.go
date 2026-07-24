package llmEngine

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
        System string `json:"system"`
        Options OllamaOptions `json:"options"`
}

type OllamaResponse struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response  string `json:"response"`
	Done      bool   `json:"done"`
}
type OllamaOptions struct {
        temp float64 `json:"temperature"`
        ctx int `json:"num_ctx"`
}

func GenOllamaResp() {
	url := "http://localhost:11434/api/generate"
	reqOptions := OllamaOptions {
                temp: 0.25,
                ctx : 2048,
        }

	reqBody := OllamaRequest{
		Model:  "llama3.2",
		Prompt: "Explain differential calculus in one sentence.",
		Stream: true,
                System: "You are a helpful assistant\n",
                Options: reqOptions,
	} 
	
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "JSON marshal error: %v\n", err)
		os.Exit(1)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Request creation error: %v\n", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "HTTP request error: %v\n", err)
		os.Exit(1)
	}
	// Defers must be placed immediately after error checking to prevent resource leaks.
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Unexpected status code: %d\n", resp.StatusCode)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Bytes()
		
		var chunk OllamaResponse
		if err := json.Unmarshal(line, &chunk); err != nil {
			fmt.Fprintf(os.Stderr, "\nJSON unmarshal error on chunk: %v\n", err)
			continue
		}

		fmt.Print(chunk.Response)

		if chunk.Done {
			break
		}
	}

	// Scanner errors must be checked after the loop terminates.
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "\nStream reading error: %v\n", err)
	}
	
	fmt.Println()
}
