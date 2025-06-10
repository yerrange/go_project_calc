package api

import (
	"context"
	"fmt"

	"github.com/yerrange/go_project_calc/internal/core"
	"github.com/yerrange/go_project_calc/internal/model"
	pb "github.com/yerrange/go_project_calc/proto"
)

// интерфейс для сервера
type GrpcServer struct {
	pb.UnimplementedCalculatorServer
}

// вывод пачкой
func (s *GrpcServer) Execute(ctx context.Context, req *pb.ExecuteRequest) (*pb.ExecuteResponse, error) {
	instructions := []model.Instruction{}
	for _, instr := range req.Instructions {
		mi, err := convertToModelInstruction(instr)
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, mi)
	}
	results, err := core.ExecuteInstructions(instructions)
	if err != nil {
		return nil, err
	}
	resp := &pb.ExecuteResponse{}
	for _, r := range results {
		resp.Items = append(resp.Items, &pb.PrintResult{
			Var:   r.Var,
			Value: r.Value,
		})
	}
	return resp, nil
}

// вывод по одному в потоке
func (s *GrpcServer) ExecuteStream(req *pb.ExecuteRequest, stream pb.Calculator_ExecuteStreamServer) error {
	instructions := []model.Instruction{}
	for _, instr := range req.Instructions {
		mi, err := convertToModelInstruction(instr)
		if err != nil {
			return err
		}
		instructions = append(instructions, mi)
	}
	return core.ExecuteStream(instructions, func(result model.PrintResult) {
		_ = stream.Send(&pb.PrintResult{
			Var:   result.Var,
			Value: result.Value,
		})
	})
}

// пытаемся привести к более конкретному типу
func parseEntity(val interface{}) interface{} {
	switch v := val.(type) {
	case string:
		var i int64
		if _, err := fmt.Sscan(v, &i); err == nil {
			return i
		}
		return v
	case float64:
		return int64(v)
	default:
		return v
	}
}

// функция преобразования из protobuf в типы Go
func convertToModelInstruction(instr *pb.Instruction) (model.Instruction, error) {
	var instrType model.InstructionType
	switch instr.Type {
	case "calc":
		instrType = model.TypeCalc
	case "print":
		instrType = model.TypePrint
	default:
		return model.Instruction{}, fmt.Errorf("unknown instruction type: %s", instr.Type)
	}

	return model.Instruction{
		Type:  instrType,
		Var:   instr.Var,
		Op:    instr.Op,
		Left:  parseEntity(instr.Left),
		Right: parseEntity(instr.Right),
	}, nil
}
