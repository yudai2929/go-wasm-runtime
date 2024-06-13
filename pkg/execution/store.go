package execution

import (
	"github.com/yudai2929/go-wasm-runtime/pkg/instruction"
	"github.com/yudai2929/go-wasm-runtime/pkg/types"
)

type Store struct {
	Funcs []FuncInst
}

type FuncInst interface {
	isFuncInst()
}

type Internal struct {
	Value InternalFuncInst
}

func (Internal) isFuncInst() {}

type InternalFuncInst struct {
	FuncType types.FuncType
	Code     Func
}

type Func struct {
	Locals types.ValueTypes
	Body   instruction.Instructions
}
