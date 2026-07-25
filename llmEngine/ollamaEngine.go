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
        Template string `json:"template"`
}

type OllamaResponse struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response  string `json:"response"`
	Done      bool   `json:"done"`
}
type OllamaOptions struct {
        //Temp float64 `json:"temperature"`
        //Ctx int `json:"num_ctx"`
        NumKeep          *int      `json:"num_keep,omitentry"`          // Specifies the number of tokens retained from the initial prompt when the context window reaches capacity.
	Seed             *int      `json:"seed,omitentry"`              // Sets the random number generator seed. Identical seeds and parameters produce deterministic outputs.
	NumPredict       *int      `json:"num_predict,omitentry"`       // Defines the maximum number of tokens generated before stopping.
	TopK             *int      `json:"top_k,omitentry"`             // Truncates the token selection pool to the k most probable tokens.
	TopP             *float64  `json:"top_p,omitentry"`             // Truncates token selection pool to the smallest set whose cumulative probability equals or exceeds p.
	MinP             *float64  `json:"min_p,omitentry"`             // Excludes tokens whose probability is less than p relative to the most likely token.
	TfsZ             *float64  `json:"tfs_z,omitentry"`             // Eliminates tail tokens by calculating the second derivative of token probabilities.
	TypicalP         *float64  `json:"typical_p,omitentry"`         // Selects tokens whose probability aligns with the expected entropy of the sequence.
	RepeatLastN      *int      `json:"repeat_last_n,omitentry"`     // Sets the lookback window (token count) used to calculate repetition penalties.
	Temp             *float64  `json:"temperature,omitentry"`       // Scales token logits before softmax. Values < 1.0 sharpen distribution; > 1.0 flatten it.
	RepeatPenalty    *float64  `json:"repeat_penalty,omitentry"`    // Applies a multiplicative penalty to logits of tokens in the repeat_last_n window.
	PresencePenalty  *float64  `json:"presence_penalty,omitentry"`  // Applies a flat, additive penalty to tokens appearing at least once in the context.
	FrequencyPenalty *float64  `json:"frequency_penalty,omitentry"` // Applies an additive penalty scaled by the exact number of token appearances.
	Mirostat         *int      `json:"mirostat,omitentry"`          // Activates Mirostat sampling (0: off, 1: V1, 2: V2) to maintain target perplexity.
	MirostatTau      *float64  `json:"mirostat_tau,omitentry"`      // Defines the target entropy (perplexity) threshold for Mirostat.
	MirostatEta      *float64  `json:"mirostat_eta"`      // Sets the learning rate for Mirostat to adjust to target entropy.
	PenalizeNewline  *bool     `json:"penalize_newline"`  // Determines whether repetition and frequency penalties apply to the newline token.
	Stop             []string `json:"stop,omitentry"`              // Defines token sequences that halt generation. Excluded from final output.
	Numa             *bool     `json:"numa,omitentry"`              // Toggles Non-Uniform Memory Access (NUMA) optimizations for multi-socket CPUs.
	Ctx              *int      `json:"num_ctx,omitentry"`           // Sets maximum context window size in tokens, encompassing prompt and response.
	NumBatch         *int      `json:"num_batch,omitentry"`         // Sets maximum prompt tokens processed in parallel during prompt evaluation.
	NumGpu           *int      `json:"num_gpu,omitentry"`           // Defines the number of model layers offloaded to the GPU.
	MainGpu          *int      `json:"main_gpu,omitentry"`          // Identifies primary GPU index for allocating tensors/overhead in multi-GPU setups.
	LowVram          *bool     `json:"low_vram,omitentry"`          // Shifts scratch space allocation to system RAM to reduce VRAM utilization.
	VocabOnly        *bool     `json:"vocab_only,omitentry"`        // Loads only vocabulary for tokenization/detokenization, bypassing weight loading.
	UseMmap          *bool     `json:"use_mmap,omitentry"`          // Loads model file via memory mapping, reading from storage instead of full RAM load.
	UseMlock         *bool     `json:"use_mlock,omitentry"`         // Locks loaded model data in physical RAM, preventing OS paging to virtual memory.
	NumThread        *int      `json:"num_thread,omitentry"`        // Allocates the number of CPU threads utilized for compute operations.
}

func GenOllamaResp(inp_prompt string) {
	url := "http://localhost:11434/api/generate"

	filepath := "./Templates/devil-modelfile"
	reqBody := ParseModelFile(filepath)
	reqBody.Model = "llama3.2" //to be dynamically created
	reqBody.Prompt = inp_prompt
	reqBody.Stream = true
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
	fmt.Printf("\nMethod: %s\n", req.Method)
        fmt.Printf("URL: %s\n", req.URL.String())
        fmt.Printf("Headers: %v\n", req.Header)
        fmt.Printf("Body: %s\n\n", string(jsonData))

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
        // Print fields without draining the body buffer
	/* fmt.Printf("Method: %s\n", req.Method)
	fmt.Printf("URL: %s\n", req.URL.String())
        fmt.Printf("Headers: %v\n", req.Header)
        fmt.Printf("Body: %s\n", string(jsonData)) 
        */
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
