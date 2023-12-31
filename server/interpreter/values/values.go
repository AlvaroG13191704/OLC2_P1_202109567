package values

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

// FLOAT
type Float struct {
	Value float64
}

func (f *Float) GetValue() interface{} {
	return f.Value
}

func (f *Float) GetType() string {
	return FloatType
}

// STRING
type String struct {
	Value string
}

func (s *String) GetValue() interface{} {
	return s.Value
}

func (s *String) GetType() string {
	return StringType
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

// CHARACTER
type Character struct {
	Value string
}

func (c *Character) GetValue() interface{} {
	return c.Value
}

func (c *Character) GetType() string {
	return CharType
}

// NIL
type Nil struct {
	Value interface{}
}

func (n *Nil) GetValue() interface{} {
	return n.Value
}

func (n *Nil) GetType() string {
	return NilType
}

// Vector
type Vector struct {
	Value []interface{}
}

func (v *Vector) GetValue() []interface{} {
	return v.Value
}

func (v *Vector) GetType() string {
	return VectorType
}

// Struct
type Struct struct {
	Value map[string]interface{}
}

func (s *Struct) GetValue() map[string]interface{} {
	return s.Value
}

func (s *Struct) GetType() string {
	return StructType
}
