package interpreter

import (
	"fmt"
	"log"
	"reflect"
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

	// fmt.Println("attribute of the struct -> ", attribute)

	// evaluate if we want to only access or change the value
	if ctx.IS_() == nil {
		// get attribute
		assertAttribute := attribute.(SymbolTable).Value

		// assert
		if reflect.TypeOf(assertAttribute).Kind() == reflect.Map {
			// get the third id
			idThird := ctx.AllID_PRIMITIVE()[2].GetText()
			// fmt.Println("idThird -> ", idThird)
			// get the value
			value := assertAttribute.(map[string]SymbolTable)
			// fmt.Println("value -> ", value)
			// fmt.Println("value -> ", value[idThird])

			if value[idThird].Value == nil {
				// add the error to the errors
				v.Errors = append(v.Errors, Error{
					Line:   ctx.GetStart().GetLine(),
					Column: ctx.GetStart().GetColumn(),
					Msg:    "Attribute '" + idThird + "' not declared or incorrect",
					Type:   "SintacticError",
				})

				log.Printf("Error: Attribute '%s' not declared", idThird)
				return values.Nil{Value: nil}
			}

			return value[idThird].Value.(values.PRIMITIVE)

		}
		// return the value
		return attribute.(SymbolTable).Value.(values.PRIMITIVE)
	}

	// change the value
	assert := v.Visit(ctx.Expr())
	if reflect.TypeOf(assert).Kind() == reflect.Map {
		symbolExpr := assert.(map[string]SymbolTable)
		// fmt.Println("symbolExpr -> ", symbolExpr)
		// assign the value to the attribute
		structScope[idAttribute] = SymbolTable{
			Id:           attribute.(SymbolTable).Id,
			TypeSymbol:   attribute.(SymbolTable).TypeSymbol,
			TypeData:     attribute.(SymbolTable).TypeData,
			TypeVariable: attribute.(SymbolTable).TypeVariable,
			Value:        symbolExpr,
			ListParams:   attribute.(SymbolTable).ListParams,
			Mutating:     attribute.(SymbolTable).Mutating,
			Line:         attribute.(SymbolTable).Line,
			Column:       attribute.(SymbolTable).Column,
			StructOf:     structCreation.(SymbolTable).Id,
		}
		return nil
	}
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
			Type:   "SintacticError",
		})

		log.Printf("Error: Struct '%s' not declared", idStruct)
		return values.Nil{}
	}

	// get the scope of the struct creation
	structScope := structCreation.(SymbolTable).Value.(map[string]SymbolTable)

	// get the id of the function
	idFunction := ctx.AllID_PRIMITIVE()[1].GetText()

	// evaluate if the function is in the scope and return it
	function, ok := v.VerifyStructScopeValue(structScope, idFunction)
	if !ok {
		// add the error to the errors
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Function '" + idFunction + "' not declared or incorrect",
			Type:   "SintacticError",
		})

		log.Printf("Error: Function '%s' not declared", idFunction)
	}

	// fmt.Println("function of the struct -> ", function)

	// active function in the self stack
	v.ActiveFunctionInSelfStack(idStruct, idFunction)
	fmt.Println("self stack before -> ", v.SelfStructs)

	v.Visit(function.(SymbolTable).Value.(parser.IBlockContext))

	// deactivate function in the self stack
	v.DeactiveFunctionInSelfStack(idStruct)
	fmt.Println("self stack after -> ", v.SelfStructs)

	// change the return value
	v.IsReturn = false
	fmt.Println("VisitCallFunctionExpr", v.ReturnValue)

	return v.ReturnValue
}

// VisitSelfFunction
func (v *Visitor) VisitSelfFunction(ctx *parser.SelfFunctionContext) interface{} {

	fmt.Println("self function")

	// get the atribute
	attribute := ctx.ID_PRIMITIVE().GetText()

	fmt.Println(v.SelfStructs)
	// // get the father

	symbolSelf, ok, functionName := v.VerifySelfStruct(attribute)
	if !ok {
		// add the error to the errors
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "self '" + attribute + "' out of struct",
			Type:   "SintacticError",
		})

		log.Printf("Error: Self '%s' out of struct", attribute)
	}

	// get the scope of the struct creation
	structScope := symbolSelf.(SymbolTable).Value.(map[string]SymbolTable)
	// get the function
	function, _ := v.VerifyStructScopeValue(structScope, functionName)
	// get the atttibute
	attributeSymbol, _ := v.VerifyStructScopeValue(structScope, attribute)

	// fmt.Println("function who's calling-> ", functionName)
	// fmt.Println("function of the struct -> ", function)
	// fmt.Println("attribute of the struct -> ", attributeSymbol)

	// verify if the function is mutating
	if function.(SymbolTable).Mutating {
		// evaluate if we want to only access or change the value
		if ctx.IS_() == nil {
			// return the value
			return attributeSymbol.(SymbolTable).Value.(values.PRIMITIVE)
		}

		// change the value
		// get the expression
		expr := v.Visit(ctx.Expr()).(values.PRIMITIVE)

		// evaluate if the type of the expression is the same as the attribute
		if expr.GetType() != attributeSymbol.(SymbolTable).TypeData || attributeSymbol.(SymbolTable).TypeVariable != "var" {
			// add the error to the errors
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Expression '" + attribute + "' is not the same type as the attribute or it's a constant",
				Type:   "SintacticError",
			})

			log.Printf("Error: Expression '%s' is not the same type as the attribute", attribute)
		}

		// update the value
		structScope[attribute] = SymbolTable{
			Id:           attributeSymbol.(SymbolTable).Id,
			TypeSymbol:   attributeSymbol.(SymbolTable).TypeSymbol,
			TypeData:     attributeSymbol.(SymbolTable).TypeData,
			TypeVariable: attributeSymbol.(SymbolTable).TypeVariable,
			Value:        expr,
			ListParams:   attributeSymbol.(SymbolTable).ListParams,
			Mutating:     attributeSymbol.(SymbolTable).Mutating,
			Line:         attributeSymbol.(SymbolTable).Line,
			Column:       attributeSymbol.(SymbolTable).Column,
			StructOf:     symbolSelf.(SymbolTable).Id,
		}
	} else {
		// evaluate if we want to only access or change the value
		if ctx.IS_() == nil {
			// return the value
			return attributeSymbol.(SymbolTable).Value.(values.PRIMITIVE)
		}
		// error
		// add the error to the errors
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "The function '" + functionName + "' is not mutating",
			Type:   "SintacticError",
		})
		return values.Nil{}
	}

	return nil
}
