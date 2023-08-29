package interpreter

import (
	"fmt"
	"log"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

// VisitStructAttribute

func (v *Visitor) VisitStructAttribute(ctx *parser.StructAttributeContext) interface{} {

	fmt.Println("access atributte")

	// get the id of the struct
	idStruct := ctx.AllID_PRIMITIVE()[0].GetText()

	// evaluate if the struct is in the scope and return it
	structCreation, ok := v.VerifyScope(idStruct)
	if !ok {
		// add the error to the errors
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Struct '" + idStruct + "' not declared",
			Type:   "StructError",
		})

		log.Printf("Error: Struct '%s' not declared", idStruct)
	}

	// fmt.Println("struct creation -> ", structCreation)

	// get the scope of the struct creation
	structScope := structCreation.(SymbolTable).Value.(map[string]SymbolTable)

	// get the id of the attribute
	idAttribute := ctx.AllID_PRIMITIVE()[1].GetText()

	// evaluate if the attribute is in the scope and return it
	attribute, ok := v.VerifyStructScopeValue(structScope, idAttribute)
	if !ok {
		// add the error to the errors
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Attribute '" + idAttribute + "' not declared or incorrect",
			Type:   "SintacticError",
		})

		log.Printf("Error: Attribute '%s' not declared", idAttribute)
	}

	fmt.Println("attribute of the struct -> ", attribute)

	// evaluate if we want to only access or change the value
	if ctx.IS_() == nil {
		// return the value
		return attribute.(SymbolTable).Value.(values.PRIMITIVE)
	}

	// change the value
	// get the expression
	expr := v.Visit(ctx.Expr()).(values.PRIMITIVE)

	// evaluate if the type of the expression is the same as the attribute
	if expr.GetType() != attribute.(SymbolTable).TypeData || attribute.(SymbolTable).TypeVariable != "var" {
		// add the error to the errors
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Expression '" + idAttribute + "' is not the same type as the attribute or it's a constant",
			Type:   "SintacticError",
		})

		log.Printf("Error: Expression '%s' is not the same type as the attribute", idAttribute)
	}

	// update the value
	structScope[idAttribute] = SymbolTable{
		Id:           attribute.(SymbolTable).Id,
		TypeSymbol:   attribute.(SymbolTable).TypeSymbol,
		TypeData:     attribute.(SymbolTable).TypeData,
		TypeVariable: attribute.(SymbolTable).TypeVariable,
		Value:        expr,
		ListParams:   attribute.(SymbolTable).ListParams,
		Mutating:     attribute.(SymbolTable).Mutating,
		Line:         attribute.(SymbolTable).Line,
		Column:       attribute.(SymbolTable).Column,
	}

	return nil
}

// VisitStructCallFunction
func (v *Visitor) VisitStructCallFunction(ctx *parser.StructCallFunctionContext) interface{} {

	fmt.Println("call function of the struct")

	return nil
}
