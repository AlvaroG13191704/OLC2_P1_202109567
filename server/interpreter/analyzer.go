package interpreter

import (
	"fmt"
	"log"
	"server/parserInterpreter/parser"

	"github.com/antlr4-go/antlr/v4"
)

func NewVisitor() *Visitor {
	return &Visitor{
		memory:      make(map[string]SymbolTable),
		symbolStack: []map[string]SymbolTable{},
		Outputs:     []string{},
	}
}

// The visit method
func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		node := tree.Accept(v)
		return node
	}
}

// visit start
func (v *Visitor) VisitStart(ctx *parser.StartContext) interface{} {
	return v.Visit(ctx.Block())
}

// visit block
func (v *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	// push the scope
	v.pushScope()
	for _, stmt := range ctx.AllStmts() {
		// evalaute if the stmt is not a interface nil
		v.Visit(stmt)
	}
	// pop the scope
	v.popScope()
	return nil
}

// visit stmts
func (v *Visitor) VisitStmts(ctx *parser.StmtsContext) interface{} {
	// verify if the stmt which stmt is
	if ctx.Printstmt() != nil {
		return v.Visit(ctx.Printstmt())
	}
	if ctx.Declaration() != nil {
		return v.Visit(ctx.Declaration())
	}
	return nil
}

// visit printstmt
func (v *Visitor) VisitPrintstmt(ctx *parser.PrintstmtContext) interface{} {
	value := v.Visit(ctx.Expr()) // visit the expression

	//TODO: evaluate the type of the value

	fmt.Println("print value ->", value)

	// add to the outputs
	v.Outputs = append(v.Outputs, fmt.Sprint(value))
	return value
}

// visit varReference
func (v *Visitor) VerifyScope(varName string) interface{} {

	for i := len(v.symbolStack) - 1; i >= 0; i-- {
		scope := v.symbolStack[i]
		if val, ok := scope[varName]; ok {
			return val
		}
	}
	log.Fatalf("Error: Variable '%s' not declared", varName)
	return nil
}
