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
