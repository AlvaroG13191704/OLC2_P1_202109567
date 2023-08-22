package interpreter

import (
	"fmt"
	"log"
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
	v.getCurrentScope()[id] = SymbolTable{
		Id:           id,
		TypeSymbol:   values.Type_Struct,
		TypeVariable: values.StructType,
		TypeData:     values.StructType,
		Value:        block,
		Line:         ctx.GetStart().GetLine(),
		Column:       ctx.GetStart().GetColumn(),
	}

	return nil
}

// VisitStructBlock
func (v *Visitor) VisitStructBlock(ctx *parser.StructBlockContext) interface{} {

	// initialize the scope
	scopeStruct := map[string]SymbolTable{}

	// get the struct block
	for _, stmt := range ctx.AllStructStmts() {
		stmtMap := v.Visit(stmt).(SymbolTable)
		// add the variable to the scope
		scopeStruct[stmtMap.Id] = stmtMap
		fmt.Println("stmtMap", stmtMap)
	}

	return scopeStruct
}

// VisitStructStmts
func (v *Visitor) VisitStructStmts(ctx *parser.StructStmtsContext) interface{} {

	if ctx.DeclarationStructs() != nil {
		return v.Visit(ctx.DeclarationStructs())
	}

	return nil
}
