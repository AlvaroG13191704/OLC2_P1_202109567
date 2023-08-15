package interpreter

import (
	"log"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

func (v *Visitor) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	// evaluate if the expression is not null

	// evaluate if the return is inside a function
	if !v.ExistsFunctionContext() {
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Return statement must be inside a function",
			Type:   "Semantic",
		})
		log.Println("Return statement must be inside a function")
		return nil
	}
	if ctx.Expr() != nil {
		// evaluate the expression
		expr := v.Visit(ctx.Expr()).(values.PRIMITIVE)

		// update the return value
		v.ReturnValue = expr
		v.IsReturn = true
		// return the expression
		return expr

	} else {
		// update the return value
		v.ReturnValue = values.Nil{Value: nil}
		v.IsReturn = true
		return nil
	}

}

// VisitBreak
func (v *Visitor) VisitBreakStmt(ctx *parser.BreakStmtContext) interface{} {

	// evaluate if the current context is inside a loop
	if v.ExistsLoopContext() {

		// update the loop context
		loopContext := v.GetLoopContext()

		if loopContext.TypeLoop == "while" || loopContext.TypeLoop == "for" || loopContext.TypeLoop == "switch" {
			// update the loop context
			loopContext.BreakFound = true

			// update the loop context
			v.UpdateLoopContext(loopContext)

		} else {
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Break statement must be inside a loop",
				Type:   "Semantic",
			})
			log.Println("Break statement must be inside a loop")

		}

	} else {
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Break statement must be inside a loop",
			Type:   "Semantic",
		})
		log.Println("Break statement must be inside a loop")
	}

	return nil
}

// VisitContinue
func (v *Visitor) VisitContinueStmt(ctx *parser.ContinueStmtContext) interface{} {

	// evaluate if the current context is inside a loop
	if v.ExistsLoopContext() {
		// fmt.Println("continue found--------")
		// update the loop context
		loopContext := v.GetLoopContext()

		if loopContext.TypeLoop == "while" || loopContext.TypeLoop == "for" {
			// update the loop context
			loopContext.ContinueFound = true

			// update the loop context
			v.UpdateLoopContext(loopContext)

		} else {
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Continue statement must be inside a loop",
				Type:   "Semantic",
			})
			log.Println("Continue statement must be inside a loop")

		}

	} else {
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Continue statement must be inside a loop",
			Type:   "Semantic",
		})
		log.Println("Continue statement must be inside a loop")
	}

	return nil
}
