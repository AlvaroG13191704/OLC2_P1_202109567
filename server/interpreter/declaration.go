package interpreter

import (
	"fmt"
	"log"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

// Visit type declaration with value or not
func (v *Visitor) VisitTypeValueDeclaration(ctx *parser.TypeValueDeclarationContext) interface{} {
	// get the id of the variable
	varId := ctx.ID_PRIMITIVE().GetText()
	// verify if the variable is in the scope
	if v.VerifyVariableCurrentScope(varId) {
		log.Printf("Error: Variable '%s' already declared \n", varId)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Variable '%s' already declared", varId),
			Type:   "Semantic",
		})
		return nil
	}
	// get the type of the variable -> var or let
	varType := ctx.Type_declaration().GetText()
	// get the type of the value -> Int, Float, String, Boolean, Character
	varTypeValue := ctx.Type_().GetText()
	// get the value of the variable
	varValue := v.Visit(ctx.Expr()).(values.PRIMITIVE)
	// cast int to float
	if varTypeValue == "Float" && varValue.GetType() == "Int" {
		// cast the value to float
		varValue = &values.Float{
			Value: float64(varValue.GetValue().(int64)),
		}
	} else if varTypeValue != varValue.GetType() { // evaluate if the type of the value is the same as the type of the variable
		// add error
		log.Printf("Error: Type of the value '%s' is not the same as the type of the variable '%s' \n", varTypeValue, varValue.GetType())
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Type of the value '%s' is not the same as the type of the variable '%s'", varTypeValue, varType),
			Type:   "Semantic",
		})
		return nil
	}

	symbol := SymbolTable{
		Id:           varId,
		TypeSymbol:   values.Type_Variable,
		TypeVariable: varType,
		TypeData:     varTypeValue,
		Value:        varValue,
		Line:         ctx.GetStart().GetLine(),
		Column:       ctx.GetStart().GetColumn(),
	}

	// add the variable to the scope
	v.getCurrentScope()[varId] = symbol

	// apppend to the TableSymbol
	v.TableSymbol = append(v.TableSymbol, symbol)

	// print the symbol table
	// fmt.Println("Current scope or symbol table ->", v.getCurrentScope())
	// fmt.Println("Global scope or symbol table ->", v.symbolStack)

	return nil
}

// Visit type declaration with value or not
func (v *Visitor) VisitTypeOptionalValueDeclaration(ctx *parser.TypeOptionalValueDeclarationContext) interface{} {
	// get the id of the variable
	varId := ctx.ID_PRIMITIVE().GetText()
	// verify if the variable is in the scope
	if v.VerifyVariableCurrentScope(varId) {
		log.Printf("Error: Variable '%s' already declared \n", varId)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Variable '%s' already declared", varId),
			Type:   "Semantic",
		})
		return nil
	}
	// get the type of the variable -> var or let
	varType := ctx.Type_declaration().GetText()
	// get the type of the value -> Int, Float, String, Boolean, Character
	varTypeValue := ctx.Type_().GetText()
	// get the question mark
	questionMark := ctx.QUESTION_MARK()

	// evaluate if the var or let
	if varType == "var" {
		// save the value of the variable as nil
		varValue := &values.Nil{
			Value: nil,
		}
		// add the variable to the scope
		symbol := SymbolTable{
			Id:           varId,
			TypeSymbol:   values.Type_Variable,
			TypeVariable: varType,
			TypeData:     varTypeValue,
			Value:        varValue,
			Line:         ctx.GetStart().GetLine(),
			Column:       ctx.GetStart().GetColumn(),
		}

		// add the variable to the scope
		v.getCurrentScope()[varId] = symbol

		// apppend to the TableSymbol
		v.TableSymbol = append(v.TableSymbol, symbol)

	} else if varType == "let" {
		if questionMark != nil {
			log.Printf("Error: Variable '%s' is a constant and cannot be optional \n", varId)
			// add error
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("Error: Variable '%s' is a constant and cannot be optional", varId),
				Type:   "Semantic",
			})
			return nil
		}
	}

	// print the symbol table
	// fmt.Println("Current scope or symbol table ->", v.getCurrentScope())
	// fmt.Println("Global scope or symbol table ->", v.symbolStack)

	return nil
}

// Visit type declaration without type value (infer)

func (v *Visitor) VisitValueDeclaration(ctx *parser.ValueDeclarationContext) interface{} {
	// get the id of the variable
	varId := ctx.ID_PRIMITIVE().GetText()
	// verify if the variable is in the scope
	if v.VerifyVariableCurrentScope(varId) {
		log.Printf("Error: Variable '%s' already declared \n", varId)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Variable '%s' already declared", varId),
			Type:   "Semantic",
		})
		return nil
	}
	// get the type of the variable -> var or let
	varType := ctx.Type_declaration().GetText()
	// get the type of the value -> Int, Float, String, Boolean, Character
	// get the value of the variable
	varValue := v.Visit(ctx.Expr()).(values.PRIMITIVE)
	// get the type of the value
	varTypeValue := varValue.GetType()

	// add the variable to the scope
	symbol := SymbolTable{
		Id:           varId,
		TypeSymbol:   values.Type_Variable,
		TypeVariable: varType,
		TypeData:     varTypeValue,
		Value:        varValue,
		Line:         ctx.GetStart().GetLine(),
		Column:       ctx.GetStart().GetColumn(),
	}

	// add the variable to the scope
	v.getCurrentScope()[varId] = symbol

	// apppend to the TableSymbol
	v.TableSymbol = append(v.TableSymbol, symbol)

	// print the symbol table
	// fmt.Println("Current scope or symbol table ->", v.getCurrentScope())
	// fmt.Println("Global scope or symbol table ->", v.symbolStack)

	return nil

}

// VisitVectorDeclaration
func (v *Visitor) VisitVectorDeclaration(ctx *parser.VectorDeclarationContext) interface{} {
	// get the id of the variable
	varId := ctx.AllID_PRIMITIVE()[0].GetText()
	// verify if the variable is in the scope
	if v.VerifyVariableCurrentScope(varId) {
		log.Printf("Error: Variable '%s' already declared \n", varId)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Variable '%s' already declared", varId),
			Type:   "Semantic",
		})
		return nil
	}
	// get the type of the variable -> var or let
	varType := ctx.Type_declaration().GetText()
	// get the type of the value -> Int, Float, String, Boolean, Character or Vector
	varTypeValue := ctx.Type_().GetText()

	if len(ctx.AllID_PRIMITIVE()) == 2 {
		// get the symbol table of the variable
		symbol, ok := v.VerifyScope(ctx.AllID_PRIMITIVE()[1].GetText())
		// evaluate if the variable is in the scope
		if !ok {
			log.Printf("Error: Variable '%s' not declared \n", ctx.AllID_PRIMITIVE()[1].GetText())
			// add error
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("Error: Variable '%s' not declared", ctx.AllID_PRIMITIVE()[1].GetText()),
				Type:   "Semantic",
			})
			return nil
		}
		// evaluate if the type of the variable is not a vector
		symbolTable := symbol.(SymbolTable)
		if symbolTable.TypeSymbol != values.Type_Vector {
			log.Printf("Error: Variable '%s' is not a vector \n", ctx.AllID_PRIMITIVE()[1].GetText())
			// add error
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("Error: Variable '%s' is not a vector", ctx.AllID_PRIMITIVE()[1].GetText()),
				Type:   "Semantic",
			})
			return nil
		}
		// get the value of the variable
		varValue := symbolTable.Value.([]values.PRIMITIVE)
		fmt.Println("varValue ->", varValue)
		// iterate the vector
		for _, value := range varValue {
			// evaluate if the type of the value is the same as the type of the variable
			if value.GetType() != varTypeValue {
				log.Printf("Error: Type of the value '%s' is not the same as the type of the variable '%s' \n", varTypeValue, value.GetType())
				// add error
				v.Errors = append(v.Errors, Error{
					Line:   ctx.GetStart().GetLine(),
					Column: ctx.GetStart().GetColumn(),
					Msg:    fmt.Sprintf("Error: Type of the value '%s' is not the same as the type of the variable '%s'", varTypeValue, value.GetType()),
					Type:   "Semantic",
				})
				return nil
			}
		}

		// add the variable to the scope
		symbolV := SymbolTable{
			Id:           varId,
			TypeSymbol:   values.Type_Vector,
			TypeVariable: varType,
			TypeData:     varTypeValue,
			Value:        varValue,
			Line:         ctx.GetStart().GetLine(),
			Column:       ctx.GetStart().GetColumn(),
		}

		// add the variable to the scope
		v.getCurrentScope()[varId] = symbolV

		// apppend to the TableSymbol
		v.TableSymbol = append(v.TableSymbol, symbolV)

		return nil
	}

	// get the list of expresions
	exprList := ctx.ExprList().(*parser.ExprListContext)
	expressions := exprList.AllExpr()

	// create slice of interface
	var listExpr []values.PRIMITIVE
	// evaluate if expressions is empty
	if len(expressions) == 1 && v.Visit(expressions[0]) == nil {
		// add nothing

	} else {
		// iterate the expressions
		for _, expr := range expressions {
			// get the value of the expression
			value := v.Visit(expr).(values.PRIMITIVE)

			fmt.Println("expr ->", value)
			// evaluate if the type of the expression is the same as the type of the varTYpeValue
			if value.GetType() != varTypeValue {
				log.Printf("Error: Type of the value '%s' is not the same as the type of the variable '%s' \n", varTypeValue, value.GetType())
				// add error
				v.Errors = append(v.Errors, Error{
					Line:   ctx.GetStart().GetLine(),
					Column: ctx.GetStart().GetColumn(),
					Msg:    fmt.Sprintf("Error: Type of the value '%s' is not the same as the type of the variable '%s'", varTypeValue, value.GetType()),
					Type:   "Semantic",
				})
				return nil
			}
			// append the value to the list
			listExpr = append(listExpr, value)
		}

	}

	// save the variable as a vector in the scope
	symbol := SymbolTable{
		Id:           varId,
		TypeSymbol:   values.Type_Vector,
		TypeVariable: varType,
		TypeData:     varTypeValue,
		Value:        listExpr, // check this
		Line:         ctx.GetStart().GetLine(),
		Column:       ctx.GetStart().GetColumn(),
	}

	// add the variable to the scope
	v.getCurrentScope()[varId] = symbol

	// apppend to the TableSymbol
	v.TableSymbol = append(v.TableSymbol, symbol)

	// print the symbol table
	// fmt.Println("Current scope or symbol table ->", v.getCurrentScope())
	// fmt.Println("Global scope or symbol table ->", v.symbolStack)

	return nil

}
