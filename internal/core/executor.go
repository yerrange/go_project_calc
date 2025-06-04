package core

import (
	"fmt"
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

func ExecuteInstructions(instructions []model.Instruction) ([]model.PrintResult, error) {
	store := NewVariableStore()
	var wg sync.WaitGroup
	var resultsMu sync.Mutex
	var results []model.PrintResult

	declared := make(map[string]struct{})

	for _, instr := range instructions {
		switch instr.Type {
		case model.TypeCalc:
			if _, exists := declared[instr.Var]; exists {
				return nil, fmt.Errorf("variable '%s' declared multiple times", instr.Var)
			}
			declared[instr.Var] = struct{}{}

			if err := validateOperands(instr.Left, instr.Right, declared); err != nil {
				return nil, fmt.Errorf("invalid operands for '%s': %v", instr.Var, err)
			}

			if _, ok := supportedOps[instr.Op]; !ok {
				return nil, fmt.Errorf("unsupported operation: %s", instr.Op)
			}

			wg.Add(1)
			go func(inst model.Instruction) {
				defer wg.Done()
				time.Sleep(50 * time.Millisecond)

				left, err := evalOperand(inst.Left, store)
				if err != nil {
					return
				}
				right, err := evalOperand(inst.Right, store)
				if err != nil {
					return
				}

				res, err := applyOperation(inst.Op, left, right)
				if err == nil {
					_ = store.Set(inst.Var, res)
				}
			}(instr)

		case model.TypePrint:
			if _, ok := declared[instr.Var]; !ok {
				return nil, fmt.Errorf("cannot print undeclared variable: %s", instr.Var)
			}

			wg.Add(1)
			go func(inst model.Instruction) {
				defer wg.Done()
				val, err := store.Get(inst.Var)
				if err != nil {
					return
				}
				resultsMu.Lock()
				results = append(results, model.PrintResult{Var: inst.Var, Value: val})
				resultsMu.Unlock()
			}(instr)

		default:
			return nil, fmt.Errorf("unknown instruction type: %s", instr.Type)
		}
	}

	wg.Wait()
	return results, nil
}

var supportedOps = map[string]struct{}{
	"+": {},
	"-": {},
	"*": {},
}

func applyOperation(op string, a, b int64) (int64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	default:
		return 0, fmt.Errorf("invalid operator: %s", op)
	}
}

func evalOperand(op interface{}, store *VariableStore) (int64, error) {
	switch v := op.(type) {
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

func validateOperands(left, right interface{}, declared map[string]struct{}) error {
	for _, op := range []interface{}{left, right} {
		switch v := op.(type) {
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

func tryParseInt(s string) (int64, error) {
	var n int64
	_, err := fmt.Sscanf(s, "%d", &n)
	return n, err
}
