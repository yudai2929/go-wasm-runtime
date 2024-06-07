package integration

import (
	"testing"

	"github.com/yudai2929/go-wasm-runtime/internal/wat"
	"github.com/yudai2929/go-wasm-runtime/pkg/wasm"
	"github.com/zeebo/assert"
)

func Test_Module(t *testing.T) {

	wasmBinary, err := wat.ParseStr("(module)")
	assert.NoError(t, err)

	module, err := wasm.NewModule(wasmBinary)
	assert.NoError(t, err)

	assert.Equal(t, wasm.DefaultModule, module)

}
