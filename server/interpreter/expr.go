package interpreter

import (
	"fmt"
	"log"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

// visit idexpr
func (v *Visitor) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	id := ctx.GetText() // get the id
	fmt.Println("Id -> ", id)
	// verify if the id is in the scope or others
	variable, ok := v.VerifyScope(id)
	value := variable.(SymbolTable)
	if ok {
		// return the value
		return value.Value

	} else {
		// add the error to the errors
		log.Fatalf("Error: Variable '%s' not declared", id)
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Variable '" + id + "' not declared",
			Type:   "VariableError",
		})
	}
	return nil
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
