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

func (v *BaseSFGrammarVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
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

func (v *BaseSFGrammarVisitor) VisitElseIfStmt(ctx *ElseIfStmtContext) interface{} {
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

func (v *BaseSFGrammarVisitor) VisitForRangeExpr(ctx *ForRangeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitForExpr(ctx *ForExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitForRange(ctx *ForRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitGuardStmt(ctx *GuardStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitFunctionWithoutParams(ctx *FunctionWithoutParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitFunctionWithParams(ctx *FunctionWithParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitListFunctionParamsEI(ctx *ListFunctionParamsEIContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitListFunctionParamsNEI(ctx *ListFunctionParamsNEIContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitListFunctionParamsBEI(ctx *ListFunctionParamsBEIContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitCallFunctionWithoutParams(ctx *CallFunctionWithoutParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitCallFunctionWithParamsEI(ctx *CallFunctionWithParamsEIContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitListCallFunctionStmtEI(ctx *ListCallFunctionStmtEIContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitListCallFunctionStmtNEI(ctx *ListCallFunctionStmtNEIContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitEmbbededFunc(ctx *EmbbededFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitPrintstmt(ctx *PrintstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitExprList(ctx *ExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitStringExpr(ctx *StringExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitNilExpr(ctx *NilExprContext) interface{} {
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

func (v *BaseSFGrammarVisitor) VisitComparationOperationExpr(ctx *ComparationOperationExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitArithmeticOperationExpr(ctx *ArithmeticOperationExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitRelationalOperationExpr(ctx *RelationalOperationExprContext) interface{} {
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

func (v *BaseSFGrammarVisitor) VisitCallFunctionExpr(ctx *CallFunctionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitBooleanExpr(ctx *BooleanExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}
