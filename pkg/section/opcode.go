package section

type OpCode int32

func NewOpCode(value uint8) OpCode {
	return OpCode(value)
}

const (
	OpCodeUnreachable  OpCode = 0x00
	OpCodeNop          OpCode = 0x01
	OpCodeBlock        OpCode = 0x02
	OpCodeLoop         OpCode = 0x03
	OpCodeIf           OpCode = 0x04
	OpCodeElse         OpCode = 0x05
	OpCodeEnd          OpCode = 0x0b
	OpCodeBr           OpCode = 0x0c
	OpCodeBrIf         OpCode = 0x0d
	OpCodeBrTable      OpCode = 0x0e
	OpCodeReturn       OpCode = 0x0f
	OpCodeCall         OpCode = 0x10
	OpCodeCallIndirect OpCode = 0x11
	OpCodeDrop         OpCode = 0x1a
	OpCodeSelect       OpCode = 0x1b
	OpCodeLocalGet     OpCode = 0x20
	OpCodeLocalSet     OpCode = 0x21
	OpCodeLocalTee     OpCode = 0x22
	OpCodeGlobalGet    OpCode = 0x23
	OpCodeGlobalSet    OpCode = 0x24
	OpCodeI32Load      OpCode = 0x28
	OpCodeI64Load      OpCode = 0x29
	OpCodeF32Load      OpCode = 0x2a
	OpCodeF64Load      OpCode = 0x2b
	OpCodeI32Load8S    OpCode = 0x2c
	OpCodeI32Load8U    OpCode = 0x2d
	OpCodeI32Load16S   OpCode = 0x2e
	OpCodeI32Load16U   OpCode = 0x2f
	OpCodeI64Load8S    OpCode = 0x30
	OpCodeI64Load8U    OpCode = 0x31
	OpCodeI64Load16S   OpCode = 0x32
	OpCodeI64Load16U   OpCode = 0x33
	OpCodeI64Load32S   OpCode = 0x34
	OpCodeI64Load32U   OpCode = 0x35
	OpCodeI32Store     OpCode = 0x36
	OpCodeI64Store     OpCode = 0x37
	OpCodeF32Store     OpCode = 0x38
	OpCodeF64Store     OpCode = 0x39
	OpCodeI32Store8    OpCode = 0x3a
	OpCodeI32Store16   OpCode = 0x3b
	OpCodeI64Store8    OpCode = 0x3c
	OpCodeI64Store16   OpCode = 0x3d
	OpCodeI64Store32   OpCode = 0x3e
	OpCodeMemorySize   OpCode = 0x3f
	OpCodeMemoryGrow   OpCode = 0x40
	OpCodeI32Const     OpCode = 0x41
	OpCodeI64Const     OpCode = 0x42
	OpCodeF32Const     OpCode = 0x43
	OpCodeF64Const     OpCode = 0x44
	OpCodeI32Eqz       OpCode = 0x45
	OpCodeI32Eq        OpCode = 0x46
	OpCodeI32Ne        OpCode = 0x47
	OpCodeI32LtS       OpCode = 0x48
	OpCodeI32LtU       OpCode = 0x49
	OpCodeI32GtS       OpCode = 0x4a
	OpCodeI32GtU       OpCode = 0x4b
	OpCodeI32LeS       OpCode = 0x4c
	OpCodeI32LeU       OpCode = 0x4d
	OpCodeI32GeS       OpCode = 0x4e
	OpCodeI32GeU       OpCode = 0x4f
	OpCodeI64Eqz       OpCode = 0x50
	OpCodeI64Eq        OpCode = 0x51
	OpCodeI64Ne        OpCode = 0x52
	OpCodeI64LtS       OpCode = 0x53
	OpCodeI64LtU       OpCode = 0x54
	OpCodeI64GtS       OpCode = 0x55
	OpCodeI64GtU       OpCode = 0x56
	OpCodeI64LeS       OpCode = 0x57
	OpCodeI64LeU       OpCode = 0x58
	OpCodeI64GeS       OpCode = 0x59
	OpCodeI64GeU       OpCode = 0x5a
	OpCodeF32Eq        OpCode = 0x5b
	OpCodeF32Ne        OpCode = 0x5c
	OpCodeF32Lt        OpCode = 0x5d
	OpCodeF32Gt        OpCode = 0x5e
	OpCodeF32Le        OpCode = 0x5f
	OpCodeF32Ge        OpCode = 0x60
	OpCodeF64Eq        OpCode = 0x61
	OpCodeF64Ne        OpCode = 0x62
	OpCodeF64Lt        OpCode = 0x63
	OpCodeF64Gt        OpCode = 0x64
	OpCodeF64Le        OpCode = 0x65
	OpCodeF64Ge        OpCode = 0x66
	OpCodeI32Clz       OpCode = 0x67
	OpCodeI32Ctz       OpCode = 0x68
	OpCodeI32Popcnt    OpCode = 0x69
	OpCodeI32Add       OpCode = 0x6a
	OpCodeI32Sub       OpCode = 0x6b
	OpCodeI32Mul       OpCode = 0x6c
	OpCodeI32DivS      OpCode = 0x6d
	OpCodeI32DivU      OpCode = 0x6e
	OpCodeI32RemS      OpCode = 0x6f
	OpCodeI32RemU      OpCode = 0x70
	OpCodeI32And       OpCode = 0x71
	OpCodeI32Or        OpCode = 0x72
	OpCodeI32Xor       OpCode = 0x73
	OpCodeI32Shl       OpCode = 0x74
	OpCodeI32ShrS      OpCode = 0x75
	OpCodeI32ShrU      OpCode = 0x76
	OpCodeI32Rotl      OpCode = 0x77
	OpCodeI32Rotr      OpCode = 0x78
	OpCodeI64Clz       OpCode = 0x79
	OpCodeI64Ctz       OpCode = 0x7a
	OpCodeI64Popcnt    OpCode = 0x7b
	OpCodeI64Add       OpCode = 0x7c
	OpCodeI64Sub       OpCode = 0x7d
	OpCodeI64Mul       OpCode = 0x7e
	OpCodeI64DivS      OpCode = 0x7f
	OpCodeI64DivU      OpCode = 0x80
	OpCodeI64RemS      OpCode = 0x81
	OpCodeI64RemU      OpCode = 0x82
	OpCodeI64And       OpCode = 0x83
	OpCodeI64Or        OpCode = 0x84
	OpCodeI64Xor       OpCode = 0x85
	OpCodeI64Shl       OpCode = 0x86

	// ...
)
