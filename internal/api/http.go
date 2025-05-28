package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/yerrange/go_project_calc/internal/core"
	"github.com/yerrange/go_project_calc/internal/model"
)

// HTTPHandler godoc
// @Summary Execute calculator instructions
// @Description Accepts a list of instructions and returns calculation results
// @Tags calculator
// @Accept json
// @Produce json
// @Param instructions body []model.Instruction true "Instructions"
// @Success 200 {object} model.PrintResults
// @Failure 400 {string} string "invalid json"
// @Failure 500 {string} string "internal error"
// @Router /execute [post]
func HandleExecute(w http.ResponseWriter, r *http.Request) {
	var instructions []model.Instruction

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &instructions); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	results, err := core.ExecuteInstructions(instructions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{"items": results}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
