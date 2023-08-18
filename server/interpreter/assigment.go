package interpreter

import (
	"fmt"
	"log"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

func (v *Visitor) VisitValueAssignment(ctx *parser.ValueAssignmentContext) interface{} {

	// get the variable name
	idName := ctx.ID_PRIMITIVE().GetText()
	// get the expr
	expr := v.Visit(ctx.Expr()).(values.PRIMITIVE)
	// get the value
	valueFromScope, ok := v.VerifyScope(idName)
	// evaluate if the name is declared
	if ok {
		// get the value
		value := valueFromScope.(SymbolTable)
		// evaluate if its a constant
		if value.TypeVariable == "let" {
			log.Fatalf("Error: Variable '%s' is a constant, cannot be assign", idName)
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Variable '" + idName + "' is a constant",
				Type:   "VariableError",
			})
			return nil
		}
		// get the type of the variable

		typeVariable := value.Value.(values.PRIMITIVE).GetType()
		// evaluate if the type of the variable is the same of the expr
		if typeVariable == expr.GetType() || typeVariable == values.NilType {
			// update the value
			value.Value = expr
			// update the value no matter the scope
			v.UpdateVariable(idName, value)

			return nil
		} else {
			log.Fatalf("Error: Variable '%s' is of type %s, cannot be assign to %s", idName, typeVariable, expr.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Variable '" + idName + "' is of type " + typeVariable + ", cannot be assign to " + expr.GetType(),
				Type:   "TypeError",
			})
			return nil
		}

	} else {
		// add the error to the errors
		log.Fatalf("Error: Variable '%s' not declared", idName)
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Variable '" + idName + "' not declared",
			Type:   "VariableError",
		})
	}

	return nil

}

func (v *Visitor) VisitPlusAssignment(ctx *parser.PlusAssignmentContext) interface{} {
	// get the variable name
	idName := ctx.ID_PRIMITIVE().GetText()
	// get the expr
	expr := v.Visit(ctx.Expr()).(values.PRIMITIVE)
	// get the value
	valueFromScope, ok := v.VerifyScope(idName)
	// evaluate if the name is declared
	if ok {
		// get the value
		value := valueFromScope.(SymbolTable)
		// evaluate if its a constant
		if value.TypeVariable == "let" {
			log.Fatalf("Error: Variable '%s' is a constant, cannot be assign", idName)
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Variable '" + idName + "' is a constant",
				Type:   "VariableError",
			})
			return nil
		}
		// get the type of the variable

		typeVariable := value.Value.(values.PRIMITIVE).GetType()
		// evaluate if the type of the variable is the same of the expr
		// if the expr is an integer, cast it to float
		if typeVariable == expr.GetType() || (typeVariable == values.FloatType && expr.GetType() == values.IntType) || (typeVariable == values.IntType && expr.GetType() == values.FloatType) {

			// make the sum
			sum := makeSum(value.Value.(values.PRIMITIVE), expr)
			// update the value
			value.Value = sum
			// update the value no matter the scope
			v.UpdateVariable(idName, value)

			return nil
		} else {
			log.Fatalf("Error: Variable '%s' is of type %s, cannot be assign to %s", idName, typeVariable, expr.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Variable '" + idName + "' is of type " + typeVariable + ", cannot be assign to " + expr.GetType(),
				Type:   "TypeError",
			})
			return nil
		}

	} else {
		// add the error to the errors
		log.Fatalf("Error: Variable '%s' not declared", idName)
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Variable '" + idName + "' not declared",
			Type:   "VariableError",
		})
	}

	return nil
}

func (v *Visitor) VisitMinusAssignment(ctx *parser.MinusAssignmentContext) interface{} {
	// get the variable name
	idName := ctx.ID_PRIMITIVE().GetText()
	// get the expr
	expr := v.Visit(ctx.Expr()).(values.PRIMITIVE)
	// get the value
	valueFromScope, ok := v.VerifyScope(idName)
	// evaluate if the name is declared
	if ok {
		// get the value
		value := valueFromScope.(SymbolTable)
		// evaluate if its a constant
		if value.TypeVariable == "let" {
			log.Fatalf("Error: Variable '%s' is a constant, cannot be assign", idName)
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Variable '" + idName + "' is a constant",
				Type:   "VariableError",
			})
			return nil
		}
		// get the type of the variable

		typeVariable := value.Value.(values.PRIMITIVE).GetType()
		// evaluate if the type of the variable is the same of the expr
		// if the expr is an integer, cast it to float
		if typeVariable == expr.GetType() || (typeVariable == values.FloatType && expr.GetType() == values.IntType) || (typeVariable == values.IntType && expr.GetType() == values.FloatType) {

			// make the sum
			sum := makeRest(value.Value.(values.PRIMITIVE), expr)
			// update the value
			value.Value = sum
			// update the value no matter the scope
			v.UpdateVariable(idName, value)

			return nil
		} else {
			log.Fatalf("Error: Variable '%s' is of type %s, cannot be assign to %s", idName, typeVariable, expr.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Variable '" + idName + "' is of type " + typeVariable + ", cannot be assign to " + expr.GetType(),
				Type:   "TypeError",
			})
			return nil
		}

	} else {
		// add the error to the errors
		log.Fatalf("Error: Variable '%s' not declared", idName)
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Variable '" + idName + "' not declared",
			Type:   "VariableError",
		})
	}

	return nil
}

func makeSum(a values.PRIMITIVE, b values.PRIMITIVE) values.PRIMITIVE {
	switch a.(type) {
	case *values.Integer:
		switch b.(type) {
		case *values.Integer:
			return &values.Integer{
				Value: a.(*values.Integer).Value + b.(*values.Integer).Value,
			}
		case *values.Float:
			return &values.Float{
				Value: float64(a.(*values.Integer).Value) + b.(*values.Float).Value,
			}
		}
	case *values.Float:
		switch b.(type) {
		case *values.Integer:
			return &values.Float{
				Value: a.(*values.Float).Value + float64(b.(*values.Integer).Value),
			}
		case *values.Float:
			return &values.Float{
				Value: a.(*values.Float).Value + b.(*values.Float).Value,
			}
		}
	case *values.String:
		switch b.(type) {
		case *values.String:
			return &values.String{
				Value: a.(*values.String).Value + b.(*values.String).Value,
			}
		}
	}
	return nil
}

func makeRest(a values.PRIMITIVE, b values.PRIMITIVE) values.PRIMITIVE {
	switch a.(type) {
	case *values.Integer:
		switch b.(type) {
		case *values.Integer:
			return &values.Integer{
				Value: a.(*values.Integer).Value - b.(*values.Integer).Value,
			}
		case *values.Float:
			return &values.Float{
				Value: float64(a.(*values.Integer).Value) - b.(*values.Float).Value,
			}
		}
	case *values.Float:
		switch b.(type) {
		case *values.Integer:
			return &values.Float{
				Value: a.(*values.Float).Value - float64(b.(*values.Integer).Value),
			}
		case *values.Float:
			return &values.Float{
				Value: a.(*values.Float).Value - b.(*values.Float).Value,
			}
		}
	}
	return nil
}

// VisitVectorAssignment
func (v *Visitor) VisitVectorAssignment(ctx *parser.VectorAssignmentContext) interface{} {
	// get the variable name
	idName := ctx.ID_PRIMITIVE().GetText()
	// get the value
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

	index := int(v.Visit(ctx.Expr(0)).(values.PRIMITIVE).GetValue().(int64))
	// evaluate if the index is out of bounds
	if index < 0 || index >= len(symbolValue.Value.([]values.PRIMITIVE)) {
		log.Printf("Error: Index '%d' is out of bounds \n", index)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Index '%d' is out of bounds", index),
			Type:   "Semantic",
		})
		return nil
	}

	// get the expr
	expr := v.Visit(ctx.Expr(1)).(values.PRIMITIVE)

	// evaluate if the expr type is the same as the symbol value
	if expr.GetType() != symbolValue.TypeData {
		log.Printf("Error: Type of the expression is different from the type of the vector")
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Type of the expression is different from the type of the vector"),
			Type:   "Semantic",
		})
		return nil
	}
	fmt.Println("before assign", symbolValue.Value)
	// update the value
	symbolValue.Value.([]values.PRIMITIVE)[index] = expr

	fmt.Println("after assign", symbolValue.Value)
	// update the value no matter the scope
	v.UpdateVariable(idName, symbolValue)

	return nil
}
