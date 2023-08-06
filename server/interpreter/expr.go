package interpreter

import "server/parserInterpreter/parser"

// visit parenexpr
func (v *Visitor) VisitParenExpr(ctx *parser.ParenExprContext) interface{} {
	return v.Visit(ctx.Expr()) // visit the expression
}

// visit negativeexpr
func (v *Visitor) VisitNegExpr(ctx *parser.NegExprContext) interface{} {
	value := v.Visit(ctx.Expr()) // visit the expression
	return -value.(float64)      // return the negative value
}
