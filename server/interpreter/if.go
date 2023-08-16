package interpreter

import (
	"log"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

func (v *Visitor) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
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
	if conditionExpr.GetValue().(bool) {
		// return the block
		return v.Visit(ctx.Block(0))

	} else if ctx.ELSE() != nil {
		// return the else block
		return v.Visit(ctx.Block(1))
	}

	return nil
}

// visit else if stmt
func (v *Visitor) VisitIfElseStmt(ctx *parser.IfElseStmtContext) interface{} {
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
	if conditionExpr.GetValue().(bool) {
		// return the block
		return v.Visit(ctx.Block())
	} else {
		// return the else block
		return v.Visit(&ctx.IfstmtContext)
	}

}
