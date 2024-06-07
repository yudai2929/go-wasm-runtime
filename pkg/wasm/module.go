package wasm

type Module struct {
	Magic   string
	Version int32
}

func NewDefaultModule() *Module {
	return &Module{
		Magic:   "\x00asm",
		Version: 1,
	}
}
