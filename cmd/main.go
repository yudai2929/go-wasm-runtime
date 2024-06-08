package main

import (
	"fmt"

	"github.com/yudai2929/go-wasm-runtime/internal/wat"
	"github.com/yudai2929/go-wasm-runtime/pkg/module"
)

func main() {
	wasmBinary, err := wat.ParseStr(`(module
		(func (param i32 i32) (result i32)
		  (local.get 0)
		  (local.get 1)
		  i32.add
		)
	  )`)
	fmt.Println(wasmBinary)
	if err != nil {
		fmt.Println(err)
		return
	}

	mod, err := module.New(wasmBinary)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Magic: %v\n", mod.Magic)
	fmt.Printf("Version: %v\n", mod.Version)
	fmt.Printf("CodeSection: %v\n", mod.CodeSection)
	fmt.Printf("FunctionSection: %v\n", mod.FunctionSection)
	fmt.Printf("TypeSection: %v\n", mod.TypeSection)

}
