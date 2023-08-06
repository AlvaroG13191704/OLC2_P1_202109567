package interpreter

import (
	"fmt"
	"server/parserInterpreter/parser"
)

// visit expr
func (v *Visitor) VisitOperationExpr(ctx *parser.OperationExprContext) interface{} {
	left := v.Visit(ctx.GetLeft()).(int64)   // visit the expression
	right := v.Visit(ctx.GetRight()).(int64) // visit the relation
	sign := ctx.GetOp().GetText()            // get the signal
	fmt.Println("Entra y es una operacion de: ", sign)
	// TODO: evaluate the type of the value
	switch sign {
	case "+":
		return left + right
	case "-":
		return left - right
	}
	return true
}
