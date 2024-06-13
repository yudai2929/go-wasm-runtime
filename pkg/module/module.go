package module

import (
	"fmt"

	"github.com/yudai2929/go-wasm-runtime/pkg/bin"
	"github.com/yudai2929/go-wasm-runtime/pkg/instruction"
	"github.com/yudai2929/go-wasm-runtime/pkg/section"
	"github.com/yudai2929/go-wasm-runtime/pkg/types"
)

type Module struct {
	Magic           string
	Version         uint32
	TypeSection     *types.FuncTypes
	FunctionSection *[]uint32
	CodeSection     *section.Functions
}

var (
	Default = &Module{
		Magic:           "\x00asm",
		Version:         1,
		TypeSection:     nil,
		FunctionSection: nil,
		CodeSection:     nil,
	}

	ErrInvalidModule  = fmt.Errorf("invalid module")
	ErrInvalidMagic   = fmt.Errorf("invalid magic")
	ErrInvalidVersion = fmt.Errorf("invalid version")
	ErrInvalidSection = fmt.Errorf("invalid section")
)

func New(binary []byte) (*Module, error) {
	remaining, module, err := decode(binary)
	if err != nil {
		return nil, err
	}

	for len(remaining) > 0 {
		rest, sectionHeader, err := decodeSectionHeader(remaining)
		if err != nil {
			return nil, ErrInvalidSection
		}

		rest, sectionContent, err := bin.Take(rest, int(sectionHeader.Size))
		if err != nil {
			return nil, ErrInvalidSection
		}

		switch sectionHeader.Code {
		case section.SectionCodeType:
			_, typeSection, err := decodeTypeSection(sectionContent)
			if err != nil {
				return nil, ErrInvalidSection
			}
			module.TypeSection = &typeSection
		case section.SectionCodeFunction:
			_, functionSection, err := decodeFunctionSection(sectionContent)
			if err != nil {
				return nil, ErrInvalidSection
			}
			module.FunctionSection = &functionSection
		case section.SectionCodeCode:
			_, codeSection, err := decodeCodeSection(sectionContent)
			if err != nil {
				return nil, ErrInvalidSection
			}
			module.CodeSection = &codeSection
		}
		remaining = rest
	}

	return module, nil
}

func decode(binary []byte) (remaining []byte, module *Module, err error) {
	if len(binary) < 8 {
		return nil, nil, ErrInvalidModule
	}

	remaining, err = bin.Tag(binary, []byte(Default.Magic))
	if err != nil {
		return nil, nil, ErrInvalidMagic
	}

	remaining, version, err := bin.LeU32(remaining)
	if err != nil {
		return nil, nil, ErrInvalidVersion
	}

	module = &Module{
		Magic:           Default.Magic,
		Version:         version,
		TypeSection:     Default.TypeSection,
		FunctionSection: Default.FunctionSection,
		CodeSection:     Default.CodeSection,
	}

	return
}

func decodeSectionHeader(binary []byte) (remaining []byte, sectionHeader *section.Header, err error) {
	binary, code, err := bin.LeU8(binary)
	if err != nil {
		return nil, nil, err
	}
	binary, size, err := bin.Leb128U32(binary)
	if err != nil {
		return nil, nil, err
	}

	sectionCode := section.NewCode(code)
	return binary, section.NewHeader(sectionCode, size), nil
}

func decodeValueType(binary []byte) (remaining []byte, valueType types.ValueType, err error) {
	remaining, value, err := bin.LeU8(binary)
	if err != nil {
		return nil, 0, err
	}
	return remaining, types.NewValueType(value), nil
}

func decodeTypeSection(binary []byte) (remaining []byte, typeSection types.FuncTypes, err error) {
	funcTypes := make(types.FuncTypes, 0)
	remaining, count, err := bin.Leb128U32(binary)
	if err != nil {
		return nil, nil, err
	}

	for i := uint32(0); i < count; i++ {
		rest, _, err := bin.LeU8(remaining)
		if err != nil {
			return nil, nil, err
		}
		var funcType types.FuncType

		// Params
		rest, size, err := bin.Leb128U32(rest)
		if err != nil {
			return nil, nil, err
		}
		rest, typeBytes, err := bin.Take(rest, int(size))
		if err != nil {
			return nil, nil, err
		}
		funcType.Params = types.NewValueTypes(typeBytes...)

		// Results
		rest, size, err = bin.Leb128U32(rest)
		if err != nil {
			return nil, nil, err
		}
		rest, typeBytes, err = bin.Take(rest, int(size))
		if err != nil {
			return nil, nil, err
		}
		funcType.Results = types.NewValueTypes(typeBytes...)

		funcTypes = append(funcTypes, funcType)
		remaining = rest
	}
	return remaining, funcTypes, nil

}

func decodeFunctionSection(binary []byte) (remaining []byte, functionSection []uint32, err error) {
	funcIndexes := make([]uint32, 0)
	remaining, count, err := bin.Leb128U32(binary)
	if err != nil {
		return nil, nil, err
	}

	for i := uint32(0); i < count; i++ {
		rest, index, err := bin.Leb128U32(remaining)
		if err != nil {
			return nil, nil, err
		}
		funcIndexes = append(funcIndexes, index)
		remaining = rest
	}

	return remaining, funcIndexes, nil
}

func decodeInstruction(binary []byte) (remaining []byte, inst instruction.Instruction, err error) {
	remaining, opcode, err := bin.LeU8(binary)
	if err != nil {
		return nil, nil, err
	}
	op := section.NewOpCode(opcode)
	switch op {
	case section.OpCodeLocalGet:
		rest, index, err := bin.Leb128U32(remaining)
		if err != nil {
			return nil, nil, err
		}
		inst = instruction.InstructionLocalGet{Value: index}
		remaining = rest
	case section.OpCodeI32Add:
		inst = instruction.InstructionI32Add{}
	case section.OpCodeEnd:
		inst = instruction.InstructionEnd{}
	}

	return remaining, inst, nil
}

func decodeFunctionBody(binary []byte) (remaining []byte, functionBody section.Function, err error) {
	remaining, count, err := bin.Leb128U32(binary)
	if err != nil {
		return nil, section.Function{}, err
	}

	for i := uint32(0); i < count; i++ {
		rest, typeCount, err := bin.Leb128U32(remaining)
		if err != nil {
			return nil, section.Function{}, err
		}
		rest, valueType, err := bin.LeU8(rest)
		if err != nil {
			return nil, section.Function{}, err
		}
		local := types.FunctionalLocal{
			Count: typeCount,
			Type:  types.NewValueType(valueType),
		}
		functionBody.Locals = append(functionBody.Locals, local)

		remaining = rest
	}

	for len(remaining) > 0 {
		rest, instruction, err := decodeInstruction(remaining)
		if err != nil {
			return nil, section.Function{}, err
		}
		functionBody.Code = append(functionBody.Code, instruction)
		remaining = rest
	}

	return remaining, functionBody, nil

}

func decodeCodeSection(binary []byte) (remaining []byte, codeSection section.Functions, err error) {
	functions := make(section.Functions, 0)
	remaining, count, err := bin.Leb128U32(binary)
	if err != nil {
		return nil, nil, err
	}

	for i := uint32(0); i < count; i++ {
		rest, size, err := bin.Leb128U32(remaining)
		if err != nil {
			return nil, nil, err
		}
		rest, body, err := bin.Take(rest, int(size))
		if err != nil {
			return nil, nil, err
		}
		_, functionBody, err := decodeFunctionBody(body)
		if err != nil {
			return nil, nil, err
		}
		functions = append(functions, functionBody)
		remaining = rest
	}
	return remaining, functions, nil
}
