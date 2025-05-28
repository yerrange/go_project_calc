package core

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/yerrange/go_project_calc/internal/model"
)

func ParseInstructionsFromFile(path string) ([]model.Instruction, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read error: %w", err)
	}

	var instructions []model.Instruction
	if err := json.Unmarshal(data, &instructions); err != nil {
		return nil, fmt.Errorf("json unmarshal error: %w", err)
	}

	return instructions, nil
}
