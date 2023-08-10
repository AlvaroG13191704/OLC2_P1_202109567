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

func (v *BaseSFGrammarVisitor) VisitTypeValueDeclaration(ctx *TypeValueDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitTypeOptionalValueDeclaration(ctx *TypeOptionalValueDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitValueDeclaration(ctx *ValueDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitType_declaration(ctx *Type_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitValueAssignment(ctx *ValueAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitPlusAssignment(ctx *PlusAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitMinusAssignment(ctx *MinusAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitIfElseStmt(ctx *IfElseStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitSwitchStmt(ctx *SwitchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitCaseBlock(ctx *CaseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitDefaultBlock(ctx *DefaultBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitWhileStmt(ctx *WhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitEmbbededFunc(ctx *EmbbededFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitExprList(ctx *ExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitStringExpr(ctx *StringExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitArithmeticOperationExpr(ctx *ArithmeticOperationExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitNilExpr(ctx *NilExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitRelationalOperationExpr(ctx *RelationalOperationExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitIdExpr(ctx *IdExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitLogicalOperationExpr(ctx *LogicalOperationExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitNegExpr(ctx *NegExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitDigitExpr(ctx *DigitExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitComparationOperationExpr(ctx *ComparationOperationExprContext) interface{} {
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

func (v *BaseSFGrammarVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}
