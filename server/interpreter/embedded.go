package interpreter

import (
	"fmt"
	"log"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

// visit printstmt
func (v *Visitor) VisitPrintstmt(ctx *parser.PrintstmtContext) interface{} {

	exprList := ctx.ExprList().(*parser.ExprListContext)
	expressiongs := exprList.AllExpr()

	var output string

	for _, expression := range expressiongs {
		value := v.Visit(expression).(values.PRIMITIVE) // visit the expression

		//evalue the type of the value
		switch value.GetType() {

		case values.IntType:
			fmt.Println("print value ->", value.GetValue().(int64))
			// v.Outputs = append(v.Outputs, fmt.Sprint(value.GetValue().(int64)))
			output += fmt.Sprint(value.GetValue().(int64))

		case values.FloatType:
			// add .#### to the float
			fmt.Println("print value ->", fmt.Sprintf("%.4f", value.GetValue().(float64)))
			// v.Outputs = append(v.Outputs, fmt.Sprintf("%.4f", value.GetValue().(float64)))
			output += fmt.Sprintf("%.4f", value.GetValue().(float64))

		case values.StringType:
			fmt.Println("print value ->", value.GetValue().(string))
			// v.Outputs = append(v.Outputs, fmt.Sprint(value.GetValue().(string)))
			output += fmt.Sprint(value.GetValue().(string))

		case values.BooleanType:
			fmt.Println("print value ->", value.GetValue().(bool))
			// v.Outputs = append(v.Outputs, fmt.Sprint(value.GetValue().(bool)))
			output += fmt.Sprint(value.GetValue().(bool))

		case values.CharType:
			fmt.Println("print value char ->", value.GetValue().(string))
			// v.Outputs = append(v.Outputs, fmt.Sprint(value.GetValue().(string)))
			output += fmt.Sprint(value.GetValue().(string))

		case values.NilType:
			fmt.Println("print value ->", value.GetValue())
			// v.Outputs = append(v.Outputs, fmt.Sprint(value.GetValue().(string)))
			output += "nil"

		default:
			log.Printf("Error: Invalid type '%s' to print \n", value.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("Error: Invalid type '%s' to print", value.GetType()),
				Type:   "Semantic",
			})
			return nil
		}

	}

	v.Outputs = append(v.Outputs, output)
	return nil

}
