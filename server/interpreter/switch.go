package interpreter

import (
	"fmt"
	"log"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

func (v *Visitor) VisitSwitchStmt(ctx *parser.SwitchStmtContext) interface{} {
	v.PushLoopContext("switch")
	defer v.PopLoopContext() // pop the loop context after the execution

	// get the condition
	conditionExpr := v.Visit(ctx.Expr()).(values.PRIMITIVE)

	// get list of cases
	casesList := ctx.AllCaseBlock()
	// iterate over the cases
	for _, caseBlock := range casesList {

		// get the case condition
		caseCondition := v.Visit(caseBlock.Expr()).(values.PRIMITIVE)
		// verify if the cases are the same type
		if caseCondition.GetType() == conditionExpr.GetType() {
			// evaluate the condition
			if caseCondition.GetValue() == conditionExpr.GetValue() {
				// return the block
				v.Visit(caseBlock.Block())
			}
		} else {
			log.Fatal("Error: switch cases must be the same type")
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Error: switch cases must be the same type",
				Type:   "Semantic",
			})
		}

	}

	// if there is a default block
	if ctx.DefaultBlock() != nil {
		// return the default block
		v.Visit(ctx.DefaultBlock())

	}

	fmt.Printf("LoopStack: %v\n", v.loopContexts)

	return nil

}

// visit default block
func (v *Visitor) VisitDefaultBlock(ctx *parser.DefaultBlockContext) interface{} {

	return v.Visit(ctx.Block())
}
