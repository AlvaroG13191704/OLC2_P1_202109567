package interpreter

import (
	"fmt"
	"server/parserInterpreter/parser"
)

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
			panic("Break statement must be inside a loop")

		}

	} else {
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Break statement must be inside a loop",
			Type:   "Semantic",
		})
		panic("Break statement must be inside a loop")
	}

	return nil
}

// VisitContinue
func (v *Visitor) VisitContinueStmt(ctx *parser.ContinueStmtContext) interface{} {

	// evaluate if the current context is inside a loop
	if v.ExistsLoopContext() {
		fmt.Println("continue found--------")
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
			panic("Continue statement must be inside a loop")

		}

	} else {
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Continue statement must be inside a loop",
			Type:   "Semantic",
		})
		panic("Continue statement must be inside a loop")
	}

	return nil
}
