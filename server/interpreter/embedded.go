package interpreter

import (
	"fmt"
	"log"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
	"strconv"
	"strings"
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

			output += fmt.Sprint(value.GetValue().(int64))

		case values.FloatType:
			// add .#### to the float
			output += fmt.Sprintf("%.4f", value.GetValue().(float64))

		case values.StringType:
			// add "" to the string
			output += fmt.Sprint(value.GetValue().(string))

		case values.BooleanType:

			output += fmt.Sprint(value.GetValue().(bool))

		case values.CharType:

			output += fmt.Sprint(value.GetValue().(string))

		case values.NilType:
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

// VisitEmbeddedFunctionExpr
func (v *Visitor) VisitEmbeddedFunctionExpr(ctx *parser.EmbeddedFunctionExprContext) interface{} {
	// return the visit embbded function
	return v.Visit(ctx.EmbbededFunc())
}

// VisitIntstmt
func (v *Visitor) VisitIntstmt(ctx *parser.IntstmtContext) interface{} {
	// get expresion
	expr := v.Visit(ctx.Expr()).(values.PRIMITIVE)

	// evaluate the type
	if expr.GetType() == values.StringType {
		// evalue if the digit has a . to know if it is a float or an integer
		if strings.Contains(expr.GetValue().(string), ".") {
			// truncate the float and return an integer
			f, ok := strconv.ParseFloat(expr.GetValue().(string), 64) // convert to float
			if ok != nil {
				log.Printf("Error: Invalid type '%s' to convert to integer \n", expr.GetType())
				v.Errors = append(v.Errors, Error{
					Line:   ctx.GetStart().GetLine(),
					Column: ctx.GetStart().GetColumn(),
					Msg:    fmt.Sprintf("Error: Invalid type '%s' to convert to integer", expr.GetType()),
					Type:   "Semantic",
				})
				return &values.Nil{}
			}
			i := int64(f) // truncate the float
			// fmt.Println("Digito primitivo int: ", i)
			return &values.Integer{Value: i}
		} else {
			i, ok := strconv.ParseInt(expr.GetValue().(string), 10, 64) // convert to integer
			if ok != nil {
				log.Printf("Error: Invalid type '%s' to convert to integer \n", expr.GetType())
				v.Errors = append(v.Errors, Error{
					Line:   ctx.GetStart().GetLine(),
					Column: ctx.GetStart().GetColumn(),
					Msg:    fmt.Sprintf("Error: Invalid type '%s' to convert to integer", expr.GetType()),
					Type:   "Semantic",
				})
				return &values.Nil{}
			}
			// fmt.Println("Digito primitivo int: ", i)
			return &values.Integer{Value: i}
		}
	} else if expr.GetType() == values.FloatType {
		// truncate the float and return an integer
		f, ok := strconv.ParseFloat(fmt.Sprintf("%.4f", expr.GetValue().(float64)), 64) // convert to float
		if ok != nil {
			log.Printf("Error: Invalid type '%s' to convert to integer \n", expr.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("Error: Invalid type '%s' to convert to integer", expr.GetType()),
				Type:   "Semantic",
			})
			return &values.Nil{}
		}

		i := int64(f) // truncate the float
		// fmt.Println("Digito primitivo int: ", i)
		return &values.Integer{Value: i}
	} else {
		log.Printf("Error: Invalid type '%s' to convert to integer \n", expr.GetType())
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Invalid type '%s' to convert to integer", expr.GetType()),
			Type:   "Semantic",
		})
		return &values.Nil{}
	}

}

// VisitFloatstmt
func (v *Visitor) VisitFloatstmt(ctx *parser.FloatstmtContext) interface{} {
	// get expresion
	expr := v.Visit(ctx.Expr()).(values.PRIMITIVE)

	// evaluate the type
	if expr.GetType() == values.StringType {
		// evalue if the digit has a . to know if it is a float or an integer
		if strings.Contains(expr.GetValue().(string), ".") {
			// truncate the float and return an integer
			f, ok := strconv.ParseFloat(expr.GetValue().(string), 64) // convert to float
			if ok != nil {
				log.Printf("Error: Invalid type '%s' to convert to float \n", expr.GetType())
				v.Errors = append(v.Errors, Error{
					Line:   ctx.GetStart().GetLine(),
					Column: ctx.GetStart().GetColumn(),
					Msg:    fmt.Sprintf("Error: Invalid type '%s' to convert to float", expr.GetType()),
					Type:   "Semantic",
				})
				return &values.Nil{}
			}
			// fmt.Println("Digito primitivo float: ", f)
			return &values.Float{Value: f}
		} else {
			i, ok := strconv.ParseInt(expr.GetValue().(string), 10, 64) // convert to integer
			if ok != nil {
				log.Printf("Error: Invalid type '%s' to convert to float \n", expr.GetType())
				v.Errors = append(v.Errors, Error{
					Line:   ctx.GetStart().GetLine(),
					Column: ctx.GetStart().GetColumn(),
					Msg:    fmt.Sprintf("Error: Invalid type '%s' to convert to float", expr.GetType()),
					Type:   "Semantic",
				})
				return &values.Nil{}
			}
			// fmt.Println("Digito primitivo float: ", i)
			return &values.Float{Value: float64(i)}
		}
	} else {
		log.Printf("Error: Invalid type '%s' to convert to float \n", expr.GetType())
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Invalid type '%s' to convert to float", expr.GetType()),
			Type:   "Semantic",
		})
		return &values.Nil{}
	}
}

// VisitStringstmt
func (v *Visitor) VisitStringstmt(ctx *parser.StringstmtContext) interface{} {
	// get expresion
	expr := v.Visit(ctx.Expr()).(values.PRIMITIVE)

	// evaluate the type
	if expr.GetType() == values.StringType {
		// fmt.Println("Digito primitivo string: ", expr.GetValue().(string))
		return &values.String{Value: expr.GetValue().(string)}

	} else if expr.GetType() == values.IntType {
		// convert the int to string
		return &values.String{Value: strconv.FormatInt(expr.GetValue().(int64), 10)}
	} else if expr.GetType() == values.FloatType {
		// convert the float to string
		return &values.String{Value: fmt.Sprintf("%.4f", expr.GetValue().(float64))}
	} else if expr.GetType() == values.BooleanType {
		// convert the boolean to string
		return &values.String{Value: strconv.FormatBool(expr.GetValue().(bool))}
	} else {
		log.Printf("Error: Invalid type '%s' to convert to string \n", expr.GetType())
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Invalid type '%s' to convert to string", expr.GetType()),
			Type:   "Semantic",
		})
		return &values.Nil{}
	}
}
