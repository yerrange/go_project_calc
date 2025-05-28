package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/yerrange/go_project_calc/internal/core"
	"github.com/yerrange/go_project_calc/internal/model"
)

func main() {
	data, err := ioutil.ReadFile("input.json")
	if err != nil {
		log.Fatalf("failed to read input.json: %v", err)
	}

	var instructions []model.Instruction
	if err := json.Unmarshal(data, &instructions); err != nil {
		log.Fatalf("failed to unmarshal: %v", err)
	}

	results, err := core.ExecuteInstructions(instructions)
	if err != nil {
		log.Fatalf("execution error: %v", err)
	}

	out := map[string]interface{}{
		"items": results,
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(out); err != nil {
		log.Fatalf("failed to encode output: %v", err)
	}
}
