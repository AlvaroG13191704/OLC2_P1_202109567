package interpreter

import (
	"fmt"
	"log"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

func (v *Visitor) VisitWhileStmt(ctx *parser.WhileStmtContext) interface{} {

	// push a loop to the loop context
	v.PushLoopContext("while")
	defer v.PopLoopContext() // pop the loop context after the execution

	// get the condition
	conditionExpr := v.Visit(ctx.Expr()).(values.PRIMITIVE)
	// verify if the condition is a Bool
	if conditionExpr.GetType() != values.BooleanType {
		// add error
		log.Printf("Error: Condition must be a boolean in while  \n")
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Error: Condition must be a boolean",
			Type:   "Semantic",
		})
		return nil
	}

	// evaluate the condition
	for conditionExpr.GetValue().(bool) {
		// check if there is a continue in the loop
		isContinue := v.GetLoopContext()
		if isContinue.ContinueFound {
			// reset the continue found
			isContinue.ContinueFound = false
			// reset the continue found
			v.UpdateLoopContext(isContinue)
			continue
		}

		// check if there is a break in the loop
		if isBreak := v.GetLoopContext().BreakFound; isBreak {
			break
		}

		// return the block
		v.Visit(ctx.Block())

		// get the condition
		conditionExpr = v.Visit(ctx.Expr()).(values.PRIMITIVE)
		// verify if the condition is a Bool
		if conditionExpr.GetType() != values.BooleanType {
			// add error
			log.Printf("Error: Condition must be a boolean in while  \n")
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Error: Condition must be a boolean",
				Type:   "Semantic",
			})
			return nil
		}

	}

	// print loop context
	fmt.Println("loop context -> ", v.loopContexts)

	return nil
}
