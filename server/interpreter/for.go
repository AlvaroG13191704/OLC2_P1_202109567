package interpreter

import (
	"fmt"
	"log"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

func (v *Visitor) VisitForRangeExpr(ctx *parser.ForRangeExprContext) interface{} {
	// push a loop to the loop context
	v.PushLoopContext("for")
	defer v.PopLoopContext() // pop the loop context after the execution

	var loopVarName string = "_"
	if ctx.ID_PRIMITIVE() != nil {
		// get the loopVarName
		loopVarName = ctx.ID_PRIMITIVE().GetText()
	}
	// get the range
	rangeValue := v.Visit(ctx.ForRange())

	// fmt.Println(loopVarName, rangeValue)
	// gett he left and right value
	leftValue := rangeValue.([]values.PRIMITIVE)[0]
	rightValue := rangeValue.([]values.PRIMITIVE)[1]

	if leftValue.GetType() == values.NilType || rightValue.GetType() == values.NilType {
		log.Printf("Error: %s", "left and right value must be an integer")
		// add the error to the errors
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "left and right value in the range must be an Integer",
			Type:   "TypeError",
		})
		return nil
	}

	// assign the value to the loopVarName
	v.AssignVariable(loopVarName, SymbolTable{
		Id:           loopVarName,
		TypeSymbol:   "variable",
		TypeVariable: "let",
		TypeData:     leftValue.GetType(),
		Value:        &values.Nil{Value: nil},
		Line:         ctx.GetStart().GetLine(),
		Column:       ctx.GetStart().GetColumn(),
	})

	for i := leftValue.GetValue().(int64); i <= rightValue.GetValue().(int64); i++ {

		// assign the value to the loopVarName
		v.AssignVariable(loopVarName, SymbolTable{
			Id:           loopVarName,
			TypeSymbol:   "variable",
			TypeVariable: "let",
			TypeData:     leftValue.GetType(),
			Value: &values.Integer{
				Value: i,
			},
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		})
		// execute the statements
		v.Visit(ctx.Block())

	}

	// delete the variable from the scope
	v.DeleteVariable(loopVarName)

	// print loop context
	fmt.Println(v.loopContexts)

	return nil
}

// VisitForExpr visit forExpr
func (v *Visitor) VisitForExpr(ctx *parser.ForExprContext) interface{} {
	v.PushLoopContext("for")
	defer v.PopLoopContext() // pop the loop context after the execution
	// get the loopVarName
	loopVarName := ctx.ID_PRIMITIVE().GetText()

	// get the id from expr
	idName := ctx.Expr().GetText()

	// search the variable in the scope
	valueFromScope, ok := v.VerifyScope(idName)
	// evaluate if the name is declared
	if !ok {
		// add the error to the errors
		log.Fatalf("Error: Variable '%s' not declared", idName)
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Variable '" + idName + "' not declared",
			Type:   "VariableError",
		})
		return nil
	}
	symbolValue := valueFromScope.(SymbolTable)

	if symbolValue.TypeSymbol == values.Type_Vector {
		// get the vector
		vector := symbolValue.Value.([]values.PRIMITIVE)

		// loop the vector
		for _, value := range vector {
			// assign the value to the loopVarName
			v.AssignVariable(loopVarName, SymbolTable{
				Id:           loopVarName,
				TypeSymbol:   "variable",
				TypeVariable: "let",
				TypeData:     value.GetType(),
				Value:        &values.Nil{Value: nil},
				Line:         ctx.GetStart().GetLine(),
				Column:       ctx.GetStart().GetColumn(),
			})

			// assign the value to the loopVarName
			v.AssignVariable(loopVarName, SymbolTable{
				Id:           loopVarName,
				TypeSymbol:   "variable",
				TypeVariable: "let",
				TypeData:     value.GetType(),
				Value:        value,
				Line:         ctx.GetStart().GetLine(),
				Column:       ctx.GetStart().GetColumn(),
			})

			// execute the statements
			v.Visit(ctx.Block())
		}

		// delete the variable from the scope

		v.DeleteVariable(loopVarName)

		return nil
	}

	// get the expr
	expr := v.Visit(ctx.Expr()).(values.PRIMITIVE)

	// verify if the expr is an string
	if expr.GetType() != values.StringType {
		log.Printf("Error: %s", "expr must be an string")
		// add the error to the errors
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "expr must be an string",
			Type:   "TypeError",
		})
		return nil
	}

	// assign the value to the loopVarName
	v.AssignVariable(loopVarName, SymbolTable{
		Id:           loopVarName,
		TypeSymbol:   "variable",
		TypeVariable: "let",
		TypeData:     expr.GetType(),
		Value: &values.Nil{
			Value: nil,
		},
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	})

	// get the string
	str := expr.GetValue().(string)

	// loop the string
	for _, char := range str {

		// assign the value to the loopVarName
		v.AssignVariable(loopVarName, SymbolTable{
			Id:           loopVarName,
			TypeSymbol:   "variable",
			TypeVariable: "let",
			TypeData:     expr.GetType(),
			Value: &values.Character{
				Value: string(char),
			},
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		})

		// execute the statements
		v.Visit(ctx.Block())

	}

	// delete the variable from the scope
	v.DeleteVariable(loopVarName)

	return nil
}

// visit forRange
func (v *Visitor) VisitForRange(ctx *parser.ForRangeContext) interface{} {
	// get left expr
	leftExpr := v.Visit(ctx.GetLeft()).(values.PRIMITIVE)
	rightExpr := v.Visit(ctx.GetRight()).(values.PRIMITIVE)

	// evaluate if the left expr is an int and the right expr is an int
	if leftExpr.GetType() == values.IntType && rightExpr.GetType() == values.IntType {
		if leftExpr.GetValue().(int64) > rightExpr.GetValue().(int64) {
			// throw an error
			log.Printf("Error: %s", "left expr must be less than right expr")
			// add the error to the errors
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "left expr must be less than right expr",
				Type:   "TypeError",
			})
			return []values.PRIMITIVE{
				&values.Nil{
					Value: nil,
				},
				&values.Nil{
					Value: nil,
				},
			}
		}
		// return an array of Primitive Intenger struct
		return []values.PRIMITIVE{
			&values.Integer{
				Value: leftExpr.GetValue().(int64),
			},
			&values.Integer{
				Value: rightExpr.GetValue().(int64),
			},
		}

	} else {
		return []values.PRIMITIVE{
			&values.Nil{
				Value: nil,
			},
			&values.Nil{
				Value: nil,
			},
		}
	}
}
