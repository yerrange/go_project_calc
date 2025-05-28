package model

type InstructionType string

const (
	TypeCalc  InstructionType = "calc"
	TypePrint InstructionType = "print"
)

type IncomingInstruction struct {
	Type  string      `json:"type"`
	Var   string      `json:"var"`
	Op    string      `json:"op,omitempty"`
	Left  interface{} `json:"left,omitempty"`
	Right interface{} `json:"right,omitempty"`
}

type Instruction struct {
	Type  InstructionType `json:"type"`
	Var   string          `json:"var"`
	Op    string          `json:"op,omitempty"`
	Left  any             `json:"left,omitempty"`
	Right any             `json:"right,omitempty"`
}

type PrintResult struct {
	Var   string `json:"var"`
	Value int64  `json:"value"`
}

type PrintResults struct {
	Items []PrintResult `json:"items"`
}
