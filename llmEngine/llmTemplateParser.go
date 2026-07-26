package llmEngine

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"log"
	"strconv"
)

//Write the code for ollama templatefile parsing

func ParseModelFile(FilePath string) OllamaRequest {
        //No need for map. Just use the Ollama objects from OllamaEngine
	ParsedOllamaReq := OllamaRequest{}
	ParsedOllamaOptions := OllamaOptions{}
	//OllamaFieldMap := make(map[string]any)
	inSystemBlock := false
	inTemplateBlock := false
        file, err := os.Open(FilePath)
	if err != nil {
		log.Fatalf("Error while opening the file %v\n",err)
		return ParsedOllamaReq
	}
	var multline string
	defer file.Close()
        scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PARAMETER") {
			GetParam(&ParsedOllamaOptions, line)
		} else if strings.HasPrefix(line, `SYSTEM """`) {
			inSystemBlock = true
		} else if strings.HasPrefix(line, "SYSTEM ") {
			ParsedOllamaReq.System = strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(line, "SYSTEM "),`"`))
		} else if strings.HasPrefix(line, `TEMPLATE """`) {
                        inTemplateBlock = true
                } else if strings.HasPrefix(line, "TEMPLATE ") {
			ParsedOllamaReq.Template = strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(line, "TEMPLATE "),`"`))
                } else if strings.HasPrefix(line, "FROM "){
			ParsedOllamaReq.Model = strings.TrimPrefix(line, "FROM ")	
		} else {
			if !strings.HasPrefix(line, `"""`) && inSystemBlock {
				multline = multline + line + "\n"
			} else if !strings.HasPrefix(line, `"""`) && inTemplateBlock {
				multline = multline + line + "\n"
			} else {
				if strings.HasPrefix(line, `"""`) && inSystemBlock {
					//OllamaFieldMap["system"] = multline
					ParsedOllamaReq.System = multline
					multline = ""
					inSystemBlock = false
				}
				if strings.HasPrefix(line, `"""`) && inTemplateBlock {
                                        //OllamaFieldMap["template"] = multline
                                        ParsedOllamaReq.Template = multline
					multline = ""
                                        inTemplateBlock = false
                                }
			}
		}
	}
	ParsedOllamaReq.Options = ParsedOllamaOptions
	return ParsedOllamaReq
}
func GetParam(opts *OllamaOptions, line string) {
	if !strings.HasPrefix(line, "PARAMETER ") {
		return
	}

	trimmedLine := strings.TrimPrefix(line, "PARAMETER ")
	segment := strings.SplitN(strings.TrimSpace(trimmedLine), " ", 2)
	
	if len(segment) < 2 {
		return
	}
	key := strings.ToLower(segment[0])
	valStr := strings.TrimSpace(segment[1])
	// Modelfile strings are often quoted
	valStr = strings.Trim(valStr, `"'`)

	switch key {
	case "num_keep":
		opts.NumKeep = parseInt(valStr)
	case "seed":
		opts.Seed = parseInt(valStr)
	case "num_predict":
		opts.NumPredict = parseInt(valStr)
	case "top_k":
		opts.TopK = parseInt(valStr)
	case "top_p":
		opts.TopP = parseFloat(valStr)
	case "min_p":
		opts.MinP = parseFloat(valStr)
	case "tfs_z":
		opts.TfsZ = parseFloat(valStr)
	case "typical_p":
		opts.TypicalP = parseFloat(valStr)
	case "repeat_last_n":
		opts.RepeatLastN = parseInt(valStr)
	case "temperature":
		opts.Temp = parseFloat(valStr)
	case "repeat_penalty":
		opts.RepeatPenalty = parseFloat(valStr)
	case "presence_penalty":
		opts.PresencePenalty = parseFloat(valStr)
	case "frequency_penalty":
		opts.FrequencyPenalty = parseFloat(valStr)
	case "mirostat":
		opts.Mirostat = parseInt(valStr)
	case "mirostat_tau":
		opts.MirostatTau = parseFloat(valStr)
	case "mirostat_eta":
		opts.MirostatEta = parseFloat(valStr)
	case "penalize_newline":
		opts.PenalizeNewline = parseBool(valStr)
	case "stop":
		opts.Stop = append(opts.Stop, valStr)
	case "numa":
		opts.Numa = parseBool(valStr)
	case "num_ctx":
		opts.Ctx = parseInt(valStr)
	case "num_batch":
		opts.NumBatch = parseInt(valStr)
	case "num_gpu":
		opts.NumGpu = parseInt(valStr)
	case "main_gpu":
		opts.MainGpu = parseInt(valStr)
	case "low_vram":
		opts.LowVram = parseBool(valStr)
	case "vocab_only":
		opts.VocabOnly = parseBool(valStr)
	case "use_mmap":
		opts.UseMmap = parseBool(valStr)
	case "use_mlock":
		opts.UseMlock = parseBool(valStr)
	case "num_thread":
		opts.NumThread = parseInt(valStr)
	}
}
func parseInt(s string) *int {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Printf("Error parsing int from %s: %v\n", s, err)
		return nil
	}
	addr := int(val)
	return &addr
}

func parseFloat(s string) *float64 {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("Error parsing float from %s: %v\n", s, err)
		return nil
	}
	return &val
}

func parseBool(s string) *bool {
	val, err := strconv.ParseBool(s)
	if err != nil {
		fmt.Printf("Error parsing bool from %s: %v\n", s, err)
		return nil
	}
	return &val
}
