package values

const (
	IntType       = "Int"
	FloatType     = "Float"
	StringType    = "String"
	BooleanType   = "Bool"
	CharType      = "Character"
	NilType       = "nil"
	ReferenceType = "Reference"
)

const (
	// Types
	Type_Variable  = "variable"
	Type_Function  = "function"
	Type_Reference = "Reference"
)

type PRIMITIVE interface {
	GetValue() interface{}
	GetType() string
}
