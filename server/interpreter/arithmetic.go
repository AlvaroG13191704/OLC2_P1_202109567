package interpreter

import (
	"server/parserInterpreter/parser"
)

// visit expr
func (v *Visitor) VisitOperationExpr(ctx *parser.OperationExprContext) interface{} {
	return nil
}
