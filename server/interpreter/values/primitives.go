package values

const (
	IntType     = "Int"
	FloatType   = "Float"
	StringType  = "String"
	BooleanType = "Bool"
	CharType    = "Character"
	nilType     = "nil"
)

type PRIMITIVE interface {
	GetValue() interface{}
	GetType() string
}
