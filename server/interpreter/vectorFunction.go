package interpreter

import (
	"fmt"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

// VisitListFunctionParamsEIVector
func (v *Visitor) VisitListFunctionParamsEIVector(ctx *parser.ListFunctionParamsEIVectorContext) interface{} {
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
			TypeSymbol:   values.Type_Vector,
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

	fmt.Println("ListFunctionParamsEIVector:\n ", params)

	return params

}

// VisitListFunctionParamsNEIVector
func (v *Visitor) VisitListFunctionParamsNEIVector(ctx *parser.ListFunctionParamsNEIVectorContext) interface{} {

	// create a map to save the params
	params := make(map[string][]SymbolTable)

	// create two keys, external and internal
	params["internal"] = []SymbolTable{}

	listIds := ctx.AllID_PRIMITIVE()
	listTypes := ctx.AllType_()

	// iterate over the list of ids
	for i, id := range listIds {
		// create the
		// get the type
		typeParam := listTypes[i/2].GetText()
		// create a new symbol table
		symbol := SymbolTable{
			Id:           id.GetText(),
			TypeSymbol:   values.Type_Vector,
			TypeVariable: "let",
			TypeData:     typeParam,
			Value:        nil,
			Line:         ctx.GetStart().GetLine(),
			Column:       ctx.GetStart().GetColumn(),
		}
		// comes external then internal
		params["internal"] = append(params["internal"], symbol)

	}

	fmt.Println("ListFunctionParamsNEIVector:\n", params)

	return params
}

// VisitListFunctionParamsBEIVector
func (v *Visitor) VisitListFunctionParamsBEIVector(ctx *parser.ListFunctionParamsBEIVectorContext) interface{} {
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
			TypeSymbol:   values.Type_Vector,
			TypeVariable: "let",
			TypeData:     typeParam,
			Value:        nil,
			Line:         ctx.GetStart().GetLine(),
			Column:       ctx.GetStart().GetColumn(),
		}
		// comes external then internal
		params["internal"] = append(params["internal"], symbol)

	}

	fmt.Println("ListFunctionParamsBEIVector:\n ", params)

	return params
}
