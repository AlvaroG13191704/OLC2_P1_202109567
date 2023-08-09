package interpreter

import (
	"log"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

func (v *Visitor) VisitRelationalOperationExpr(ctx *parser.RelationalOperationExprContext) interface{} {
	// get the left value
	leftValue := v.Visit(ctx.GetLeft()).(values.PRIMITIVE)
	// get the right value
	rightValue := v.Visit(ctx.GetRight()).(values.PRIMITIVE)

	// get the operator
	sign := ctx.GetOp().GetText()

	// verify the type of the values
	switch sign {
	case "<":
		if leftValue.GetType() == values.IntType && rightValue.GetType() == values.IntType { // Int < Int

			return &values.Boolean{Value: leftValue.GetValue().(int64) < rightValue.GetValue().(int64)}

		} else if leftValue.GetType() == values.FloatType && rightValue.GetType() == values.FloatType { // Float < Float

			return &values.Boolean{Value: leftValue.GetValue().(float64) < rightValue.GetValue().(float64)}

		} else if leftValue.GetType() == values.StringType && rightValue.GetType() == values.StringType { // String < String

			return &values.Boolean{Value: leftValue.GetValue().(string) < rightValue.GetValue().(string)}

		} else if leftValue.GetType() == values.CharType && rightValue.GetType() == values.CharType { // Char < Char

			return &values.Boolean{Value: leftValue.GetValue().(string) < rightValue.GetValue().(string)}

		} else {
			// error
			log.Printf("Error: Invalid operation between '<' '%s' and '%s' \n", leftValue.GetType(), rightValue.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Invalid operation between '<'",
				Type:   "Semantic",
			})
			return &values.Nil{
				Value: nil,
			}
		}
	case "<=":
		if leftValue.GetType() == values.IntType && rightValue.GetType() == values.IntType { // Int <= Int

			return &values.Boolean{Value: leftValue.GetValue().(int64) <= rightValue.GetValue().(int64)}

		} else if leftValue.GetType() == values.FloatType && rightValue.GetType() == values.FloatType { // Float <= Float

			return &values.Boolean{Value: leftValue.GetValue().(float64) <= rightValue.GetValue().(float64)}

		} else if leftValue.GetType() == values.StringType && rightValue.GetType() == values.StringType { // String <= String

			return &values.Boolean{Value: leftValue.GetValue().(string) <= rightValue.GetValue().(string)}

		} else if leftValue.GetType() == values.CharType && rightValue.GetType() == values.CharType { // Char <= Char

			return &values.Boolean{Value: leftValue.GetValue().(string) <= rightValue.GetValue().(string)}

		} else {

			// error
			log.Printf("Error: Invalid operation between '<=' '%s' and '%s' \n", leftValue.GetType(), rightValue.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Invalid operation between '<='",
				Type:   "Semantic",
			})

			return &values.Nil{
				Value: nil,
			}

		}
	case ">":
		if leftValue.GetType() == values.IntType && rightValue.GetType() == values.IntType { // Int > Int

			return &values.Boolean{Value: leftValue.GetValue().(int64) > rightValue.GetValue().(int64)}

		} else if leftValue.GetType() == values.FloatType && rightValue.GetType() == values.FloatType { // Float > Float

			return &values.Boolean{Value: leftValue.GetValue().(float64) > rightValue.GetValue().(float64)}

		} else if leftValue.GetType() == values.StringType && rightValue.GetType() == values.StringType { // String > String

			return &values.Boolean{Value: leftValue.GetValue().(string) > rightValue.GetValue().(string)}

		} else if leftValue.GetType() == values.CharType && rightValue.GetType() == values.CharType { // Char > Char

			return &values.Boolean{Value: leftValue.GetValue().(string) > rightValue.GetValue().(string)}

		} else {

			// error
			log.Printf("Error: Invalid operation between '>' '%s' and '%s' \n", leftValue.GetType(), rightValue.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Invalid operation between '>'",
				Type:   "Semantic",
			})

			return &values.Nil{
				Value: nil,
			}
		}

	case ">=":
		if leftValue.GetType() == values.IntType && rightValue.GetType() == values.IntType { // Int >= Int

			return &values.Boolean{Value: leftValue.GetValue().(int64) >= rightValue.GetValue().(int64)}

		} else if leftValue.GetType() == values.FloatType && rightValue.GetType() == values.FloatType { // Float >= Float

			return &values.Boolean{Value: leftValue.GetValue().(float64) >= rightValue.GetValue().(float64)}

		} else if leftValue.GetType() == values.StringType && rightValue.GetType() == values.StringType { // String >= String

			return &values.Boolean{Value: leftValue.GetValue().(string) >= rightValue.GetValue().(string)}

		} else if leftValue.GetType() == values.CharType && rightValue.GetType() == values.CharType { // Char >= Char

			return &values.Boolean{Value: leftValue.GetValue().(string) >= rightValue.GetValue().(string)}

		} else {

			// error
			log.Printf("Error: Invalid operation between '>=' '%s' and '%s' \n", leftValue.GetType(), rightValue.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Invalid operation between '>='",
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
