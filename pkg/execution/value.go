package execution

import "errors"

type Value interface {
	isValue()
	Add(Value) (Value, error)
}

func NewValueFromInt32(value int32) Value {
	return I32{Value: value}
}

func NewValueFromInt64(value int64) Value {
	return I64{Value: value}
}

type I32 struct {
	Value int32
}

func (I32) isValue() {}

func (v I32) Add(other Value) (Value, error) {
	if other, ok := other.(I32); ok {
		return I32{Value: v.Value + other.Value}, nil
	}
	return nil, errors.New("type mismatch")
}

type I64 struct {
	Value int64
}

func (I64) isValue() {}

func (v I64) Add(other Value) (Value, error) {
	if other, ok := other.(I64); ok {
		return I64{Value: v.Value + other.Value}, nil
	}
	return nil, errors.New("type mismatch")
}
