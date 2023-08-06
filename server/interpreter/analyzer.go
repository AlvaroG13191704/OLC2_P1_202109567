package interpreter

import (
	"fmt"
	"log"
	"server/parserInterpreter/parser"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

func NewVisitor() *Visitor {
	return &Visitor{
		Memory: make(map[string]interface{}),
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
	out := ""
	for i := 0; ctx.Stmts(i) != nil; i++ {
		out += strconv.FormatInt(v.Visit(ctx.Stmts(i)).(int64), 10) + ""
	}
	return out
}

// visit stmts
func (v *Visitor) VisitStmts(ctx *parser.StmtsContext) interface{} {
	// verify if the stmt which stmt is
	if ctx.Printstmt() != nil {
		return v.Visit(ctx.Printstmt())
	}
	return nil
}

// visit printstmt
func (v *Visitor) VisitPrintstmt(ctx *parser.PrintstmtContext) interface{} {
	value := v.Visit(ctx.Expr()) // visit the expression
	//TODO: evaluate the type of the value
	fmt.Println(value)
	return value
}
