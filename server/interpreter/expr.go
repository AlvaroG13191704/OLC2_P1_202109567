package interpreter

import (
	"fmt"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

// visit idexpr
func (v *Visitor) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	id := ctx.GetText() // get the id
	fmt.Println("Id -> ", id)
	// verify if the id is in the scope or others
	value := v.VerifyScope(id).(SymbolTable) // type assertion
	// return the value
	return value.Value
}

// visit parenexpr
func (v *Visitor) VisitParenExpr(ctx *parser.ParenExprContext) interface{} {
	return v.Visit(ctx.Expr()) // visit the expression
}

// visit negativeexpr
func (v *Visitor) VisitNegExpr(ctx *parser.NegExprContext) interface{} {
	// get the value
	value := v.Visit(ctx.Expr()).(values.PRIMITIVE)
	// verify if the value is an integer or a float
	if value.GetType() == values.IntType {
		return &values.Integer{Value: -value.GetValue().(int64)} // return the negative of the value
	} else if value.GetType() == values.FloatType {
		return &values.Float{Value: -value.GetValue().(float64)} // return the negative of the value
	}
	return nil
}
