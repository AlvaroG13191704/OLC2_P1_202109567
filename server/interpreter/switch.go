package interpreter

import (
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

func (v *Visitor) VisitSwitchStmt(ctx *parser.SwitchStmtContext) interface{} {
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
				return v.Visit(caseBlock.Block())
			}
		}

	}

	// if there is a default block
	if ctx.DefaultBlock() != nil {
		// return the default block
		v.Visit(ctx.DefaultBlock())

	}

	return nil

}

// visit default block
func (v *Visitor) VisitDefaultBlock(ctx *parser.DefaultBlockContext) interface{} {
	return v.Visit(ctx.Block())
}
