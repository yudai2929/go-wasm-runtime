package types

type FuncType struct {
	Params  ValueTypes
	Results ValueTypes
}

type FuncTypes []FuncType

type ValueTypes []ValueType

type ValueType int32

const (
	ValueTypeI32 ValueType = 0x7F
	ValueTypeI64 ValueType = 0x7E
)

func NewValueTypes(valueTypes ...uint8) []ValueType {
	vts := make([]ValueType, len(valueTypes))
	for i, vt := range valueTypes {
		vts[i] = NewValueType(vt)
	}
	return vts
}

func NewValueType(valueType uint8) ValueType {
	return ValueType(valueType)
}

type FunctionalLocal struct {
	Count uint32
	Type  ValueType
}

type FunctionalLocals []FunctionalLocal
