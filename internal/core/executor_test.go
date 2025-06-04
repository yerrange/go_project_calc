package core

import (
	"testing"

	"github.com/yerrange/go_project_calc/internal/model"
)

// простое сложение
func TestExecuteInstructions_SimpleAddition(t *testing.T) {
	instructions := []model.Instruction{
		{Type: model.TypeCalc, Var: "x", Op: "+", Left: int64(2), Right: int64(3)},
		{Type: model.TypePrint, Var: "x"},
	}

	results, err := ExecuteInstructions(instructions)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 1 || results[0].Var != "x" || results[0].Value != 5 {
		t.Errorf("unexpected result: %+v", results)
	}
}

// неподдерживаемая операция
func TestExecuteInstructions_InvalidOp(t *testing.T) {
	instructions := []model.Instruction{
		{Type: model.TypeCalc, Var: "x", Op: "!", Left: int64(1), Right: int64(2)},
	}

	_, err := ExecuteInstructions(instructions)
	if err == nil {
		t.Error("expected error due to unsupported operation")
	}
}

// переопределение переменной
func TestExecuteInstructions_Redefinition(t *testing.T) {
	instructions := []model.Instruction{
		{Type: model.TypeCalc, Var: "x", Op: "+", Left: int64(1), Right: int64(2)},
		{Type: model.TypeCalc, Var: "x", Op: "*", Left: int64(3), Right: int64(4)},
	}

	_, err := ExecuteInstructions(instructions)
	if err == nil {
		t.Error("expected error due to variable redefinition")
	}
}

// вызов несуществующей переменной
func TestExecuteInstructions_UndeclaredPrint(t *testing.T) {
	instructions := []model.Instruction{
		{Type: model.TypePrint, Var: "missing"},
	}

	_, err := ExecuteInstructions(instructions)
	if err == nil {
		t.Error("expected error for undeclared variable in print")
	}
}

// число-строка
func TestExecuteInstructions_StringLiteralNumber(t *testing.T) {
	instructions := []model.Instruction{
		{Type: "calc", Op: "+", Var: "x", Left: "5", Right: int64(10)},
		{Type: "print", Var: "x"},
	}

	results, err := ExecuteInstructions(instructions)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 1 || results[0].Var != "x" || results[0].Value != 15 {
		t.Fatalf("unexpected result: %+v", results)
	}
}

// либо число-строка, либо существует переменная
func TestExecuteInstructions_UndeclaredStringVariable(t *testing.T) {
	instructions := []model.Instruction{
		{Type: "calc", Op: "+", Var: "x", Left: "missing_var", Right: 10},
		{Type: "print", Var: "x"},
	}

	_, err := ExecuteInstructions(instructions)
	if err == nil {
		t.Fatal("expected error due to missing variable, got none")
	}
}
