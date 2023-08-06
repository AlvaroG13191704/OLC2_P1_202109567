// Code generated from SFGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SFGrammar
import "github.com/antlr4-go/antlr/v4"

type BaseSFGrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSFGrammarVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitStmts(ctx *StmtsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitPrintstmt(ctx *PrintstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitStringExpr(ctx *StringExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitIdExpr(ctx *IdExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitOperationExpr(ctx *OperationExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitNegExpr(ctx *NegExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitDigitExpr(ctx *DigitExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitBooleanExpr(ctx *BooleanExprContext) interface{} {
	return v.VisitChildren(ctx)
}