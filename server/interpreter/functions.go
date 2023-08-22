package interpreter

import (
	"fmt"
	"log"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

func (v *Visitor) VisitFunctionWithoutParams(ctx *parser.FunctionWithoutParamsContext) interface{} {

	// get the id
	idFunction := ctx.ID_PRIMITIVE().GetText()
	// verify if the function is already declared
	if v.VerifyFunctionScope(idFunction) {
		log.Printf("Function %s already declared \n", idFunction)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Function '%s' already declared", idFunction),
			Type:   "Semantic",
		})
		return nil
	}
	// get the block
	block := ctx.Block()

	// get the return type if it exists
	var returnType string
	if ctx.ARROW_FUNCTION() != nil {
		returnType = ctx.Type_().GetText()
		// save the function in the symbol table
		v.getCurrentScope()[idFunction] = SymbolTable{
			Id:           idFunction,
			TypeSymbol:   values.Type_Function,
			TypeVariable: returnType,
			TypeData:     values.Type_Function,
			Value:        block,
			Line:         ctx.GetStart().GetLine(),
			Column:       ctx.GetStart().GetColumn(),
		}

	} else {
		returnType = "void"
		// save the function in the symbol table
		v.getCurrentScope()[idFunction] = SymbolTable{
			Id:           idFunction,
			TypeSymbol:   values.Type_Function,
			TypeVariable: returnType,
			TypeData:     values.Type_Function,
			Value:        block,
			Line:         ctx.GetStart().GetLine(),
			Column:       ctx.GetStart().GetColumn(),
		}
	}

	fmt.Println("FunctionWithoutParams: ", idFunction, returnType)

	return nil
}

// VisitFunctionWithParams
func (v *Visitor) VisitFunctionWithParams(ctx *parser.FunctionWithParamsContext) interface{} {
	// get the id
	idFunction := ctx.ID_PRIMITIVE().GetText()
	// verify if the function is already declared
	if v.VerifyFunctionScope(idFunction) {
		log.Printf("Function %s already declared \n", idFunction)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Function '%s' already declared", idFunction),
			Type:   "Semantic",
		})
		return nil
	}

	// get list of params
	params := v.Visit(ctx.ListFunctionParams()).(map[string][]SymbolTable)
	// fmt.Print("ListFunctionParams:\n", params, "\n")
	// get the block
	block := ctx.Block()

	// get the return type if it exists
	var returnType string
	if ctx.ARROW_FUNCTION() != nil {
		returnType = ctx.Type_().GetText()
		// save the function in the symbol table
		v.getCurrentScope()[idFunction] = SymbolTable{
			Id:           idFunction,
			TypeSymbol:   values.Type_Function,
			TypeVariable: returnType,
			TypeData:     values.Type_Function,
			Value:        block,
			ListParams:   params,
			Line:         ctx.GetStart().GetLine(),
			Column:       ctx.GetStart().GetColumn(),
		}

	} else {
		returnType = "void"
		// save the function in the symbol table
		v.getCurrentScope()[idFunction] = SymbolTable{
			Id:           idFunction,
			TypeSymbol:   values.Type_Function,
			TypeVariable: returnType,
			TypeData:     values.Type_Function,
			Value:        block,
			ListParams:   params,
			Line:         ctx.GetStart().GetLine(),
			Column:       ctx.GetStart().GetColumn(),
		}
	}

	return nil
}

// VisitListFunctionParamsEI
func (v *Visitor) VisitListFunctionParamsEI(ctx *parser.ListFunctionParamsEIContext) interface{} {

	// create a map to save the params
	params := make(map[string][]SymbolTable)

	// create two keys, external and internal
	params["external"] = []SymbolTable{}
	params["internal"] = []SymbolTable{}

	listIds := ctx.AllID_PRIMITIVE()
	listTypes := ctx.AllType_()

	// iterate over the list of ids
	for i, id := range listIds {
		// get the type
		typeParam := listTypes[i/2].GetText()

		// create a new symbol table
		symbol := SymbolTable{
			Id:           id.GetText(),
			TypeSymbol:   values.Type_Variable,
			TypeVariable: "let",
			TypeData:     typeParam,
			Value:        nil,
			Line:         ctx.GetStart().GetLine(),
			Column:       ctx.GetStart().GetColumn(),
		}
		// comes external then internal
		if i%2 == 0 {
			params["external"] = append(params["external"], symbol)
		} else {
			params["internal"] = append(params["internal"], symbol)
		}

	}

	// fmt.Println("ListFunctionParamsEI:\n ", params)

	return params
}

// VisitListFunctionParamsNEI
func (v *Visitor) VisitListFunctionParamsNEI(ctx *parser.ListFunctionParamsNEIContext) interface{} {

	// create a map to save the params
	params := make(map[string][]SymbolTable)

	// create two keys, external and internal
	params["internal"] = []SymbolTable{}

	listIds := ctx.AllID_PRIMITIVE()
	listTypes := ctx.AllType_()

	// iterate over the list of ids
	for i, id := range listIds {
		// get the type
		typeParam := listTypes[i/2].GetText()
		// create a new symbol table
		symbol := SymbolTable{
			Id:           id.GetText(),
			TypeSymbol:   values.Type_Variable,
			TypeVariable: "let",
			TypeData:     typeParam,
			Value:        nil,
			Line:         ctx.GetStart().GetLine(),
			Column:       ctx.GetStart().GetColumn(),
		}
		// comes external then internal
		params["internal"] = append(params["internal"], symbol)

	}

	fmt.Println("ListFunctionParamsNEI:\n", params)

	return params
}

// VisitListFunctionParamsBEI
func (v *Visitor) VisitListFunctionParamsBEI(ctx *parser.ListFunctionParamsBEIContext) interface{} {

	// create a map to save the params
	params := make(map[string][]SymbolTable)

	// create two keys, external and internal
	params["internal"] = []SymbolTable{}

	listIds := ctx.AllID_PRIMITIVE()
	listTypes := ctx.AllType_()

	// iterate over the list of ids
	for i, id := range listIds {
		// get the type
		typeParam := listTypes[i/2].GetText()

		// create a new symbol table
		symbol := SymbolTable{
			Id:           id.GetText(),
			TypeSymbol:   values.Type_Variable,
			TypeVariable: "let",
			TypeData:     typeParam,
			Value:        nil,
			Line:         ctx.GetStart().GetLine(),
			Column:       ctx.GetStart().GetColumn(),
		}
		// comes external then internal
		params["internal"] = append(params["internal"], symbol)

	}

	// fmt.Println("ListFunctionParamsBEI:\n ", params)

	return params
}
