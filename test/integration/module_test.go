package integration

import (
	"testing"

	"github.com/yudai2929/go-wasm-runtime/internal/wat"
	"github.com/zeebo/assert"
)

func Test_Module(t *testing.T) {

	wasmBinary, err := wat.ParseStr("(module)")

	assert.NoError(t, err)

	assert.Equal(t, []byte{0, 97, 115, 109, 1, 0, 0, 0}, wasmBinary)

}
