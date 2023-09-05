package interpreter

import (
	"fmt"
	"log"
	"reflect"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

// visit idexpr
func (v *Visitor) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	id := ctx.GetText() // get the id
	// fmt.Println("Id -> ", id)

	// verify if the id is in the scope or others
	variable, ok := v.VerifyScope(id)
	if ok {
		value := variable.(SymbolTable)
		// return the value
		return value.Value

	} else {
		fmt.Printf("Scope -> %v\n", v.SymbolStack)
		// add the error to the errors
		log.Printf("Error: Variable '%s' not declared", id)
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Variable '" + id + "' not declared",
			Type:   "VariableError",
		})
	}
	return &values.Nil{
		Value: nil,
	}
}

// visit parenexpr
func (v *Visitor) VisitParenExpr(ctx *parser.ParenExprContext) interface{} {
	return v.Visit(ctx.Expr()) // visit the expression
}

// visit negativeexpr
func (v *Visitor) VisitNegExpr(ctx *parser.NegExprContext) interface{} {
	// get the value
	value := v.Visit(ctx.Expr()).(values.PRIMITIVE)
	// verify if the value is an integer or a float
	if value.GetType() == values.IntType {
		return &values.Integer{Value: -value.GetValue().(int64)} // return the negative of the value
	} else if value.GetType() == values.FloatType {
		return &values.Float{Value: -value.GetValue().(float64)} // return the negative of the value
	}
	return &values.Nil{
		Value: nil,
	}
}

// visit notexpr
func (v *Visitor) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	// get the value
	value := v.Visit(ctx.Expr()).(values.PRIMITIVE)
	// verify if the value is a boolean
	if value.GetType() == values.BooleanType {
		return &values.Boolean{Value: !value.GetValue().(bool)} // return the negative of the value
	}
	return &values.Nil{
		Value: nil,
	}

}

// VisitStructAsArgument
func (v *Visitor) VisitStructAsArgument(ctx *parser.StructAsArgumentContext) interface{} {
	fmt.Println("Struct as argument")
	// get the symbol table of the struct
	structValue, ok := v.VerifyScope(ctx.ID_PRIMITIVE().GetText())

	if !ok {
		v.Errors = append(v.Errors, Error{Line: ctx.GetStart().GetLine(), Column: ctx.GetStart().GetColumn(), Msg: "The struct " + ctx.ID_PRIMITIVE().GetText() + " does not exists", Type: "Variable"})
		log.Printf("The struct %s does not exists", ctx.ID_PRIMITIVE().GetText())
		return nil
	}
	// get list of init arguments
	assertionArguments := v.Visit(ctx.StructCallList())
	// new Scope
	newScope := map[string]SymbolTable{}
	for key, value := range assertionArguments.(map[string]interface{}) {

		// get the value
		assertlistInitArguments := value

		// assert the type of the struct
		if reflect.TypeOf(assertlistInitArguments).Kind() == reflect.Map {
			listInitArguments := assertlistInitArguments.(map[string]SymbolTable)

			fmt.Println("listInitArguments map (another struct)", listInitArguments)

			newScope[key] = SymbolTable{
				Id:           key,
				TypeSymbol:   values.Type_Variable,
				TypeVariable: "var",
				TypeData:     values.StructType,
				Value:        listInitArguments,
				Line:         ctx.GetStart().GetLine(),
				Column:       ctx.GetStart().GetColumn(),
			}

			continue
		}
		listInitArguments := assertlistInitArguments.(values.PRIMITIVE)
		fmt.Println("listInitArguments values primitve", listInitArguments)
		// assert the type of the struct
		symbolStruct := structValue.(SymbolTable)
		// get the scope of the struct
		scopeStruct := symbolStruct.Value.(map[string]SymbolTable)

		// iterate over the symbolStruct to verify
		// 1. If the variable is var
		if scopeStruct[key].TypeVariable == "var" {
			// if value is not nil and the argument doesn't comes, continue
			if scopeStruct[key].Value.(values.PRIMITIVE).GetType() != values.NilType && listInitArguments == nil {
				// add the existing values
				newScope[key] = SymbolTable{
					Id:           scopeStruct[key].Id,
					TypeSymbol:   values.Type_Variable,
					TypeVariable: scopeStruct[key].TypeVariable,
					TypeData:     scopeStruct[key].TypeData,
					Value:        scopeStruct[key].Value,
					Line:         scopeStruct[key].Line,
					Column:       scopeStruct[key].Column,
					ListParams:   scopeStruct[key].ListParams,
					Mutating:     scopeStruct[key].Mutating,
					StructOf:     scopeStruct[key].StructOf,
				}

			}
			// if value is nil and the argument comes, assing the value
			if scopeStruct[key].Value.(values.PRIMITIVE).GetType() == values.NilType && listInitArguments != nil {
				// evaluate if the type of the argument is the same as the type of the variable
				if scopeStruct[key].TypeData == listInitArguments.GetType() {
					newScope[key] = SymbolTable{
						Id:           scopeStruct[key].Id,
						TypeSymbol:   values.Type_Variable,
						TypeVariable: scopeStruct[key].TypeVariable,
						TypeData:     scopeStruct[key].TypeData,
						Value:        listInitArguments,
						Line:         scopeStruct[key].Line,
						Column:       scopeStruct[key].Column,
					}
				} else {
					v.Errors = append(v.Errors, Error{Line: ctx.GetStart().GetLine(), Column: ctx.GetStart().GetColumn(), Msg: "The type of the argument " + scopeStruct[key].Id + " is not the same as the type of the variable", Type: "Variable"})
					log.Printf("The type of the argument %s is not the same as the type of the variable", scopeStruct[key].Id)
					return nil
				}

			}
			// if value is not nil and the argument comes, assing the value
			if scopeStruct[key].Value.(values.PRIMITIVE).GetType() != values.NilType && listInitArguments != nil {
				// evaluate if the type of the argument is the same as the type of the variable
				if scopeStruct[key].TypeData == listInitArguments.GetType() {
					newScope[key] = SymbolTable{
						Id:           scopeStruct[key].Id,
						TypeSymbol:   values.Type_Variable,
						TypeVariable: scopeStruct[key].TypeVariable,
						TypeData:     scopeStruct[key].TypeData,
						Value:        listInitArguments,
						Line:         scopeStruct[key].Line,
						Column:       scopeStruct[key].Column,
					}
				} else {
					v.Errors = append(v.Errors, Error{Line: ctx.GetStart().GetLine(), Column: ctx.GetStart().GetColumn(), Msg: "The type of the argument " + scopeStruct[key].Id + " is not the same as the type of the variable", Type: "Variable"})
					log.Printf("The type of the argument %s is not the same as the type of the variable", scopeStruct[key].Id)
					return nil
				}
			}
			// throw error if the value is nil and the argument doesn't comes
			if scopeStruct[key].Value.(values.PRIMITIVE).GetType() == values.NilType && listInitArguments == nil {
				v.Errors = append(v.Errors, Error{Line: ctx.GetStart().GetLine(), Column: ctx.GetStart().GetColumn(), Msg: "The variable " + scopeStruct[key].Id + " is not initialized or the name is not the same", Type: "Variable"})
				log.Printf("The variable %s is not initialized or the name is not the same", scopeStruct[key].Id)
				return nil
			}
		} else if scopeStruct[key].TypeVariable == "let" {
			// if value is not nil and the argument doesn't comes, continue
			if scopeStruct[key].Value.(values.PRIMITIVE).GetType() != values.NilType && listInitArguments == nil {
				// save the value
				newScope[key] = SymbolTable{
					Id:           scopeStruct[key].Id,
					TypeSymbol:   values.Type_Variable,
					TypeVariable: scopeStruct[key].TypeVariable,
					TypeData:     scopeStruct[key].TypeData,
					Value:        scopeStruct[key].Value,
					Line:         scopeStruct[key].Line,
					Column:       scopeStruct[key].Column,
					ListParams:   scopeStruct[key].ListParams,
					Mutating:     scopeStruct[key].Mutating,
				}
			}
			// if value is nil and the argument comes, assing the value
			if scopeStruct[key].Value.(values.PRIMITIVE).GetType() == values.NilType && listInitArguments != nil {
				// evaluate if the type of the argument is the same as the type of the variable
				if scopeStruct[key].TypeData == listInitArguments.GetType() {
					newScope[key] = SymbolTable{
						Id:           scopeStruct[key].Id,
						TypeSymbol:   values.Type_Variable,
						TypeVariable: scopeStruct[key].TypeVariable,
						TypeData:     scopeStruct[key].TypeData,
						Value:        listInitArguments,
						Line:         scopeStruct[key].Line,
						Column:       scopeStruct[key].Column,
					}
				} else {
					v.Errors = append(v.Errors, Error{Line: ctx.GetStart().GetLine(), Column: ctx.GetStart().GetColumn(), Msg: "The type of the argument " + scopeStruct[key].Id + " is not the same as the type of the variable", Type: "Variable"})
					log.Printf("The type of the argument %s is not the same as the type of the variable", scopeStruct[key].Id)
					return nil
				}
			}
			// if value is not nil and the argument comes, throw error because the variable is let
			if scopeStruct[key].Value.(values.PRIMITIVE).GetType() != values.NilType && listInitArguments != nil {
				v.Errors = append(v.Errors, Error{Line: ctx.GetStart().GetLine(), Column: ctx.GetStart().GetColumn(), Msg: "The variable " + scopeStruct[key].Id + " is let", Type: "Variable"})
				log.Printf("The variable %s is inmutable, cannot be initializated", scopeStruct[key].Id)
				return nil
			}
			// throw error if the value is nil and the argument doesn't comes
			if scopeStruct[key].Value.(values.PRIMITIVE).GetType() == values.NilType && listInitArguments == nil {
				v.Errors = append(v.Errors, Error{Line: ctx.GetStart().GetLine(), Column: ctx.GetStart().GetColumn(), Msg: "The variable " + scopeStruct[key].Id + " is not initialized or the name is not the same", Type: "Variable"})
				log.Printf("The variable %s is not initialized or the name is not the same", scopeStruct[key].Id)
				return nil
			}
		}

		// copy functions, iterate over the scopeStruct to verify if the variable is a function
		for _, symbol := range scopeStruct {
			if symbol.TypeSymbol == values.Type_Function {
				// copy
				newScope[symbol.Id] = SymbolTable{
					Id:           symbol.Id,
					TypeSymbol:   symbol.TypeSymbol,
					TypeVariable: symbol.TypeVariable,
					TypeData:     symbol.TypeData,
					Value:        symbol.Value,
					ListParams:   symbol.ListParams,
					Mutating:     symbol.Mutating,
					Line:         symbol.Line,
					Column:       symbol.Column,
				}
			}
		}

		fmt.Println("scopeStruct after changes", newScope)

		newSymbolStruct := SymbolTable{
			Id:           ctx.ID_PRIMITIVE().GetText(),
			TypeSymbol:   values.Type_Variable,
			TypeVariable: ctx.ID_PRIMITIVE().GetText(),
			TypeData:     values.StructType,
			StructOf:     ctx.ID_PRIMITIVE().GetText(),
			Value:        newScope,
			Line:         ctx.GetStart().GetLine(),
			Column:       ctx.GetStart().GetColumn(),
		}

		v.getCurrentScope()[ctx.ID_PRIMITIVE().GetText()] = newSymbolStruct

		// append his self where the was created
		v.SelfStructs[newSymbolStruct.Id] = SelfStruct{VarId: newSymbolStruct.Id, StructOf: newSymbolStruct.StructOf}

		v.TableSymbol = append(v.TableSymbol, newSymbolStruct)
	}

	return newScope
}

// VisitCallBackExpr -> Visit a parse tree produced by SFGrammarParser#CallBackExpr.
func (v *Visitor) VisitCallBackExpr(ctx *parser.CallBackExprContext) interface{} {

	return v.Visit(ctx.CallBack())
}
