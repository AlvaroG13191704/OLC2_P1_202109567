// Code generated from SFGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SFGrammar
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SFGrammarParser.
type SFGrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SFGrammarParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by SFGrammarParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by SFGrammarParser#stmts.
	VisitStmts(ctx *StmtsContext) interface{}

	// Visit a parse tree produced by SFGrammarParser#printstmt.
	VisitPrintstmt(ctx *PrintstmtContext) interface{}

	// Visit a parse tree produced by SFGrammarParser#TypeValueDeclaration.
	VisitTypeValueDeclaration(ctx *TypeValueDeclarationContext) interface{}

	// Visit a parse tree produced by SFGrammarParser#TypeOptionalValueDeclaration.
	VisitTypeOptionalValueDeclaration(ctx *TypeOptionalValueDeclarationContext) interface{}

	// Visit a parse tree produced by SFGrammarParser#ValueDeclaration.
	VisitValueDeclaration(ctx *ValueDeclarationContext) interface{}

	// Visit a parse tree produced by SFGrammarParser#type_declaration.
	VisitType_declaration(ctx *Type_declarationContext) interface{}

	// Visit a parse tree produced by SFGrammarParser#ValueAssignment.
	VisitValueAssignment(ctx *ValueAssignmentContext) interface{}

	// Visit a parse tree produced by SFGrammarParser#PlusAssignment.
	VisitPlusAssignment(ctx *PlusAssignmentContext) interface{}

	// Visit a parse tree produced by SFGrammarParser#MinusAssignment.
	VisitMinusAssignment(ctx *MinusAssignmentContext) interface{}

	// Visit a parse tree produced by SFGrammarParser#exprList.
	VisitExprList(ctx *ExprListContext) interface{}

	// Visit a parse tree produced by SFGrammarParser#StringExpr.
	VisitStringExpr(ctx *StringExprContext) interface{}

	// Visit a parse tree produced by SFGrammarParser#ArithmeticOperationExpr.
	VisitArithmeticOperationExpr(ctx *ArithmeticOperationExprContext) interface{}

	// Visit a parse tree produced by SFGrammarParser#NilExpr.
	VisitNilExpr(ctx *NilExprContext) interface{}

	// Visit a parse tree produced by SFGrammarParser#IdExpr.
	VisitIdExpr(ctx *IdExprContext) interface{}

	// Visit a parse tree produced by SFGrammarParser#NegExpr.
	VisitNegExpr(ctx *NegExprContext) interface{}

	// Visit a parse tree produced by SFGrammarParser#DigitExpr.
	VisitDigitExpr(ctx *DigitExprContext) interface{}

	// Visit a parse tree produced by SFGrammarParser#NotExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by SFGrammarParser#ParenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by SFGrammarParser#BooleanExpr.
	VisitBooleanExpr(ctx *BooleanExprContext) interface{}

	// Visit a parse tree produced by SFGrammarParser#type.
	VisitType(ctx *TypeContext) interface{}
}
