package values

const (
	IntType      = "Int"
	FloatType    = "Float"
	StringType   = "String"
	BooleanType  = "Boolean"
	CharType     = "Character"
	nilType      = "nil"
	variableType = "Variable"
)

type PRIMITIVE interface {
	GetValue() interface{}
	GetType() string
}
