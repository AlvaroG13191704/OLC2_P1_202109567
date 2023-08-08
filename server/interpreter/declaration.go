package interpreter

import (
	"fmt"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

func (v *Visitor) VisitDeclaration(ctx *parser.DeclarationContext) interface{} {
	// get the declaration type -> var or let
	declarationType1 := ctx.DECLARATION_1().GetText()
	// declarationType2 := ctx.DECLARATION_2().GetText()
	fmt.Println(declarationType1)
	// get the variable name
	varName := ctx.ID_PRIMITIVE().GetText()
	fmt.Println(varName)
	// get the value of the variable
	varValue := v.Visit(ctx.Expr()).(values.PRIMITIVE)
	fmt.Println(varValue)
	// get the type of the variable
	varType_fromSyntax := ctx.Type_().GetText()
	fmt.Println(varType_fromSyntax)
	// get the type of the variable
	varType_fromContext := varValue.GetType()
	fmt.Println(varType_fromContext)

	v.getCurrentScope()[varName] = SymbolTable{
		Id:         varName,
		TypeSymbol: "Variable",
		TypeData:   varType_fromContext,
		Value:      varValue.GetValue(),
		Line:       ctx.GetStart().GetLine(),
		Column:     ctx.GetStart().GetColumn(),
	}

	// print the symbol table
	fmt.Println("Current scope or symbol table ->", v.symbolStack)
	return nil
}
