package interpreter

import (
	"fmt"
	"log"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

// visit expr
func (v *Visitor) VisitArithmeticOperationExpr(ctx *parser.ArithmeticOperationExprContext) interface{} {
	// get the left value
	leftValue := v.Visit(ctx.GetLeft()).(values.PRIMITIVE)
	// get the right value
	rightValue := v.Visit(ctx.GetRight()).(values.PRIMITIVE)

	sign := ctx.GetOp().GetText() // Get the operator

	// for each sing evaluate the operation and verify if the operation is valid
	switch sign {
	case "+":
		if leftValue.GetType() == values.IntType && rightValue.GetType() == values.IntType { // Int + Int

			return &values.Integer{Value: leftValue.GetValue().(int64) + rightValue.GetValue().(int64)}

		} else if leftValue.GetType() == values.IntType && rightValue.GetType() == values.FloatType { // Int + Float

			// if right  value are int cast to float
			if leftValue.GetType() == values.IntType {
				leftValue = &values.Float{
					Value: float64(leftValue.GetValue().(int64)),
				}
			}

			return &values.Float{Value: leftValue.GetValue().(float64) + rightValue.GetValue().(float64)}

		} else if leftValue.GetType() == values.FloatType && rightValue.GetType() == values.FloatType { // Float + Float

			return &values.Float{Value: leftValue.GetValue().(float64) + rightValue.GetValue().(float64)}

		} else if leftValue.GetType() == values.FloatType && rightValue.GetType() == values.IntType { // Float + Int

			// if right  value are int cast to float
			if rightValue.GetType() == values.IntType {
				rightValue = &values.Float{
					Value: float64(rightValue.GetValue().(int64)),
				}
			}

			return &values.Float{Value: leftValue.GetValue().(float64) + rightValue.GetValue().(float64)}

		} else if leftValue.GetType() == values.StringType && rightValue.GetType() == values.StringType { // String + String

			return &values.String{Value: leftValue.GetValue().(string) + rightValue.GetValue().(string)}

		} else {
			// throw an error
			log.Printf("Error: Invalid operation between '+' '%s' and '%s' \n", leftValue.GetType(), rightValue.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("Error: Invalid operation between '%s' and '%s'", leftValue.GetType(), rightValue.GetType()),
				Type:   "Semantic",
			})
		}
	//
	case "-":
		if leftValue.GetType() == values.IntType && rightValue.GetType() == values.IntType { // Int - Int

			return &values.Integer{Value: leftValue.GetValue().(int64) - rightValue.GetValue().(int64)}

		} else if leftValue.GetType() == values.IntType && rightValue.GetType() == values.FloatType { // Int - Float

			// if right  value are int cast to float
			if leftValue.GetType() == values.IntType {
				leftValue = &values.Float{
					Value: float64(leftValue.GetValue().(int64)),
				}
			}

			return &values.Float{Value: leftValue.GetValue().(float64) - rightValue.GetValue().(float64)}

		} else if leftValue.GetType() == values.FloatType && rightValue.GetType() == values.FloatType { // Float - Float

			return &values.Float{Value: leftValue.GetValue().(float64) - rightValue.GetValue().(float64)}

		} else if leftValue.GetType() == values.FloatType && rightValue.GetType() == values.IntType { // Float - Int

			// if right  value are int cast to float
			if rightValue.GetType() == values.IntType {
				rightValue = &values.Float{
					Value: float64(rightValue.GetValue().(int64)),
				}
			}

			return &values.Float{Value: leftValue.GetValue().(float64) - rightValue.GetValue().(float64)}
		} else {
			// throw an error
			log.Printf("Error: Invalid operation '-'  between '%s' and '%s' \n", leftValue.GetType(), rightValue.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("Error: Invalid operation between '%s' and '%s'", leftValue.GetType(), rightValue.GetType()),
				Type:   "Semantic",
			})
		}

	case "*":
		if leftValue.GetType() == values.IntType && rightValue.GetType() == values.IntType { // Int * Int

			return &values.Integer{Value: leftValue.GetValue().(int64) * rightValue.GetValue().(int64)}

		} else if leftValue.GetType() == values.IntType && rightValue.GetType() == values.FloatType { // Int * Float

			// if right  value are int cast to float
			if leftValue.GetType() == values.IntType {
				leftValue = &values.Float{
					Value: float64(leftValue.GetValue().(int64)),
				}
			}

			return &values.Float{Value: leftValue.GetValue().(float64) * rightValue.GetValue().(float64)}

		} else if leftValue.GetType() == values.FloatType && rightValue.GetType() == values.FloatType { // Float * Float

			return &values.Float{Value: leftValue.GetValue().(float64) * rightValue.GetValue().(float64)}

		} else if leftValue.GetType() == values.FloatType && rightValue.GetType() == values.IntType { // Float * Int

			// if right  value are int cast to float
			if rightValue.GetType() == values.IntType {
				rightValue = &values.Float{
					Value: float64(rightValue.GetValue().(int64)),
				}
			}

			return &values.Float{Value: leftValue.GetValue().(float64) * rightValue.GetValue().(float64)}

		} else {
			// throw an error
			log.Printf("Error: Invalid operation '*' between '%s' and '%s' \n", leftValue.GetType(), rightValue.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("Error: Invalid operation between '%s' and '%s'", leftValue.GetType(), rightValue.GetType()),
				Type:   "Semantic",
			})
		}

	case "/":
		// verify if the right value is not 0
		if rightValue.GetType() == values.IntType && rightValue.GetValue().(int64) == 0 || rightValue.GetType() == values.FloatType && rightValue.GetValue().(float64) == 0 {
			// throw an error
			log.Printf("Error: Division by zero \n")
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Error: Division by zero",
				Type:   "Semantic",
			})
			return &values.Nil{
				Value: nil,
			}
		}

		if leftValue.GetType() == values.IntType && rightValue.GetType() == values.IntType { // Int / Int

			return &values.Integer{Value: leftValue.GetValue().(int64) / rightValue.GetValue().(int64)}

		} else if leftValue.GetType() == values.IntType && rightValue.GetType() == values.FloatType { // Int / Float

			// if right  value are int cast to float
			if leftValue.GetType() == values.IntType {
				leftValue = &values.Float{
					Value: float64(leftValue.GetValue().(int64)),
				}
			}

			return &values.Float{Value: leftValue.GetValue().(float64) / rightValue.GetValue().(float64)}

		} else if leftValue.GetType() == values.FloatType && rightValue.GetType() == values.FloatType { // Float / Float

			return &values.Float{Value: leftValue.GetValue().(float64) / rightValue.GetValue().(float64)}

		} else if leftValue.GetType() == values.FloatType && rightValue.GetType() == values.IntType { // Float / Int

			// if right  value are int cast to float
			if rightValue.GetType() == values.IntType {
				rightValue = &values.Float{
					Value: float64(rightValue.GetValue().(int64)),
				}
			}

			return &values.Float{Value: leftValue.GetValue().(float64) / rightValue.GetValue().(float64)}

		} else {
			// throw an error
			log.Printf("Error: Invalid operation '/' between '%s' and '%s' \n", leftValue.GetType(), rightValue.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("Error: Invalid operation between '%s' and '%s'", leftValue.GetType(), rightValue.GetType()),
				Type:   "Semantic",
			})
		}

	case "%":
		// verify if the right value is not 0
		if rightValue.GetType() == values.IntType && rightValue.GetValue().(int64) == 0 || rightValue.GetType() == values.FloatType && rightValue.GetValue().(float64) == 0 {
			// throw an error
			log.Printf("Error: Division by zero \n")
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Error: Division by zero",
				Type:   "Semantic",
			})
			return &values.Nil{
				Value: nil,
			}
		}

		if leftValue.GetType() == values.IntType && rightValue.GetType() == values.IntType { // Int % Int

			return &values.Integer{Value: leftValue.GetValue().(int64) % rightValue.GetValue().(int64)}

		} else {
			// throw an error
			log.Printf("Error: Invalid operation '%%' between '%s' and '%s' \n", leftValue.GetType(), rightValue.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("Error: Invalid operation between '%s' and '%s'", leftValue.GetType(), rightValue.GetType()),
				Type:   "Semantic",
			})
		}

	default:
		// throw an error
		log.Printf("Error: Unknown operator: %s \n", sign)
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Unknown operator: %s", sign),
			Type:   "Semantic",
		})

	}
	return &values.Nil{
		Value: nil,
	}

}
