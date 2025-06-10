package core

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/yerrange/go_project_calc/internal/model"
)

type VariableStore struct {
	mu    sync.RWMutex
	data  map[string]int64
	ready map[string]*sync.Cond
}

func NewVariableStore() *VariableStore {
	return &VariableStore{
		data:  make(map[string]int64),
		ready: make(map[string]*sync.Cond),
	}
}

func (vs *VariableStore) Set(name string, value int64) error {
	vs.mu.Lock()
	defer vs.mu.Unlock()

	if _, exists := vs.data[name]; exists {
		return fmt.Errorf("variable %s already set", name)
	}
	vs.data[name] = value

	if cond, ok := vs.ready[name]; ok {
		cond.Broadcast()
	}
	return nil
}

func (vs *VariableStore) Get(name string) (int64, error) {
	vs.mu.RLock()
	val, ok := vs.data[name]
	vs.mu.RUnlock()
	if ok {
		return val, nil
	}

	vs.mu.Lock()
	defer vs.mu.Unlock()

	cond, ok := vs.ready[name]
	if !ok {
		cond = sync.NewCond(&vs.mu)
		vs.ready[name] = cond
	}

	for {
		if val, ok := vs.data[name]; ok {
			return val, nil
		}
		cond.Wait()
	}
}

func ExecuteInstructionsGeneric(instructions []model.Instruction, onPrint func(model.PrintResult)) error {
	store := NewVariableStore()
	var wg sync.WaitGroup
	var printWg sync.WaitGroup
	declared := make(map[string]struct{})
	errCh := make(chan error, len(instructions))

	// собираем все переменные, которые будут объявлены
	for _, instr := range instructions {
		if instr.Type == model.TypeCalc {
			if _, exists := declared[instr.Var]; exists {
				return fmt.Errorf("variable '%s' declared multiple times", instr.Var)
			}
			declared[instr.Var] = struct{}{}
		}
	}

	// проверяем, что в print-инструкциях не вызываются переменные, которые не будут рассчитаны
	for _, instr := range instructions {
		if instr.Type == model.TypePrint {
			if _, ok := declared[instr.Var]; !ok {
				return fmt.Errorf("cannot print undeclared variable: %s", instr.Var)
			}
		}
	}

	// основной цикл исполнения инструкций
	for _, instr := range instructions {
		switch instr.Type {
		case model.TypeCalc:
			if err := validateOperands(instr.Left, instr.Right, declared); err != nil {
				return fmt.Errorf("invalid operands for '%s': %v", instr.Var, err)
			}
			if _, ok := supportedOps[instr.Op]; !ok {
				return fmt.Errorf("unsupported operation: %s", instr.Op)
			}

			wg.Add(1)
			go func(inst model.Instruction) {
				defer wg.Done()
				time.Sleep(500 * time.Millisecond)
				left, err := evalOperand(inst.Left, store)
				if err != nil {
					errCh <- err
					return
				}
				right, err := evalOperand(inst.Right, store)
				if err != nil {
					errCh <- err
					return
				}
				res, err := applyOperation(inst.Op, left, right)
				if err == nil {
					_ = store.Set(inst.Var, res)
				}
			}(instr)

		case model.TypePrint:
			printWg.Add(1)
			go func(inst model.Instruction) {
				defer printWg.Done()
				time.Sleep(time.Duration(rand.Intn(5000)+5000) * time.Millisecond)
				val, err := store.Get(inst.Var)
				if err != nil {
					errCh <- fmt.Errorf("error getting %s: %v", inst.Var, err)
					return
				}
				fmt.Printf("Sending: %s = %d\n", inst.Var, val)
				onPrint(model.PrintResult{Var: inst.Var, Value: val})
			}(instr)

		default:
			return fmt.Errorf("unknown instruction type: %s", instr.Type)
		}
	}

	wg.Wait()
	printWg.Wait()

	select {
	case err := <-errCh:
		return err
	default:
		return nil
	}
}

func ExecuteInstructions(instructions []model.Instruction) ([]model.PrintResult, error) {
	var results []model.PrintResult
	err := ExecuteInstructionsGeneric(instructions, func(r model.PrintResult) {
		results = append(results, r)
	})
	return results, err
}

func ExecuteStream(instructions []model.Instruction, onPrint func(model.PrintResult)) error {
	return ExecuteInstructionsGeneric(instructions, onPrint)
}

var supportedOps = map[string]struct{}{
	"+": {},
	"-": {},
	"*": {},
}

// валидация операндов
func validateOperands(left, right interface{}, declared map[string]struct{}) error {
	for _, operand := range []interface{}{left, right} {
		switch v := operand.(type) {
		case int64, float64:
			continue
		case string:
			if _, ok := declared[v]; ok {
				continue
			}
			if _, err := tryParseInt(v); err != nil {
				return fmt.Errorf("undeclared variable or invalid literal '%s'", v)
			}
		default:
			return fmt.Errorf("unsupported operand type: %T", v)
		}
	}
	return nil
}

// преобразование операндов
func evalOperand(operand interface{}, store *VariableStore) (int64, error) {
	switch v := operand.(type) {
	case int64:
		return v, nil
	case float64:
		return int64(v), nil
	case string:
		if n, err := tryParseInt(v); err == nil {
			return n, nil
		}
		return store.Get(v)
	default:
		return 0, fmt.Errorf("unsupported operand type: %T", v)
	}
}

// доступные операции
func applyOperation(operation string, a, b int64) (int64, error) {
	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", operation)
	}
}

// s - число?
func tryParseInt(s string) (int64, error) {
	var n int64
	_, err := fmt.Sscanf(s, "%d", &n)
	return n, err
}
