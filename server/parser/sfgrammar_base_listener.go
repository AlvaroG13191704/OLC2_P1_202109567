// Code generated from SFGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SFGrammar
import "github.com/antlr4-go/antlr/v4"

// BaseSFGrammarListener is a complete listener for a parse tree produced by SFGrammarParser.
type BaseSFGrammarListener struct{}

var _ SFGrammarListener = &BaseSFGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSFGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSFGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSFGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSFGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseSFGrammarListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseSFGrammarListener) ExitStart(ctx *StartContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseSFGrammarListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseSFGrammarListener) ExitBlock(ctx *BlockContext) {}

// EnterStmts is called when production stmts is entered.
func (s *BaseSFGrammarListener) EnterStmts(ctx *StmtsContext) {}

// ExitStmts is called when production stmts is exited.
func (s *BaseSFGrammarListener) ExitStmts(ctx *StmtsContext) {}

// EnterBreakStmt is called when production BreakStmt is entered.
func (s *BaseSFGrammarListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production BreakStmt is exited.
func (s *BaseSFGrammarListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterContinueStmt is called when production ContinueStmt is entered.
func (s *BaseSFGrammarListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production ContinueStmt is exited.
func (s *BaseSFGrammarListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterReturnStmt is called when production ReturnStmt is entered.
func (s *BaseSFGrammarListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production ReturnStmt is exited.
func (s *BaseSFGrammarListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterStructStmt is called when production structStmt is entered.
func (s *BaseSFGrammarListener) EnterStructStmt(ctx *StructStmtContext) {}

// ExitStructStmt is called when production structStmt is exited.
func (s *BaseSFGrammarListener) ExitStructStmt(ctx *StructStmtContext) {}

// EnterStructBlock is called when production structBlock is entered.
func (s *BaseSFGrammarListener) EnterStructBlock(ctx *StructBlockContext) {}

// ExitStructBlock is called when production structBlock is exited.
func (s *BaseSFGrammarListener) ExitStructBlock(ctx *StructBlockContext) {}

// EnterStructStmts is called when production structStmts is entered.
func (s *BaseSFGrammarListener) EnterStructStmts(ctx *StructStmtsContext) {}

// ExitStructStmts is called when production structStmts is exited.
func (s *BaseSFGrammarListener) ExitStructStmts(ctx *StructStmtsContext) {}

// EnterStructDeclarationWithValueAndType is called when production StructDeclarationWithValueAndType is entered.
func (s *BaseSFGrammarListener) EnterStructDeclarationWithValueAndType(ctx *StructDeclarationWithValueAndTypeContext) {
}

// ExitStructDeclarationWithValueAndType is called when production StructDeclarationWithValueAndType is exited.
func (s *BaseSFGrammarListener) ExitStructDeclarationWithValueAndType(ctx *StructDeclarationWithValueAndTypeContext) {
}

// EnterStructDeclarationWithoutValue is called when production StructDeclarationWithoutValue is entered.
func (s *BaseSFGrammarListener) EnterStructDeclarationWithoutValue(ctx *StructDeclarationWithoutValueContext) {
}

// ExitStructDeclarationWithoutValue is called when production StructDeclarationWithoutValue is exited.
func (s *BaseSFGrammarListener) ExitStructDeclarationWithoutValue(ctx *StructDeclarationWithoutValueContext) {
}

// EnterStructDeclarationImplicitValue is called when production StructDeclarationImplicitValue is entered.
func (s *BaseSFGrammarListener) EnterStructDeclarationImplicitValue(ctx *StructDeclarationImplicitValueContext) {
}

// ExitStructDeclarationImplicitValue is called when production StructDeclarationImplicitValue is exited.
func (s *BaseSFGrammarListener) ExitStructDeclarationImplicitValue(ctx *StructDeclarationImplicitValueContext) {
}

// EnterStructDeclarationVector is called when production StructDeclarationVector is entered.
func (s *BaseSFGrammarListener) EnterStructDeclarationVector(ctx *StructDeclarationVectorContext) {}

// ExitStructDeclarationVector is called when production StructDeclarationVector is exited.
func (s *BaseSFGrammarListener) ExitStructDeclarationVector(ctx *StructDeclarationVectorContext) {}

// EnterStructFunctionWithoutParams is called when production StructFunctionWithoutParams is entered.
func (s *BaseSFGrammarListener) EnterStructFunctionWithoutParams(ctx *StructFunctionWithoutParamsContext) {
}

// ExitStructFunctionWithoutParams is called when production StructFunctionWithoutParams is exited.
func (s *BaseSFGrammarListener) ExitStructFunctionWithoutParams(ctx *StructFunctionWithoutParamsContext) {
}

// EnterStructFunctionWithParams is called when production StructFunctionWithParams is entered.
func (s *BaseSFGrammarListener) EnterStructFunctionWithParams(ctx *StructFunctionWithParamsContext) {}

// ExitStructFunctionWithParams is called when production StructFunctionWithParams is exited.
func (s *BaseSFGrammarListener) ExitStructFunctionWithParams(ctx *StructFunctionWithParamsContext) {}

// EnterStructCallList is called when production structCallList is entered.
func (s *BaseSFGrammarListener) EnterStructCallList(ctx *StructCallListContext) {}

// ExitStructCallList is called when production structCallList is exited.
func (s *BaseSFGrammarListener) ExitStructCallList(ctx *StructCallListContext) {}

// EnterStructCreation is called when production StructCreation is entered.
func (s *BaseSFGrammarListener) EnterStructCreation(ctx *StructCreationContext) {}

// ExitStructCreation is called when production StructCreation is exited.
func (s *BaseSFGrammarListener) ExitStructCreation(ctx *StructCreationContext) {}

// EnterTypeValueDeclaration is called when production TypeValueDeclaration is entered.
func (s *BaseSFGrammarListener) EnterTypeValueDeclaration(ctx *TypeValueDeclarationContext) {}

// ExitTypeValueDeclaration is called when production TypeValueDeclaration is exited.
func (s *BaseSFGrammarListener) ExitTypeValueDeclaration(ctx *TypeValueDeclarationContext) {}

// EnterTypeOptionalValueDeclaration is called when production TypeOptionalValueDeclaration is entered.
func (s *BaseSFGrammarListener) EnterTypeOptionalValueDeclaration(ctx *TypeOptionalValueDeclarationContext) {
}

// ExitTypeOptionalValueDeclaration is called when production TypeOptionalValueDeclaration is exited.
func (s *BaseSFGrammarListener) ExitTypeOptionalValueDeclaration(ctx *TypeOptionalValueDeclarationContext) {
}

// EnterValueDeclaration is called when production ValueDeclaration is entered.
func (s *BaseSFGrammarListener) EnterValueDeclaration(ctx *ValueDeclarationContext) {}

// ExitValueDeclaration is called when production ValueDeclaration is exited.
func (s *BaseSFGrammarListener) ExitValueDeclaration(ctx *ValueDeclarationContext) {}

// EnterVectorDeclaration is called when production VectorDeclaration is entered.
func (s *BaseSFGrammarListener) EnterVectorDeclaration(ctx *VectorDeclarationContext) {}

// ExitVectorDeclaration is called when production VectorDeclaration is exited.
func (s *BaseSFGrammarListener) ExitVectorDeclaration(ctx *VectorDeclarationContext) {}

// EnterType_declaration is called when production type_declaration is entered.
func (s *BaseSFGrammarListener) EnterType_declaration(ctx *Type_declarationContext) {}

// ExitType_declaration is called when production type_declaration is exited.
func (s *BaseSFGrammarListener) ExitType_declaration(ctx *Type_declarationContext) {}

// EnterValueAssignment is called when production ValueAssignment is entered.
func (s *BaseSFGrammarListener) EnterValueAssignment(ctx *ValueAssignmentContext) {}

// ExitValueAssignment is called when production ValueAssignment is exited.
func (s *BaseSFGrammarListener) ExitValueAssignment(ctx *ValueAssignmentContext) {}

// EnterPlusAssignment is called when production PlusAssignment is entered.
func (s *BaseSFGrammarListener) EnterPlusAssignment(ctx *PlusAssignmentContext) {}

// ExitPlusAssignment is called when production PlusAssignment is exited.
func (s *BaseSFGrammarListener) ExitPlusAssignment(ctx *PlusAssignmentContext) {}

// EnterMinusAssignment is called when production MinusAssignment is entered.
func (s *BaseSFGrammarListener) EnterMinusAssignment(ctx *MinusAssignmentContext) {}

// ExitMinusAssignment is called when production MinusAssignment is exited.
func (s *BaseSFGrammarListener) ExitMinusAssignment(ctx *MinusAssignmentContext) {}

// EnterVectorAssignment is called when production VectorAssignment is entered.
func (s *BaseSFGrammarListener) EnterVectorAssignment(ctx *VectorAssignmentContext) {}

// ExitVectorAssignment is called when production VectorAssignment is exited.
func (s *BaseSFGrammarListener) ExitVectorAssignment(ctx *VectorAssignmentContext) {}

// EnterVectorMinusAssignment is called when production VectorMinusAssignment is entered.
func (s *BaseSFGrammarListener) EnterVectorMinusAssignment(ctx *VectorMinusAssignmentContext) {}

// ExitVectorMinusAssignment is called when production VectorMinusAssignment is exited.
func (s *BaseSFGrammarListener) ExitVectorMinusAssignment(ctx *VectorMinusAssignmentContext) {}

// EnterVectorPlusAssignment is called when production VectorPlusAssignment is entered.
func (s *BaseSFGrammarListener) EnterVectorPlusAssignment(ctx *VectorPlusAssignmentContext) {}

// ExitVectorPlusAssignment is called when production VectorPlusAssignment is exited.
func (s *BaseSFGrammarListener) ExitVectorPlusAssignment(ctx *VectorPlusAssignmentContext) {}

// EnterIfElseStmt is called when production IfElseStmt is entered.
func (s *BaseSFGrammarListener) EnterIfElseStmt(ctx *IfElseStmtContext) {}

// ExitIfElseStmt is called when production IfElseStmt is exited.
func (s *BaseSFGrammarListener) ExitIfElseStmt(ctx *IfElseStmtContext) {}

// EnterElseIfStmt is called when production ElseIfStmt is entered.
func (s *BaseSFGrammarListener) EnterElseIfStmt(ctx *ElseIfStmtContext) {}

// ExitElseIfStmt is called when production ElseIfStmt is exited.
func (s *BaseSFGrammarListener) ExitElseIfStmt(ctx *ElseIfStmtContext) {}

// EnterSwitchStmt is called when production switchStmt is entered.
func (s *BaseSFGrammarListener) EnterSwitchStmt(ctx *SwitchStmtContext) {}

// ExitSwitchStmt is called when production switchStmt is exited.
func (s *BaseSFGrammarListener) ExitSwitchStmt(ctx *SwitchStmtContext) {}

// EnterCaseBlock is called when production caseBlock is entered.
func (s *BaseSFGrammarListener) EnterCaseBlock(ctx *CaseBlockContext) {}

// ExitCaseBlock is called when production caseBlock is exited.
func (s *BaseSFGrammarListener) ExitCaseBlock(ctx *CaseBlockContext) {}

// EnterDefaultBlock is called when production defaultBlock is entered.
func (s *BaseSFGrammarListener) EnterDefaultBlock(ctx *DefaultBlockContext) {}

// ExitDefaultBlock is called when production defaultBlock is exited.
func (s *BaseSFGrammarListener) ExitDefaultBlock(ctx *DefaultBlockContext) {}

// EnterWhileStmt is called when production whileStmt is entered.
func (s *BaseSFGrammarListener) EnterWhileStmt(ctx *WhileStmtContext) {}

// ExitWhileStmt is called when production whileStmt is exited.
func (s *BaseSFGrammarListener) ExitWhileStmt(ctx *WhileStmtContext) {}

// EnterForRangeExpr is called when production ForRangeExpr is entered.
func (s *BaseSFGrammarListener) EnterForRangeExpr(ctx *ForRangeExprContext) {}

// ExitForRangeExpr is called when production ForRangeExpr is exited.
func (s *BaseSFGrammarListener) ExitForRangeExpr(ctx *ForRangeExprContext) {}

// EnterForExpr is called when production ForExpr is entered.
func (s *BaseSFGrammarListener) EnterForExpr(ctx *ForExprContext) {}

// ExitForExpr is called when production ForExpr is exited.
func (s *BaseSFGrammarListener) ExitForExpr(ctx *ForExprContext) {}

// EnterForRange is called when production forRange is entered.
func (s *BaseSFGrammarListener) EnterForRange(ctx *ForRangeContext) {}

// ExitForRange is called when production forRange is exited.
func (s *BaseSFGrammarListener) ExitForRange(ctx *ForRangeContext) {}

// EnterGuardStmt is called when production guardStmt is entered.
func (s *BaseSFGrammarListener) EnterGuardStmt(ctx *GuardStmtContext) {}

// ExitGuardStmt is called when production guardStmt is exited.
func (s *BaseSFGrammarListener) ExitGuardStmt(ctx *GuardStmtContext) {}

// EnterFunctionWithoutParams is called when production FunctionWithoutParams is entered.
func (s *BaseSFGrammarListener) EnterFunctionWithoutParams(ctx *FunctionWithoutParamsContext) {}

// ExitFunctionWithoutParams is called when production FunctionWithoutParams is exited.
func (s *BaseSFGrammarListener) ExitFunctionWithoutParams(ctx *FunctionWithoutParamsContext) {}

// EnterFunctionWithParams is called when production FunctionWithParams is entered.
func (s *BaseSFGrammarListener) EnterFunctionWithParams(ctx *FunctionWithParamsContext) {}

// ExitFunctionWithParams is called when production FunctionWithParams is exited.
func (s *BaseSFGrammarListener) ExitFunctionWithParams(ctx *FunctionWithParamsContext) {}

// EnterListFunctionParamsEI is called when production listFunctionParamsEI is entered.
func (s *BaseSFGrammarListener) EnterListFunctionParamsEI(ctx *ListFunctionParamsEIContext) {}

// ExitListFunctionParamsEI is called when production listFunctionParamsEI is exited.
func (s *BaseSFGrammarListener) ExitListFunctionParamsEI(ctx *ListFunctionParamsEIContext) {}

// EnterListFunctionParamsNEI is called when production listFunctionParamsNEI is entered.
func (s *BaseSFGrammarListener) EnterListFunctionParamsNEI(ctx *ListFunctionParamsNEIContext) {}

// ExitListFunctionParamsNEI is called when production listFunctionParamsNEI is exited.
func (s *BaseSFGrammarListener) ExitListFunctionParamsNEI(ctx *ListFunctionParamsNEIContext) {}

// EnterListFunctionParamsBEI is called when production listFunctionParamsBEI is entered.
func (s *BaseSFGrammarListener) EnterListFunctionParamsBEI(ctx *ListFunctionParamsBEIContext) {}

// ExitListFunctionParamsBEI is called when production listFunctionParamsBEI is exited.
func (s *BaseSFGrammarListener) ExitListFunctionParamsBEI(ctx *ListFunctionParamsBEIContext) {}

// EnterListFunctionParamsEIVector is called when production listFunctionParamsEIVector is entered.
func (s *BaseSFGrammarListener) EnterListFunctionParamsEIVector(ctx *ListFunctionParamsEIVectorContext) {
}

// ExitListFunctionParamsEIVector is called when production listFunctionParamsEIVector is exited.
func (s *BaseSFGrammarListener) ExitListFunctionParamsEIVector(ctx *ListFunctionParamsEIVectorContext) {
}

// EnterListFunctionParamsNEIVector is called when production listFunctionParamsNEIVector is entered.
func (s *BaseSFGrammarListener) EnterListFunctionParamsNEIVector(ctx *ListFunctionParamsNEIVectorContext) {
}

// ExitListFunctionParamsNEIVector is called when production listFunctionParamsNEIVector is exited.
func (s *BaseSFGrammarListener) ExitListFunctionParamsNEIVector(ctx *ListFunctionParamsNEIVectorContext) {
}

// EnterListFunctionParamsBEIVector is called when production listFunctionParamsBEIVector is entered.
func (s *BaseSFGrammarListener) EnterListFunctionParamsBEIVector(ctx *ListFunctionParamsBEIVectorContext) {
}

// ExitListFunctionParamsBEIVector is called when production listFunctionParamsBEIVector is exited.
func (s *BaseSFGrammarListener) ExitListFunctionParamsBEIVector(ctx *ListFunctionParamsBEIVectorContext) {
}

// EnterCallFunctionWithoutParams is called when production CallFunctionWithoutParams is entered.
func (s *BaseSFGrammarListener) EnterCallFunctionWithoutParams(ctx *CallFunctionWithoutParamsContext) {
}

// ExitCallFunctionWithoutParams is called when production CallFunctionWithoutParams is exited.
func (s *BaseSFGrammarListener) ExitCallFunctionWithoutParams(ctx *CallFunctionWithoutParamsContext) {
}

// EnterCallFunctionWithParams is called when production CallFunctionWithParams is entered.
func (s *BaseSFGrammarListener) EnterCallFunctionWithParams(ctx *CallFunctionWithParamsContext) {}

// ExitCallFunctionWithParams is called when production CallFunctionWithParams is exited.
func (s *BaseSFGrammarListener) ExitCallFunctionWithParams(ctx *CallFunctionWithParamsContext) {}

// EnterListCallFunctionStmtEI is called when production listCallFunctionStmtEI is entered.
func (s *BaseSFGrammarListener) EnterListCallFunctionStmtEI(ctx *ListCallFunctionStmtEIContext) {}

// ExitListCallFunctionStmtEI is called when production listCallFunctionStmtEI is exited.
func (s *BaseSFGrammarListener) ExitListCallFunctionStmtEI(ctx *ListCallFunctionStmtEIContext) {}

// EnterListCallFunctionStmtNEI is called when production listCallFunctionStmtNEI is entered.
func (s *BaseSFGrammarListener) EnterListCallFunctionStmtNEI(ctx *ListCallFunctionStmtNEIContext) {}

// ExitListCallFunctionStmtNEI is called when production listCallFunctionStmtNEI is exited.
func (s *BaseSFGrammarListener) ExitListCallFunctionStmtNEI(ctx *ListCallFunctionStmtNEIContext) {}

// EnterAppendVector is called when production AppendVector is entered.
func (s *BaseSFGrammarListener) EnterAppendVector(ctx *AppendVectorContext) {}

// ExitAppendVector is called when production AppendVector is exited.
func (s *BaseSFGrammarListener) ExitAppendVector(ctx *AppendVectorContext) {}

// EnterRemoveLastVector is called when production RemoveLastVector is entered.
func (s *BaseSFGrammarListener) EnterRemoveLastVector(ctx *RemoveLastVectorContext) {}

// ExitRemoveLastVector is called when production RemoveLastVector is exited.
func (s *BaseSFGrammarListener) ExitRemoveLastVector(ctx *RemoveLastVectorContext) {}

// EnterRemoveAtVector is called when production RemoveAtVector is entered.
func (s *BaseSFGrammarListener) EnterRemoveAtVector(ctx *RemoveAtVectorContext) {}

// ExitRemoveAtVector is called when production RemoveAtVector is exited.
func (s *BaseSFGrammarListener) ExitRemoveAtVector(ctx *RemoveAtVectorContext) {}

// EnterIsEmptyVector is called when production IsEmptyVector is entered.
func (s *BaseSFGrammarListener) EnterIsEmptyVector(ctx *IsEmptyVectorContext) {}

// ExitIsEmptyVector is called when production IsEmptyVector is exited.
func (s *BaseSFGrammarListener) ExitIsEmptyVector(ctx *IsEmptyVectorContext) {}

// EnterCountVector is called when production CountVector is entered.
func (s *BaseSFGrammarListener) EnterCountVector(ctx *CountVectorContext) {}

// ExitCountVector is called when production CountVector is exited.
func (s *BaseSFGrammarListener) ExitCountVector(ctx *CountVectorContext) {}

// EnterAccessVector is called when production AccessVector is entered.
func (s *BaseSFGrammarListener) EnterAccessVector(ctx *AccessVectorContext) {}

// ExitAccessVector is called when production AccessVector is exited.
func (s *BaseSFGrammarListener) ExitAccessVector(ctx *AccessVectorContext) {}

// EnterStructCallFunction is called when production StructCallFunction is entered.
func (s *BaseSFGrammarListener) EnterStructCallFunction(ctx *StructCallFunctionContext) {}

// ExitStructCallFunction is called when production StructCallFunction is exited.
func (s *BaseSFGrammarListener) ExitStructCallFunction(ctx *StructCallFunctionContext) {}

// EnterStructAttribute is called when production StructAttribute is entered.
func (s *BaseSFGrammarListener) EnterStructAttribute(ctx *StructAttributeContext) {}

// ExitStructAttribute is called when production StructAttribute is exited.
func (s *BaseSFGrammarListener) ExitStructAttribute(ctx *StructAttributeContext) {}

// EnterSelfFunction is called when production SelfFunction is entered.
func (s *BaseSFGrammarListener) EnterSelfFunction(ctx *SelfFunctionContext) {}

// ExitSelfFunction is called when production SelfFunction is exited.
func (s *BaseSFGrammarListener) ExitSelfFunction(ctx *SelfFunctionContext) {}

// EnterEmbbededFunc is called when production embbededFunc is entered.
func (s *BaseSFGrammarListener) EnterEmbbededFunc(ctx *EmbbededFuncContext) {}

// ExitEmbbededFunc is called when production embbededFunc is exited.
func (s *BaseSFGrammarListener) ExitEmbbededFunc(ctx *EmbbededFuncContext) {}

// EnterPrintstmt is called when production printstmt is entered.
func (s *BaseSFGrammarListener) EnterPrintstmt(ctx *PrintstmtContext) {}

// ExitPrintstmt is called when production printstmt is exited.
func (s *BaseSFGrammarListener) ExitPrintstmt(ctx *PrintstmtContext) {}

// EnterExprList is called when production exprList is entered.
func (s *BaseSFGrammarListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production exprList is exited.
func (s *BaseSFGrammarListener) ExitExprList(ctx *ExprListContext) {}

// EnterIntstmt is called when production intstmt is entered.
func (s *BaseSFGrammarListener) EnterIntstmt(ctx *IntstmtContext) {}

// ExitIntstmt is called when production intstmt is exited.
func (s *BaseSFGrammarListener) ExitIntstmt(ctx *IntstmtContext) {}

// EnterFloatstmt is called when production floatstmt is entered.
func (s *BaseSFGrammarListener) EnterFloatstmt(ctx *FloatstmtContext) {}

// ExitFloatstmt is called when production floatstmt is exited.
func (s *BaseSFGrammarListener) ExitFloatstmt(ctx *FloatstmtContext) {}

// EnterStringstmt is called when production stringstmt is entered.
func (s *BaseSFGrammarListener) EnterStringstmt(ctx *StringstmtContext) {}

// ExitStringstmt is called when production stringstmt is exited.
func (s *BaseSFGrammarListener) ExitStringstmt(ctx *StringstmtContext) {}

// EnterStringExpr is called when production StringExpr is entered.
func (s *BaseSFGrammarListener) EnterStringExpr(ctx *StringExprContext) {}

// ExitStringExpr is called when production StringExpr is exited.
func (s *BaseSFGrammarListener) ExitStringExpr(ctx *StringExprContext) {}

// EnterEmbeddedFunctionExpr is called when production EmbeddedFunctionExpr is entered.
func (s *BaseSFGrammarListener) EnterEmbeddedFunctionExpr(ctx *EmbeddedFunctionExprContext) {}

// ExitEmbeddedFunctionExpr is called when production EmbeddedFunctionExpr is exited.
func (s *BaseSFGrammarListener) ExitEmbeddedFunctionExpr(ctx *EmbeddedFunctionExprContext) {}

// EnterNilExpr is called when production NilExpr is entered.
func (s *BaseSFGrammarListener) EnterNilExpr(ctx *NilExprContext) {}

// ExitNilExpr is called when production NilExpr is exited.
func (s *BaseSFGrammarListener) ExitNilExpr(ctx *NilExprContext) {}

// EnterIdExpr is called when production IdExpr is entered.
func (s *BaseSFGrammarListener) EnterIdExpr(ctx *IdExprContext) {}

// ExitIdExpr is called when production IdExpr is exited.
func (s *BaseSFGrammarListener) ExitIdExpr(ctx *IdExprContext) {}

// EnterCallBackExpr is called when production CallBackExpr is entered.
func (s *BaseSFGrammarListener) EnterCallBackExpr(ctx *CallBackExprContext) {}

// ExitCallBackExpr is called when production CallBackExpr is exited.
func (s *BaseSFGrammarListener) ExitCallBackExpr(ctx *CallBackExprContext) {}

// EnterLogicalOperationExpr is called when production LogicalOperationExpr is entered.
func (s *BaseSFGrammarListener) EnterLogicalOperationExpr(ctx *LogicalOperationExprContext) {}

// ExitLogicalOperationExpr is called when production LogicalOperationExpr is exited.
func (s *BaseSFGrammarListener) ExitLogicalOperationExpr(ctx *LogicalOperationExprContext) {}

// EnterNegExpr is called when production NegExpr is entered.
func (s *BaseSFGrammarListener) EnterNegExpr(ctx *NegExprContext) {}

// ExitNegExpr is called when production NegExpr is exited.
func (s *BaseSFGrammarListener) ExitNegExpr(ctx *NegExprContext) {}

// EnterComparationOperationExpr is called when production ComparationOperationExpr is entered.
func (s *BaseSFGrammarListener) EnterComparationOperationExpr(ctx *ComparationOperationExprContext) {}

// ExitComparationOperationExpr is called when production ComparationOperationExpr is exited.
func (s *BaseSFGrammarListener) ExitComparationOperationExpr(ctx *ComparationOperationExprContext) {}

// EnterStructAsArgument is called when production StructAsArgument is entered.
func (s *BaseSFGrammarListener) EnterStructAsArgument(ctx *StructAsArgumentContext) {}

// ExitStructAsArgument is called when production StructAsArgument is exited.
func (s *BaseSFGrammarListener) ExitStructAsArgument(ctx *StructAsArgumentContext) {}

// EnterArithmeticOperationExpr is called when production ArithmeticOperationExpr is entered.
func (s *BaseSFGrammarListener) EnterArithmeticOperationExpr(ctx *ArithmeticOperationExprContext) {}

// ExitArithmeticOperationExpr is called when production ArithmeticOperationExpr is exited.
func (s *BaseSFGrammarListener) ExitArithmeticOperationExpr(ctx *ArithmeticOperationExprContext) {}

// EnterRelationalOperationExpr is called when production RelationalOperationExpr is entered.
func (s *BaseSFGrammarListener) EnterRelationalOperationExpr(ctx *RelationalOperationExprContext) {}

// ExitRelationalOperationExpr is called when production RelationalOperationExpr is exited.
func (s *BaseSFGrammarListener) ExitRelationalOperationExpr(ctx *RelationalOperationExprContext) {}

// EnterDigitExpr is called when production DigitExpr is entered.
func (s *BaseSFGrammarListener) EnterDigitExpr(ctx *DigitExprContext) {}

// ExitDigitExpr is called when production DigitExpr is exited.
func (s *BaseSFGrammarListener) ExitDigitExpr(ctx *DigitExprContext) {}

// EnterNotExpr is called when production NotExpr is entered.
func (s *BaseSFGrammarListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production NotExpr is exited.
func (s *BaseSFGrammarListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterParenExpr is called when production ParenExpr is entered.
func (s *BaseSFGrammarListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production ParenExpr is exited.
func (s *BaseSFGrammarListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterCallFunctionExpr is called when production CallFunctionExpr is entered.
func (s *BaseSFGrammarListener) EnterCallFunctionExpr(ctx *CallFunctionExprContext) {}

// ExitCallFunctionExpr is called when production CallFunctionExpr is exited.
func (s *BaseSFGrammarListener) ExitCallFunctionExpr(ctx *CallFunctionExprContext) {}

// EnterBooleanExpr is called when production BooleanExpr is entered.
func (s *BaseSFGrammarListener) EnterBooleanExpr(ctx *BooleanExprContext) {}

// ExitBooleanExpr is called when production BooleanExpr is exited.
func (s *BaseSFGrammarListener) ExitBooleanExpr(ctx *BooleanExprContext) {}

// EnterType is called when production type is entered.
func (s *BaseSFGrammarListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseSFGrammarListener) ExitType(ctx *TypeContext) {}
