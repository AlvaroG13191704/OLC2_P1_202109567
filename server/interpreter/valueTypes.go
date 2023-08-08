package interpreter

// add methods to cast to int, string, etc
type GenericValue struct {
	value interface{}
}

// Methods to cast to int, float, string, bool, char
func (v *GenericValue) Int() int64 {
	return v.value.(int64)
}

func (v *GenericValue) Float() float64 {
	return v.value.(float64)
}

func (v *GenericValue) String() string {
	return v.value.(string)
}

func (v *GenericValue) Bool() bool {
	return v.value.(bool)
}

func (v *GenericValue) Char() rune {
	return v.value.(rune)
}

// method to void the value
func (v *GenericValue) Void() {
	v.value = nil
}
