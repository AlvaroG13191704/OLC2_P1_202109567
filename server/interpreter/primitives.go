package interpreter

import (
	"fmt"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
	"strconv"
	"strings"
)

type PrimitiveType int

type PrimitiveValue struct {
	Value interface{}
	Type  PrimitiveType
}

// visit digitexpr
func (v *Visitor) VisitDigitExpr(ctx *parser.DigitExprContext) interface{} {
	i, _ := strconv.ParseInt(ctx.GetText(), 10, 64) // convert the string to int
	// cast

	fmt.Println("Digito primitivo: ", i)
	return &values.Integer{Value: i}
}

// visit idexpr
func (v *Visitor) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	id := ctx.GetText() // get the id
	fmt.Println("ID: ", id)
	// cast
	value := v.VerifyScope(id).(SymbolTable)
	fmt.Println("Value: ", value)
	return value.Value
}

// visit stringexpr
func (v *Visitor) VisitStringExpr(ctx *parser.StringExprContext) interface{} {
	str := strings.Trim(ctx.GetText(), "\"") // get the string
	fmt.Println("String: ", str)
	return str // return the string
}

// visit boolean
func (v *Visitor) VisitBooleanExpr(ctx *parser.BooleanExprContext) interface{} {
	value := strings.Trim(ctx.GetText(), "\"")
	return value
}
