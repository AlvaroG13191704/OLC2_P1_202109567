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
	listInitArguments := v.Visit(ctx.StructCallList()).(map[string]values.PRIMITIVE)
	// assert the type of the struct
	symbolStruct := structValue.(SymbolTable)

	// get the scope of the struct
	scopeStruct := symbolStruct.Value.(map[string]SymbolTable)
	fmt.Println("scopeStruct before changes", scopeStruct)

	// new Scope
	newScope := map[string]SymbolTable{}

	// iterate over the symbolStruct to verify
	// 1. If the variable is var
	for _, symbol := range scopeStruct {

		// TODO: CORRECT IMPLEMENTATION OF VECTORS
		// if symbol is a vector
		if symbol.TypeSymbol == values.Type_Vector {
			// print the argument
			// save the vector
			newScope[symbol.Id] = SymbolTable{
				Id:           symbol.Id,
				TypeSymbol:   values.Type_Vector,
				TypeVariable: symbol.TypeVariable,
				TypeData:     symbol.TypeData,
				Value:        symbol.Value,
				Line:         symbol.Line,
				Column:       symbol.Column,
			}
			continue
		}

		// if var?
		if symbol.TypeVariable == "var" {
			// if value is not nil and the argument doesn't comes, continue
			if symbol.Value.(values.PRIMITIVE).GetType() != values.NilType && listInitArguments[symbol.Id] == nil {
				continue
			}
			// if value is nil and the argument comes, assing the value
			if symbol.Value.(values.PRIMITIVE).GetType() == values.NilType && listInitArguments[symbol.Id] != nil {
				// evaluate if the type of the argument is the same as the type of the variable
				if symbol.TypeData == listInitArguments[symbol.Id].GetType() {
					newScope[symbol.Id] = SymbolTable{
						Id:           symbol.Id,
						TypeSymbol:   values.Type_Variable,
						TypeVariable: symbol.TypeVariable,
						TypeData:     symbol.TypeData,
						Value:        listInitArguments[symbol.Id],
						Line:         symbol.Line,
						Column:       symbol.Column,
					}

				} else {
					v.Errors = append(v.Errors, Error{Line: ctx.GetStart().GetLine(), Column: ctx.GetStart().GetColumn(), Msg: "The type of the argument " + symbol.Id + " is not the same as the type of the variable", Type: "Variable"})
					log.Printf("The type of the argument %s is not the same as the type of the variable", symbol.Id)
					return nil
				}
				continue
			}
			// if value is not nil and the argument comes, assing the value
			if symbol.Value.(values.PRIMITIVE).GetType() != values.NilType && listInitArguments[symbol.Id] != nil {
				// evaluate if the type of the argument is the same as the type of the variable
				if symbol.TypeData == listInitArguments[symbol.Id].GetType() {

					newScope[symbol.Id] = SymbolTable{
						Id:           symbol.Id,
						TypeSymbol:   values.Type_Variable,
						TypeVariable: symbol.TypeVariable,
						TypeData:     symbol.TypeData,
						Value:        listInitArguments[symbol.Id],
						Line:         symbol.Line,
						Column:       symbol.Column,
					}

				} else {
					v.Errors = append(v.Errors, Error{Line: ctx.GetStart().GetLine(), Column: ctx.GetStart().GetColumn(), Msg: "The type of the argument " + symbol.Id + " is not the same as the type of the variable", Type: "Variable"})
					log.Printf("The type of the argument %s is not the same as the type of the variable", symbol.Id)
					return nil
				}
				continue
			}
			// throw error if the value is nil and the argument doesn't comes
			if symbol.Value.(values.PRIMITIVE).GetType() == values.NilType && listInitArguments[symbol.Id] == nil {
				v.Errors = append(v.Errors, Error{Line: ctx.GetStart().GetLine(), Column: ctx.GetStart().GetColumn(), Msg: "The variable " + symbol.Id + " is not initialized or the name is not the same", Type: "Variable"})
				log.Printf("The variable %s is not initialized or the name is not the same", symbol.Id)
				return nil
			}

		} else if symbol.TypeVariable == "let" {
			// if value is not nil and the argument doesn't comes, continue
			if symbol.Value.(values.PRIMITIVE).GetType() != values.NilType && listInitArguments[symbol.Id] == nil {
				continue
			}
			// if value is nil and the argument comes, assing the value
			if symbol.Value.(values.PRIMITIVE).GetType() == values.NilType && listInitArguments[symbol.Id] != nil {
				// evaluate if the type of the argument is the same as the type of the variable
				if symbol.TypeData == listInitArguments[symbol.Id].GetType() {
					newScope[symbol.Id] = SymbolTable{
						Id:           symbol.Id,
						TypeSymbol:   values.Type_Variable,
						TypeVariable: symbol.TypeVariable,
						TypeData:     symbol.TypeData,
						Value:        listInitArguments[symbol.Id],
						Line:         symbol.Line,
						Column:       symbol.Column,
					}

				} else {
					v.Errors = append(v.Errors, Error{Line: ctx.GetStart().GetLine(), Column: ctx.GetStart().GetColumn(), Msg: "The type of the argument " + symbol.Id + " is not the same as the type of the variable", Type: "Variable"})
					log.Printf("The type of the argument %s is not the same as the type of the variable", symbol.Id)
					return nil
				}
				continue
			}
			// if value is not nil and the argument comes, throw error because the variable is let
			if symbol.Value.(values.PRIMITIVE).GetType() != values.NilType && listInitArguments[symbol.Id] != nil {
				v.Errors = append(v.Errors, Error{Line: ctx.GetStart().GetLine(), Column: ctx.GetStart().GetColumn(), Msg: "The variable " + symbol.Id + " is let", Type: "Variable"})
				log.Printf("The variable %s is inmutable, cannot be initializated", symbol.Id)
				return nil
			}
			// throw error if the value is nil and the argument doesn't comes
			if symbol.Value.(values.PRIMITIVE).GetType() == values.NilType && listInitArguments[symbol.Id] == nil {
				v.Errors = append(v.Errors, Error{Line: ctx.GetStart().GetLine(), Column: ctx.GetStart().GetColumn(), Msg: "The variable " + symbol.Id + " is not initialized or the name is not the same", Type: "Variable"})
				log.Printf("The variable %s is not initialized or the name is not the same", symbol.Id)
				return nil
			}
		} else if symbol.TypeSymbol == values.Type_Function {
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
		Id:           ctx.AllID_PRIMITIVE()[0].GetText(),
		TypeSymbol:   values.Type_Variable,
		TypeVariable: ctx.Type_declaration().GetText(),
		TypeData:     values.StructType,
		Value:        newScope,
		Line:         ctx.GetStart().GetLine(),
		Column:       ctx.GetStart().GetColumn(),
	}

	v.getCurrentScope()[ctx.AllID_PRIMITIVE()[0].GetText()] = newSymbolStruct

	v.TableSymbol = append(v.TableSymbol, newSymbolStruct)

	fmt.Println("----------------------------------------------------")
	fmt.Println("Global scope or symbol table ->", v.SymbolStack)
	fmt.Println("----------------------------------------------------")

	return nil
}

// VisitStructCallList
func (v *Visitor) VisitStructCallList(ctx *parser.StructCallListContext) interface{} {
	// create a map of values
	list := map[string]values.PRIMITIVE{}

	// TODO: IMPLEMENT VECTORS AND STRUCTS AS ARGUMENTS

	// get the list of values
	for i, value := range ctx.AllID_PRIMITIVE() {

		assertion := v.Visit(ctx.AllExpr()[i])

		// evaluate if a struct
		if reflect.TypeOf(assertion).Kind() == reflect.Struct {
			fmt.Println("assertion struct -> ", assertion)

		} else {
			fmt.Println("assertion primitive -> ", assertion)
			list[value.GetText()] = assertion.(values.PRIMITIVE)
		}
	}

	return list
}
