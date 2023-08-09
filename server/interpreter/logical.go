package interpreter

import (
	"log"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

func (v *Visitor) VisitLogicalOperationExpr(ctx *parser.LogicalOperationExprContext) interface{} {
	// get the left value
	leftValue := v.Visit(ctx.GetLeft()).(values.PRIMITIVE)
	// get the right value
	rightValue := v.Visit(ctx.GetRight()).(values.PRIMITIVE)

	// get the operator
	sign := ctx.GetOp().GetText()

	// verify the type of the values
	switch sign {
	case "&&":
		if leftValue.GetType() == values.BooleanType && rightValue.GetType() == values.BooleanType { // Boolean && Boolean

			return &values.Boolean{Value: leftValue.GetValue().(bool) && rightValue.GetValue().(bool)}

		} else {

			// error
			log.Printf("Error: Invalid operation between '&&' '%s' and '%s' \n", leftValue.GetType(), rightValue.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Invalid operation between '&&'",
				Type:   "Semantic",
			})

			return &values.Nil{
				Value: nil,
			}

		}
	case "||":
		if leftValue.GetType() == values.BooleanType && rightValue.GetType() == values.BooleanType { // Boolean || Boolean

			return &values.Boolean{Value: leftValue.GetValue().(bool) || rightValue.GetValue().(bool)}

		} else {

			// error
			log.Printf("Error: Invalid operation between '||' '%s' and '%s' \n", leftValue.GetType(), rightValue.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Invalid operation between '||'",
				Type:   "Semantic",
			})

			return &values.Nil{
				Value: nil,
			}

		}

	}

	return &values.Nil{
		Value: nil,
	}
}
