package main

import (
	"fmt"
	"log"
	"server/parserInterpreter/parser"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gofiber/fiber/v2"
)

// Value is a wrapper for any type
// if it's get more complex add methods to it
type Value struct {
	value interface{} // generic value
}

type Visitor struct {
	memory map[string]Value // memory for variables
}

// Visit receives a tree and returns a value
func (v *Visitor) Visit(tree antlr.ParseTree) Value {
	switch val := tree.(type) { // the val is the context of each statement
	case *parser.ProgContext:
		return v.VisitProg(val)
	case *parser.BlockContext:
		return v.VisitBlock(val)
	case *parser.StmtContext:
		return v.VisitStmt(val)
	case *parser.AssignstmtContext:
		return v.VisitAssignstmt(val)
	case *parser.IfstmtContext:
		return v.VisitIfstmt(val)
	case *parser.PrintlnstmtContext:
		return v.VisitPrintlnstmt(val)
	case *parser.WhilestmtContext:
		return v.VisitWhilestmt(val)
	case *parser.ForstmtContext:
		return v.VisitForstmt(val)
	case *parser.PrintstmtContext:
		return v.VisitPrintstmt(val)
	case *parser.OpExprContext:
		return v.VisitOpExpr(val)
	case *parser.IntExprContext:
		return v.VisitIntExpr(val)
	case *parser.IdExprContext:
		return v.VisitIdExpr(val)
	case *parser.StrExprContext:
		return v.VisitStrExpr(val)
	case *parser.BoolExprContext:
		return v.VisitBoolExpr(val)
	default:
		panic("Unknown context " + val.GetText())
	}

}

// override the visitProg function
func (v *Visitor) VisitProg(ctx *parser.ProgContext) Value {
	return v.Visit(ctx.Block()) // When we visit a prog we should visit the block
}

// override the visitBlock function
func (v *Visitor) VisitBlock(ctx *parser.BlockContext) Value { // in the grammar we use * so it will work as a slice of stmts
	for i := 0; ctx.Stmt(i) != nil; i++ { // iterate over all the statements
		v.Visit(ctx.Stmt(i)) // visit each statement
	}
	return Value{value: true}
}

// override the visitStmt function
func (v *Visitor) VisitStmt(ctx *parser.StmtContext) Value {
	// verify which type of statement it is
	if ctx.Assignstmt() != nil {
		return v.Visit(ctx.Assignstmt())
	}
	if ctx.Printlnstmt() != nil {
		return v.Visit(ctx.Printlnstmt())
	}
	if ctx.Ifstmt() != nil {
		return v.Visit(ctx.Ifstmt())
	}
	if ctx.Whilestmt() != nil {
		return v.Visit(ctx.Whilestmt())
	}
	if ctx.Forstmt() != nil {
		return v.Visit(ctx.Forstmt())
	}
	if ctx.Printstmt() != nil {
		return v.Visit(ctx.Printstmt())
	}

	return Value{value: false} // if it's not any of the above return false
}

// override the visitAssignstmt function
func (v *Visitor) VisitAssignstmt(ctx *parser.AssignstmtContext) Value {
	id := ctx.ID().GetText()   // get the id
	val := v.Visit(ctx.Expr()) // get the value
	v.memory[id] = val         // assign the value to the id
	return Value{value: true}
}

// override the visitIfstmt function
func (v *Visitor) VisitIfstmt(ctx *parser.IfstmtContext) Value {
	value, ok := v.Visit(ctx.Expr()).value.(bool) // get the value of the expression

	if ok && value { // if the value is true
		return v.Visit(ctx.Block()) // visit the first block
	}
	return Value{value: false}
}

// override the visitPrintlnstmt function
func (v *Visitor) VisitPrintlnstmt(ctx *parser.PrintlnstmtContext) Value {
	val := v.Visit(ctx.Expr()).value // get the value
	fmt.Println(val)                 // print the value
	return Value{value: true}
}

// override the visitPrintstmt function
func (v *Visitor) VisitPrintstmt(ctx *parser.PrintstmtContext) Value {
	val := v.Visit(ctx.Expr()).value // get the value
	fmt.Print(val)                   // print the value
	return Value{value: true}
}

// override the visitWhilestmt function
func (v *Visitor) VisitWhilestmt(ctx *parser.WhilestmtContext) Value {
	value, ok := v.Visit(ctx.Expr()).value.(bool) // get the value of the expression

	for ok && value { // if the value is true
		v.Visit(ctx.Block()) // visit the block
		value, ok = v.Visit(ctx.Expr()).value.(bool)
	}
	return Value{value: true}
}

// override the visitForstmt function
func (v *Visitor) VisitForstmt(ctx *parser.ForstmtContext) Value {
	// visit the first assign statement
	v.Visit(ctx.AllAssignstmt()[0])

	condValue, ok := v.Visit(ctx.Expr()).value.(bool) // get the value of the expression the condition
	for ok && condValue {
		v.Visit(ctx.Block()) // visit the block
		// update the value using the second assign statement
		v.Visit(ctx.AllAssignstmt()[1])
		// visit the condition again
		condValue = v.Visit(ctx.Expr()).value.(bool)
	}
	return Value{value: true}
}

// override the visitOpExpr function
func (v *Visitor) VisitOpExpr(ctx *parser.OpExprContext) Value {
	l := v.Visit(ctx.GetLeft()).value.(int64)
	r := v.Visit(ctx.GetRight()).value.(int64)
	op := ctx.GetOp().GetText()
	switch op {
	case "+":
		return Value{value: l + r}
	case "-":
		return Value{value: l - r}
	case "*":
		return Value{value: l * r}
	case "/":
		return Value{value: l / r}
	case "<":
		if l < r {
			return Value{value: true}
		} else {
			return Value{value: false}
		}
	case ">":
		if l > r {
			return Value{value: true}
		} else {
			return Value{value: false}
		}
	}
	return Value{value: false}
}

// override the visitIntExpr function
func (v *Visitor) VisitIntExpr(ctx *parser.IntExprContext) Value {
	i, _ := strconv.ParseInt(ctx.GetText(), 10, 64)
	return Value{value: i}
}

// override the visitStrExpr function
func (v *Visitor) VisitStrExpr(ctx *parser.StrExprContext) Value {
	value := strings.Trim(ctx.GetText(), "\"")
	return Value{value: value}
}

// override the visitIdExpr function
func (v *Visitor) VisitIdExpr(ctx *parser.IdExprContext) Value {
	id := ctx.GetText()
	value, ok := v.memory[id]
	if ok {
		return value
	} else {
		log.Println("no such variable: " + id)
	}
	return Value{value: false}
}

// override the visitBoolExpr function
func (v *Visitor) VisitBoolExpr(ctx *parser.BoolExprContext) Value {
	value, _ := strconv.ParseBool(ctx.GetText())
	return Value{value: value}
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		prog := `
		a = true
		b = 0
		if (a) {
			println("a is true!")
		}
		while (b < 5) {
			println(b)
			b = b + 2
		}
		println("-----------------")
	
		for (i = 0; i < 10; i = i + 1) {
			println(i)
		}`
		input := antlr.NewInputStream(prog)             // convert string to stream
		lexer := parser.NewControlLexer(input)          // create lexer
		tokens := antlr.NewCommonTokenStream(lexer, 0)  // create tokens
		p := parser.NewControlParser(tokens)            // create parser
		p.BuildParseTrees = true                        // tell the parser to build parse trees
		tree := p.Prog()                                // parse
		eval := Visitor{memory: make(map[string]Value)} // create visitor
		eval.Visit(tree)
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
