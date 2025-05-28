package core

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/yerrange/go_project_calc/internal/model"
)

type VariableStore struct {
	mu    sync.RWMutex
	store map[string]int64
	ready map[string]*sync.Cond
}

func NewVariableStore() *VariableStore {
	return &VariableStore{
		store: make(map[string]int64),
		ready: make(map[string]*sync.Cond),
	}
}

func (vs *VariableStore) Set(varName string, value int64) error {
	vs.mu.Lock()
	defer vs.mu.Unlock()

	if _, exists := vs.store[varName]; exists {
		return fmt.Errorf("variable %s already set", varName)
	}
	vs.store[varName] = value

	if cond, ok := vs.ready[varName]; ok {
		cond.Broadcast()
	}

	return nil
}

func (vs *VariableStore) Get(varName string) (int64, error) {
	vs.mu.RLock()
	val, ok := vs.store[varName]
	vs.mu.RUnlock()
	if ok {
		return val, nil
	}

	vs.mu.Lock()
	cond, ok := vs.ready[varName]
	if !ok {
		cond = sync.NewCond(&vs.mu)
		vs.ready[varName] = cond
	}
	for {
		if val, ok := vs.store[varName]; ok {
			vs.mu.Unlock()
			return val, nil
		}
		cond.Wait()
	}
}

func ExecuteInstructions(instructions []model.Instruction) ([]model.PrintResult, error) {
	store := NewVariableStore()
	var wg sync.WaitGroup
	printResults := []model.PrintResult{}
	printMu := sync.Mutex{}

	declaredVars := map[string]struct{}{}
	for _, instr := range instructions {
		if instr.Type == model.TypeCalc {
			if _, exists := declaredVars[instr.Var]; exists {
				return nil, fmt.Errorf("variable '%s' is defined more than once", instr.Var)
			}
			declaredVars[instr.Var] = struct{}{}
		}
	}

	for _, instr := range instructions {
		if instr.Type != model.TypeCalc && instr.Type != model.TypePrint {
			return nil, fmt.Errorf("unsupported instruction type: '%s'", instr.Type)
		}

		if instr.Type == model.TypeCalc {
			if instr.Op != "+" && instr.Op != "-" && instr.Op != "*" {
				return nil, fmt.Errorf("unsupported operation '%s' for variable '%s'", instr.Op, instr.Var)
			}
			if err := validateOperand(instr.Left, declaredVars); err != nil {
				return nil, fmt.Errorf("invalid left operand for '%s': %v", instr.Var, err)
			}
			if err := validateOperand(instr.Right, declaredVars); err != nil {
				return nil, fmt.Errorf("invalid right operand for '%s': %v", instr.Var, err)
			}
		}

		if instr.Type == model.TypePrint {
			if _, ok := declaredVars[instr.Var]; !ok {
				return nil, fmt.Errorf("cannot print undeclared variable '%s'", instr.Var)
			}
		}

		switch instr.Type {
		case model.TypeCalc:
			wg.Add(1)
			go func(inst model.Instruction) {
				defer wg.Done()
				time.Sleep(50 * time.Millisecond)

				leftVal, err := resolveOperand(inst.Left, store)
				if err != nil {
					fmt.Println("error left:", err)
					return
				}
				rightVal, err := resolveOperand(inst.Right, store)
				if err != nil {
					fmt.Println("error right:", err)
					return
				}

				var res int64
				switch inst.Op {
				case "+":
					res = leftVal + rightVal
				case "-":
					res = leftVal - rightVal
				case "*":
					res = leftVal * rightVal
				}

				if err := store.Set(inst.Var, res); err != nil {
					fmt.Println("set error:", err)
					return
				}
			}(instr)

		case model.TypePrint:
			wg.Add(1)
			go func(inst model.Instruction) {
				defer wg.Done()
				val, err := store.Get(inst.Var)
				if err != nil {
					fmt.Println("print error:", err)
					return
				}
				printMu.Lock()
				printResults = append(printResults, model.PrintResult{
					Var:   inst.Var,
					Value: val,
				})
				printMu.Unlock()
			}(instr)

		default:
			return nil, errors.New("unknown instruction type")
		}
	}

	wg.Wait()
	return printResults, nil
}

func resolveOperand(op interface{}, vars *VariableStore) (int64, error) {
	switch v := op.(type) {
	case int64:
		return v, nil
	case float64:
		return int64(v), nil
	case string:
		val, err := vars.Get(v)
		if err == nil {
			return val, nil
		}
		var parsed int64
		if _, err := fmt.Sscan(v, &parsed); err == nil {
			return parsed, nil
		}
		return 0, fmt.Errorf("invalid operand: %v", v)
	default:
		return 0, fmt.Errorf("unsupported operand type: %T", v)
	}
}

func validateOperand(op interface{}, declaredVars map[string]struct{}) error {
	switch v := op.(type) {
	case int64, float64:
		return nil
	case string:
		var tmp int64
		if _, err := fmt.Sscan(v, &tmp); err == nil {
			return nil // строка — число
		}
		if _, ok := declaredVars[v]; ok {
			return nil // переменная будет определена
		}
		return fmt.Errorf("unknown variable or invalid string literal: '%s'", v)
	default:
		return fmt.Errorf("unsupported operand type: %T", v)
	}
}

func (vs *VariableStore) Exists(name string) bool {
	vs.mu.RLock()
	defer vs.mu.RUnlock()
	_, exists := vs.store[name]
	return exists
}
