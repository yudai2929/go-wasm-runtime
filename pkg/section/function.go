package section

type Function struct {
	Locals FunctionalLocals
	Code   Instructions
}

type Functions []Function
