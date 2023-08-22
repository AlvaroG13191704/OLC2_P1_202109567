package values

const (
	IntType       = "Int"
	FloatType     = "Float"
	StringType    = "String"
	BooleanType   = "Bool"
	CharType      = "Character"
	NilType       = "nil"
	ReferenceType = "Reference"
	VectorType    = "Vector"
	StructType    = "Struct"
)

const (
	// Types
	Type_Variable  = "variable"
	Type_Function  = "function"
	Type_Vector    = "vector"
	Type_Reference = "Reference"
	Type_Struct    = "Struct"
)

type PRIMITIVE interface {
	GetValue() interface{}
	GetType() string
}
