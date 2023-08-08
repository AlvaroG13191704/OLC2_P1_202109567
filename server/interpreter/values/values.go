package values

// VARIABLE
type Variable struct {
	Value string
}

func (v *Variable) GetValue() interface{} {
	return v.Value
}

func (v *Variable) GetType() string {
	return variableType
}

// INTEGER
type Integer struct {
	Value int64
}

func (i *Integer) GetValue() interface{} {
	return i.Value
}

func (i *Integer) GetType() string {
	return IntType
}

// BOOLEAN
type Boolean struct {
	Value bool
}

func (b *Boolean) GetValue() interface{} {
	return b.Value
}

func (b *Boolean) GetType() string {
	return BooleanType
}
