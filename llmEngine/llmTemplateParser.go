package llmEngine

import (
	"os"
	"bufio"
	"strings"
	"log"
)

//Write the code for ollama templatefile parsing

func ParseModelFile(FilePath string) map[string]any {
	OllamaFieldMap := make(map[string]any)
	inSystemBlock := false
	inTemplateBlock := false
        file, err := os.Open(FilePath)
	if err != nil {
		log.fatalf("Error while opening the file %v\n",err)
		OllamaFieldMap["ERR"] := err
		return OllamaFieldMap // Handle Err
	}
	var multline string
	defer file.Close()
        scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PARAMETER") {
			GetParam(OllamaFieldMap, line)
		}
		else if strings.HasPrefix(line, `SYSTEM """`) {
			inSystemBlock = true
		}
		else if strings.HasPrefix(line, "SYSTEM ") {
			OllamaFieldMap["system"] = strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(line, "SYSTEM "),`"`))
		}
		else if strings.HasPrefix(line, `TEMPLATE """`) {
                        inTemplateBlock = true
                }
                else if strings.HasPrefix(line, "TEMPLATE ") {
			OllamaFieldMap["template"] = strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(line, "TEMPLATE "),`"`))
                }
		else {
			if !strings.HasPrefix(line, `"""`) && inSystemBlock {
				multline = multline + line + "\n"
			}
			else if !strings.HasPrefix(line, `"""`) && inTemplateBlock {
				multline = multline + line + "\n"
			}
			else {
				if strings.HasPrefix(`"""`) && inSystemBlock {
					OllamaFieldMap["system"] = multline
					multline = ""
					inSystemBlock = false
				}
				if strings.HasPrefix(`"""`) && inTemplateBlock {
                                        OllamaFieldMap["template"] = multline
                                        multline = ""
                                        inTemplateBlock = false
                                }
			}
		}
	}
	return OllamaFieldMap
}
func GetParam(FieldMap map[string]any, line string) {
	trimmedLine := strings.TrimPrefix(line, "PARAMATER ")
	segment := strings.Split(trimmedLine, " ")
	FieldMap[segment[0]] = segment[1]
}
