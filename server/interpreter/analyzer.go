package interpreter

import (
	"log"
	"server/parserInterpreter/parser"

	"github.com/antlr4-go/antlr/v4"
)

func NewVisitor() *Visitor {
	return &Visitor{
		symbolStack: []map[string]SymbolTable{},
		Outputs:     []string{},
		Errors:      []Error{},
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
	if ctx.Declaration() != nil {
		return v.Visit(ctx.Declaration())
	}
	if ctx.Assignment() != nil {
		return v.Visit(ctx.Assignment())
	}
	if ctx.EmbbededFunc() != nil {
		return v.Visit(ctx.EmbbededFunc())
	}
	if ctx.Ifstmt() != nil {
		return v.Visit(ctx.Ifstmt())
	}
	if ctx.SwitchStmt() != nil {
		return v.Visit(ctx.SwitchStmt())
	}
	if ctx.WhileStmt() != nil {
		return v.Visit(ctx.WhileStmt())
	}
	if ctx.ForStmt() != nil {
		return v.Visit(ctx.ForStmt())
	}

	return nil
}

// visit embbededFunc
func (v *Visitor) VisitEmbbededFunc(ctx *parser.EmbbededFuncContext) interface{} {
	// verify if the function is a print
	if ctx.Printstmt() != nil {
		return v.Visit(ctx.Printstmt())
	}

	return nil

}
