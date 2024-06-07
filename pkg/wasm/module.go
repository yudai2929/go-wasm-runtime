package wasm

import (
	"fmt"

	"github.com/yudai2929/go-wasm-runtime/pkg/bin"
)

type Module struct {
	Magic   string
	Version uint32
}

var (
	DefaultModule = &Module{
		Magic:   "\x00asm",
		Version: 1,
	}

	ErrInvalidModule = fmt.Errorf("invalid module")
)

func NewModule(binary []byte) (*Module, error) {
	if len(binary) < 8 {
		return nil, ErrInvalidModule
	}

	binary, err := bin.Tag(binary, []byte(DefaultModule.Magic))
	if err != nil {
		return nil, err
	}

	_, version, err := bin.LeU32(binary)
	if err != nil {
		return nil, err
	}

	return &Module{
		Magic:   DefaultModule.Magic,
		Version: version,
	}, nil
}
