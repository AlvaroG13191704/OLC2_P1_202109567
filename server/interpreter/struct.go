package interpreter

import (
	"fmt"
	"log"
	"reflect"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

// VisitStructStmt
func (v *Visitor) VisitStructStmt(ctx *parser.StructStmtContext) interface{} {
	fmt.Println("StructStmt")
	// get the id
	id := ctx.ID_PRIMITIVE().GetText()
	if v.VerifyVariableScope(id) {
		v.Errors = append(v.Errors, Error{Line: ctx.GetStart().GetLine(), Column: ctx.GetStart().GetColumn(), Msg: "The variable " + id + " already exists", Type: "Variable"})
		log.Printf("The variable %s already exists", id)
		return nil
	}

	// get the struct block
	block := v.Visit(ctx.StructBlock()).(map[string]SymbolTable)

	// add the variable to the scope
	symbol := SymbolTable{
		Id:           id,
		TypeSymbol:   values.Type_Struct,
		TypeVariable: values.StructType,
		TypeData:     values.StructType,
		Value:        block,
		Line:         ctx.GetStart().GetLine(),
		Column:       ctx.GetStart().GetColumn(),
	}

	// add the variable to the scope
	v.getCurrentScope()[id] = symbol

	return nil
}

// VisitStructBlock
func (v *Visitor) VisitStructBlock(ctx *parser.StructBlockContext) interface{} {

	// initialize the scope
	scopeStruct := map[string]SymbolTable{}

	// get the struct block
	for _, stmt := range ctx.AllStructStmts() {
		stmtMap := v.Visit(stmt).(SymbolTable)

		// verify if the variable is already defined in the scopeStruct
		if _, ok := scopeStruct[stmtMap.Id]; ok {
			v.Errors = append(v.Errors, Error{Line: stmtMap.Line, Column: stmtMap.Column, Msg: "The variable " + stmtMap.Id + " already exists", Type: "Variable"})
			log.Printf("The variable %s already exists", stmtMap.Id)
			return nil
		}

		// add the variable to the scope
		scopeStruct[stmtMap.Id] = stmtMap
		// fmt.Println("stmtMap", stmtMap)
	}

	return scopeStruct
}

// VisitStructStmts
func (v *Visitor) VisitStructStmts(ctx *parser.StructStmtsContext) interface{} {

	if ctx.DeclarationStructs() != nil {
		return v.Visit(ctx.DeclarationStructs())
	}

	if ctx.FunctionStructs() != nil {
		return v.Visit(ctx.FunctionStructs())
	}
	return nil
}

// VisitStructCreation
func (v *Visitor) VisitStructCreation(ctx *parser.StructCreationContext) interface{} {
	fmt.Println("StructCreation")
	// verify if the variable is already defined in the scope
	if v.VerifyVariableScope(ctx.AllID_PRIMITIVE()[0].GetText()) {
		v.Errors = append(v.Errors, Error{Line: ctx.GetStart().GetLine(), Column: ctx.GetStart().GetColumn(), Msg: "The variable " + ctx.AllID_PRIMITIVE()[0].GetText() + " already exists", Type: "Variable"})
		log.Printf("The variable %s already exists", ctx.AllID_PRIMITIVE()[0].GetText())
		return nil
	}
	// get the symbol table of the struct
	structValue, ok := v.VerifyScope(ctx.AllID_PRIMITIVE()[1].GetText())

	if !ok {
		v.Errors = append(v.Errors, Error{Line: ctx.GetStart().GetLine(), Column: ctx.GetStart().GetColumn(), Msg: "The struct " + ctx.AllID_PRIMITIVE()[1].GetText() + " does not exists", Type: "Variable"})
		log.Printf("The struct %s does not exists", ctx.AllID_PRIMITIVE()[1].GetText())
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
		// assert the type of the struct
		symbolStruct := structValue.(SymbolTable)
		// get the scope of the struct
		scopeStruct := symbolStruct.Value.(map[string]SymbolTable)

		// iterate over the symbolStruct to verify
		// 1. If the variable is var
		if scopeStruct[key].TypeVariable == "var" {
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
			if symbol.TypeSymbol == values.Type_Variable {
				if key != symbol.Id {
					fmt.Println("symbol ----", symbol)
					fmt.Println("symbol.Value ----", symbol.Value)
					if symbol.Value == nil {
						// save the value
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
					} else if symbol.Value != nil {
						// verify if not already saved in the newScope
						if _, ok := newScope[symbol.Id]; !ok {
							// save the value
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
				}
			}
		}

		fmt.Println("scopeStruct after changes", newScope)

		newSymbolStruct := SymbolTable{
			Id:           ctx.AllID_PRIMITIVE()[0].GetText(),
			TypeSymbol:   values.Type_Variable,
			TypeVariable: ctx.Type_declaration().GetText(),
			TypeData:     values.StructType,
			StructOf:     ctx.AllID_PRIMITIVE()[1].GetText(),
			Value:        newScope,
			Line:         ctx.GetStart().GetLine(),
			Column:       ctx.GetStart().GetColumn(),
		}

		v.getCurrentScope()[ctx.AllID_PRIMITIVE()[0].GetText()] = newSymbolStruct

		// append his self where the was created
		v.SelfStructs[newSymbolStruct.Id] = SelfStruct{VarId: newSymbolStruct.Id, StructOf: newSymbolStruct.StructOf}

		v.TableSymbol = append(v.TableSymbol, newSymbolStruct)
	}

	return nil
}

// VisitStructCallList
func (v *Visitor) VisitStructCallList(ctx *parser.StructCallListContext) interface{} {
	// create a map of values
	// list := map[string]values.PRIMITIVE{}
	list := map[string]interface{}{}

	// TODO: IMPLEMENT VECTORS AND STRUCTS AS ARGUMENTS

	// get the list of values
	for i, value := range ctx.AllID_PRIMITIVE() {

		// get the value
		assertion := v.Visit(ctx.AllExpr()[i])
		// evaluate if a struct
		if reflect.TypeOf(assertion).Kind() == reflect.Map {
			fmt.Println("assertion map -> ", assertion)
			list[value.GetText()] = assertion.(map[string]SymbolTable)
		} else {
			fmt.Println("assertion primitive -> ", assertion)
			list[value.GetText()] = assertion.(values.PRIMITIVE)
		}

	}

	fmt.Println("list of arguments -> ", list)
	return list
}
