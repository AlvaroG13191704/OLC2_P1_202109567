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

func (v *BaseSFGrammarVisitor) VisitStructStmt(ctx *StructStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitStructBlock(ctx *StructBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitStructStmts(ctx *StructStmtsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitStructDeclarationWithValueAndType(ctx *StructDeclarationWithValueAndTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitStructDeclarationWithoutValue(ctx *StructDeclarationWithoutValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitStructDeclarationImplicitValue(ctx *StructDeclarationImplicitValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitStructDeclarationVector(ctx *StructDeclarationVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitStructFunctionWithoutParams(ctx *StructFunctionWithoutParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitStructFunctionWithParams(ctx *StructFunctionWithParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitStructCallList(ctx *StructCallListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitStructCreation(ctx *StructCreationContext) interface{} {
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

func (v *BaseSFGrammarVisitor) VisitVectorOfStructDeclaration(ctx *VectorOfStructDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitVectorDeclaration(ctx *VectorDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitVectorOfStructCreation(ctx *VectorOfStructCreationContext) interface{} {
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

func (v *BaseSFGrammarVisitor) VisitVectorAssignment(ctx *VectorAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitVectorMinusAssignment(ctx *VectorMinusAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitVectorPlusAssignment(ctx *VectorPlusAssignmentContext) interface{} {
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

func (v *BaseSFGrammarVisitor) VisitListFunctionParamsEIVector(ctx *ListFunctionParamsEIVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitListFunctionParamsNEIVector(ctx *ListFunctionParamsNEIVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitListFunctionParamsBEIVector(ctx *ListFunctionParamsBEIVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitCallFunctionWithoutParams(ctx *CallFunctionWithoutParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitCallFunctionWithParams(ctx *CallFunctionWithParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitListCallFunctionStmtEI(ctx *ListCallFunctionStmtEIContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitListCallFunctionStmtNEI(ctx *ListCallFunctionStmtNEIContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitAppendVector(ctx *AppendVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitRemoveLastVector(ctx *RemoveLastVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitRemoveAtVector(ctx *RemoveAtVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitIsEmptyVector(ctx *IsEmptyVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitCountVector(ctx *CountVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitAccessVectorStruct(ctx *AccessVectorStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitAccessVector(ctx *AccessVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitStructCallFunction(ctx *StructCallFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitStructAttribute(ctx *StructAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitSelfFunction(ctx *SelfFunctionContext) interface{} {
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

func (v *BaseSFGrammarVisitor) VisitIntstmt(ctx *IntstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitFloatstmt(ctx *FloatstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitStringstmt(ctx *StringstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitStringExpr(ctx *StringExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitEmbeddedFunctionExpr(ctx *EmbeddedFunctionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitNilExpr(ctx *NilExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitIdExpr(ctx *IdExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSFGrammarVisitor) VisitCallBackExpr(ctx *CallBackExprContext) interface{} {
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

func (v *BaseSFGrammarVisitor) VisitStructAsArgument(ctx *StructAsArgumentContext) interface{} {
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
