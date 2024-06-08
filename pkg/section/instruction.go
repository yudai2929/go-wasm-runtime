package section

type Instruction interface {
	isInstruction()
}

type Instructions []Instruction

// End struct to represent the End instruction.
type InstructionEnd struct{}

func (e InstructionEnd) isInstruction() {}

// LocalGet struct to represent the LocalGet instruction with a u32 value.
type InstructionLocalGet struct {
	Value uint32
}

func (lg InstructionLocalGet) isInstruction() {}

// I32Add struct to represent the I32Add instruction.
type InstructionI32Add struct{}

func (ia InstructionI32Add) isInstruction() {}
