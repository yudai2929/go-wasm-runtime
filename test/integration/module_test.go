package integration

import (
	"testing"

	"github.com/yudai2929/go-wasm-runtime/internal/wat"
	"github.com/yudai2929/go-wasm-runtime/pkg/module"

	"github.com/zeebo/assert"
)

func Test_Module(t *testing.T) {

	wasmBinary, err := wat.ParseStr("(module (func (param i32 i64)))")
	assert.NoError(t, err)

	mod, err := module.New(wasmBinary)

	assert.NoError(t, err)

	assert.Equal(t, module.Default, mod)

}
