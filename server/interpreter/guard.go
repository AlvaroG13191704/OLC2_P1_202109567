package interpreter

import (
	"log"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

func (v *Visitor) VisitGuardStmt(ctx *parser.GuardStmtContext) interface{} {

	// get the condition
	conditionExpr := v.Visit(ctx.Expr()).(values.PRIMITIVE)
	// verify if the condition is a Bool
	if conditionExpr.GetType() != values.BooleanType {
		// add error
		log.Printf("Error: Condition must be a boolean in if  \n")
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Error: Condition must be a boolean",
			Type:   "Semantic",
		})
		return nil
	}

	// evaluate the condition
	if !conditionExpr.GetValue().(bool) {

		// return the block
		v.Visit(ctx.Block())

	}

	// return
	return nil
}
