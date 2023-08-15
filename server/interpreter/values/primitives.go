package values

const (
	IntType     = "Int"
	FloatType   = "Float"
	StringType  = "String"
	BooleanType = "Bool"
	CharType    = "Character"
	NilType     = "nil"
)

const (
	// Types
	Type_Variable = "variable"
	Type_Function = "function"
)

type PRIMITIVE interface {
	GetValue() interface{}
	GetType() string
}
