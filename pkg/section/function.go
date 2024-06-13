package section

import (
	"github.com/yudai2929/go-wasm-runtime/pkg/instruction"
	"github.com/yudai2929/go-wasm-runtime/pkg/types"
)

type Function struct {
	Locals types.FunctionalLocals
	Code   instruction.Instructions
}

type Functions []Function
