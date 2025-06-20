// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.0
// source: proto/calculator.proto

package calculator

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Instruction struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Op            string                 `protobuf:"bytes,2,opt,name=op,proto3" json:"op,omitempty"`
	Var           string                 `protobuf:"bytes,3,opt,name=var,proto3" json:"var,omitempty"`
	Left          string                 `protobuf:"bytes,4,opt,name=left,proto3" json:"left,omitempty"`
	Right         string                 `protobuf:"bytes,5,opt,name=right,proto3" json:"right,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Instruction) Reset() {
	*x = Instruction{}
	mi := &file_proto_calculator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Instruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instruction) ProtoMessage() {}

func (x *Instruction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instruction.ProtoReflect.Descriptor instead.
func (*Instruction) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *Instruction) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Instruction) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *Instruction) GetVar() string {
	if x != nil {
		return x.Var
	}
	return ""
}

func (x *Instruction) GetLeft() string {
	if x != nil {
		return x.Left
	}
	return ""
}

func (x *Instruction) GetRight() string {
	if x != nil {
		return x.Right
	}
	return ""
}

type ExecuteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Instructions  []*Instruction         `protobuf:"bytes,1,rep,name=instructions,proto3" json:"instructions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExecuteRequest) Reset() {
	*x = ExecuteRequest{}
	mi := &file_proto_calculator_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecuteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteRequest) ProtoMessage() {}

func (x *ExecuteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteRequest.ProtoReflect.Descriptor instead.
func (*ExecuteRequest) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *ExecuteRequest) GetInstructions() []*Instruction {
	if x != nil {
		return x.Instructions
	}
	return nil
}

type PrintResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Var           string                 `protobuf:"bytes,1,opt,name=var,proto3" json:"var,omitempty"`
	Value         int64                  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrintResult) Reset() {
	*x = PrintResult{}
	mi := &file_proto_calculator_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrintResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintResult) ProtoMessage() {}

func (x *PrintResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintResult.ProtoReflect.Descriptor instead.
func (*PrintResult) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{2}
}

func (x *PrintResult) GetVar() string {
	if x != nil {
		return x.Var
	}
	return ""
}

func (x *PrintResult) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type ExecuteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*PrintResult         `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExecuteResponse) Reset() {
	*x = ExecuteResponse{}
	mi := &file_proto_calculator_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecuteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteResponse) ProtoMessage() {}

func (x *ExecuteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteResponse.ProtoReflect.Descriptor instead.
func (*ExecuteResponse) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{3}
}

func (x *ExecuteResponse) GetItems() []*PrintResult {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_proto_calculator_proto protoreflect.FileDescriptor

const file_proto_calculator_proto_rawDesc = "" +
	"\n" +
	"\x16proto/calculator.proto\x12\n" +
	"calculator\"m\n" +
	"\vInstruction\x12\x12\n" +
	"\x04type\x18\x01 \x01(\tR\x04type\x12\x0e\n" +
	"\x02op\x18\x02 \x01(\tR\x02op\x12\x10\n" +
	"\x03var\x18\x03 \x01(\tR\x03var\x12\x12\n" +
	"\x04left\x18\x04 \x01(\tR\x04left\x12\x14\n" +
	"\x05right\x18\x05 \x01(\tR\x05right\"M\n" +
	"\x0eExecuteRequest\x12;\n" +
	"\finstructions\x18\x01 \x03(\v2\x17.calculator.InstructionR\finstructions\"5\n" +
	"\vPrintResult\x12\x10\n" +
	"\x03var\x18\x01 \x01(\tR\x03var\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x03R\x05value\"@\n" +
	"\x0fExecuteResponse\x12-\n" +
	"\x05items\x18\x01 \x03(\v2\x17.calculator.PrintResultR\x05items2\x98\x01\n" +
	"\n" +
	"Calculator\x12B\n" +
	"\aExecute\x12\x1a.calculator.ExecuteRequest\x1a\x1b.calculator.ExecuteResponse\x12F\n" +
	"\rExecuteStream\x12\x1a.calculator.ExecuteRequest\x1a\x17.calculator.PrintResult0\x01B\"Z go_project_calc/proto;calculatorb\x06proto3"

var (
	file_proto_calculator_proto_rawDescOnce sync.Once
	file_proto_calculator_proto_rawDescData []byte
)

func file_proto_calculator_proto_rawDescGZIP() []byte {
	file_proto_calculator_proto_rawDescOnce.Do(func() {
		file_proto_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_calculator_proto_rawDesc), len(file_proto_calculator_proto_rawDesc)))
	})
	return file_proto_calculator_proto_rawDescData
}

var file_proto_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_calculator_proto_goTypes = []any{
	(*Instruction)(nil),     // 0: calculator.Instruction
	(*ExecuteRequest)(nil),  // 1: calculator.ExecuteRequest
	(*PrintResult)(nil),     // 2: calculator.PrintResult
	(*ExecuteResponse)(nil), // 3: calculator.ExecuteResponse
}
var file_proto_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.ExecuteRequest.instructions:type_name -> calculator.Instruction
	2, // 1: calculator.ExecuteResponse.items:type_name -> calculator.PrintResult
	1, // 2: calculator.Calculator.Execute:input_type -> calculator.ExecuteRequest
	1, // 3: calculator.Calculator.ExecuteStream:input_type -> calculator.ExecuteRequest
	3, // 4: calculator.Calculator.Execute:output_type -> calculator.ExecuteResponse
	2, // 5: calculator.Calculator.ExecuteStream:output_type -> calculator.PrintResult
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_calculator_proto_init() }
func file_proto_calculator_proto_init() {
	if File_proto_calculator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_calculator_proto_rawDesc), len(file_proto_calculator_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_calculator_proto_goTypes,
		DependencyIndexes: file_proto_calculator_proto_depIdxs,
		MessageInfos:      file_proto_calculator_proto_msgTypes,
	}.Build()
	File_proto_calculator_proto = out.File
	file_proto_calculator_proto_goTypes = nil
	file_proto_calculator_proto_depIdxs = nil
}
