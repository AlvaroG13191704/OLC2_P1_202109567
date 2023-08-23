package interpreter

import (
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

// VisitStructFunctionWithoutParams
func (v *Visitor) VisitStructFunctionWithoutParams(ctx *parser.StructFunctionWithoutParamsContext) interface{} {
	var symbolTable SymbolTable
	var mutating = false
	// get the id
	idFunction := ctx.ID_PRIMITIVE().GetText()

	// get the block
	block := ctx.Block()

	// evaluate if there is mutating
	if ctx.MUTATING() != nil {
		mutating = true
	}

	// get the return type if it exists
	var returnType string
	if ctx.ARROW_FUNCTION() != nil {
		returnType = ctx.Type_().GetText()
		// save the function in the symbol table
		symbolTable = SymbolTable{
			Id:           idFunction,
			TypeSymbol:   values.Type_Function,
			TypeVariable: returnType,
			TypeData:     values.Type_Function,
			Value:        block,
			Mutating:     mutating,
			Line:         ctx.GetStart().GetLine(),
			Column:       ctx.GetStart().GetColumn(),
		}

	} else {
		returnType = "void"
		// save the function in the symbol table
		symbolTable = SymbolTable{
			Id:           idFunction,
			TypeSymbol:   values.Type_Function,
			TypeVariable: returnType,
			TypeData:     values.Type_Function,
			Value:        block,
			Mutating:     mutating,
			Line:         ctx.GetStart().GetLine(),
			Column:       ctx.GetStart().GetColumn(),
		}
	}

	// fmt.Println("FunctionWithoutParams: ", idFunction, returnType)

	return symbolTable
}

// VisitStructFunctionWithParams
func (v *Visitor) VisitStructFunctionWithParams(ctx *parser.StructFunctionWithParamsContext) interface{} {
	var symbolTable SymbolTable
	var mutating = false
	// get the id
	idFunction := ctx.ID_PRIMITIVE().GetText()
	// get list of params
	params := v.Visit(ctx.ListFunctionParams()).(map[string][]SymbolTable)

	// evaluate if there is mutating
	if ctx.MUTATING() != nil {
		mutating = true
	}
	// get the block
	block := ctx.Block()

	// get the return type if it exists
	var returnType string
	if ctx.ARROW_FUNCTION() != nil {
		returnType = ctx.Type_().GetText()
		// save the function in the symbol table
		symbolTable = SymbolTable{
			Id:           idFunction,
			TypeSymbol:   values.Type_Function,
			TypeVariable: returnType,
			TypeData:     values.Type_Function,
			Value:        block,
			ListParams:   params,
			Mutating:     mutating,
			Line:         ctx.GetStart().GetLine(),
			Column:       ctx.GetStart().GetColumn(),
		}

	} else {
		returnType = "void"
		// save the function in the symbol table
		symbolTable = SymbolTable{
			Id:           idFunction,
			TypeSymbol:   values.Type_Function,
			TypeVariable: returnType,
			TypeData:     values.Type_Function,
			Value:        block,
			ListParams:   params,
			Mutating:     mutating,
			Line:         ctx.GetStart().GetLine(),
			Column:       ctx.GetStart().GetColumn(),
		}
	}

	return symbolTable
}
