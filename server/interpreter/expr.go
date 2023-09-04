package interpreter

import (
	"fmt"
	"log"
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
	listInitArguments := v.Visit(ctx.StructCallList()).(map[string]values.PRIMITIVE)
	// assert the type of the struct
	symbolStruct := structValue.(SymbolTable)

	// get the scope of the struct
	scopeStruct := symbolStruct.Value.(map[string]SymbolTable)
	fmt.Println("scopeStruct before changes in struct as argument", scopeStruct)

	// new Scope
	newScope := map[string]SymbolTable{}
	// iterate over the symbolStruct to verify
	// 1. If the variable is var
	for _, symbol := range scopeStruct {

		// TODO: CORRECT IMPLEMENTATION OF VECTORS

		// if var?
		if symbol.TypeVariable == "var" {
			// if value is not nil and the argument doesn't comes, continue
			if symbol.Value.(values.PRIMITIVE).GetType() != values.NilType && listInitArguments[symbol.Id] == nil {
				// add the existing values
				newScope[symbol.Id] = SymbolTable{
					Id:           symbol.Id,
					TypeSymbol:   values.Type_Variable,
					TypeVariable: symbol.TypeVariable,
					TypeData:     symbol.TypeData,
					Value:        symbol.Value,
					Line:         symbol.Line,
					Column:       symbol.Column,
					ListParams:   symbol.ListParams,
					Mutating:     symbol.Mutating,
					StructOf:     symbol.StructOf,
				}
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
				// save the value
				newScope[symbol.Id] = SymbolTable{
					Id:           symbol.Id,
					TypeSymbol:   values.Type_Variable,
					TypeVariable: symbol.TypeVariable,
					TypeData:     symbol.TypeData,
					Value:        symbol.Value,
					Line:         symbol.Line,
					Column:       symbol.Column,
					ListParams:   symbol.ListParams,
					Mutating:     symbol.Mutating,
				}
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

	// newSymbolStruct := SymbolTable{
	// 	Id:           ctx.ID_PRIMITIVE().GetText(),
	// 	TypeSymbol:   values.Type_Variable,
	// 	TypeVariable: values.StructType,
	// 	TypeData:     values.StructType,
	// 	StructOf:     ctx.ID_PRIMITIVE().GetText(),
	// 	Value:        newScope,
	// 	Line:         ctx.GetStart().GetLine(),
	// 	Column:       ctx.GetStart().GetColumn(),
	// }

	// v.SelfStructs[newSymbolStruct.Id] = SelfStruct{VarId: newSymbolStruct.Id, StructOf: newSymbolStruct.StructOf}
	// v.TableSymbol = append(v.TableSymbol, newSymbolStruct)

	return newScope
}

// VisitCallBackExpr -> Visit a parse tree produced by SFGrammarParser#CallBackExpr.
func (v *Visitor) VisitCallBackExpr(ctx *parser.CallBackExprContext) interface{} {

	return v.Visit(ctx.CallBack())
}
