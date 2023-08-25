// Code generated from SFGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SFGrammar
import "github.com/antlr4-go/antlr/v4"

// SFGrammarListener is a complete listener for a parse tree produced by SFGrammarParser.
type SFGrammarListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStmts is called when entering the stmts production.
	EnterStmts(c *StmtsContext)

	// EnterBreakStmt is called when entering the BreakStmt production.
	EnterBreakStmt(c *BreakStmtContext)

	// EnterContinueStmt is called when entering the ContinueStmt production.
	EnterContinueStmt(c *ContinueStmtContext)

	// EnterReturnStmt is called when entering the ReturnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterStructStmt is called when entering the structStmt production.
	EnterStructStmt(c *StructStmtContext)

	// EnterStructBlock is called when entering the structBlock production.
	EnterStructBlock(c *StructBlockContext)

	// EnterStructStmts is called when entering the structStmts production.
	EnterStructStmts(c *StructStmtsContext)

	// EnterStructDeclarationWithValueAndType is called when entering the StructDeclarationWithValueAndType production.
	EnterStructDeclarationWithValueAndType(c *StructDeclarationWithValueAndTypeContext)

	// EnterStructDeclarationWithoutValue is called when entering the StructDeclarationWithoutValue production.
	EnterStructDeclarationWithoutValue(c *StructDeclarationWithoutValueContext)

	// EnterStructDeclarationImplicitValue is called when entering the StructDeclarationImplicitValue production.
	EnterStructDeclarationImplicitValue(c *StructDeclarationImplicitValueContext)

	// EnterStructDeclarationVector is called when entering the StructDeclarationVector production.
	EnterStructDeclarationVector(c *StructDeclarationVectorContext)

	// EnterStructFunctionWithoutParams is called when entering the StructFunctionWithoutParams production.
	EnterStructFunctionWithoutParams(c *StructFunctionWithoutParamsContext)

	// EnterStructFunctionWithParams is called when entering the StructFunctionWithParams production.
	EnterStructFunctionWithParams(c *StructFunctionWithParamsContext)

	// EnterStructCallList is called when entering the structCallList production.
	EnterStructCallList(c *StructCallListContext)

	// EnterStructCreation is called when entering the StructCreation production.
	EnterStructCreation(c *StructCreationContext)

	// EnterTypeValueDeclaration is called when entering the TypeValueDeclaration production.
	EnterTypeValueDeclaration(c *TypeValueDeclarationContext)

	// EnterTypeOptionalValueDeclaration is called when entering the TypeOptionalValueDeclaration production.
	EnterTypeOptionalValueDeclaration(c *TypeOptionalValueDeclarationContext)

	// EnterValueDeclaration is called when entering the ValueDeclaration production.
	EnterValueDeclaration(c *ValueDeclarationContext)

	// EnterVectorDeclaration is called when entering the VectorDeclaration production.
	EnterVectorDeclaration(c *VectorDeclarationContext)

	// EnterType_declaration is called when entering the type_declaration production.
	EnterType_declaration(c *Type_declarationContext)

	// EnterValueAssignment is called when entering the ValueAssignment production.
	EnterValueAssignment(c *ValueAssignmentContext)

	// EnterPlusAssignment is called when entering the PlusAssignment production.
	EnterPlusAssignment(c *PlusAssignmentContext)

	// EnterMinusAssignment is called when entering the MinusAssignment production.
	EnterMinusAssignment(c *MinusAssignmentContext)

	// EnterVectorAssignment is called when entering the VectorAssignment production.
	EnterVectorAssignment(c *VectorAssignmentContext)

	// EnterVectorMinusAssignment is called when entering the VectorMinusAssignment production.
	EnterVectorMinusAssignment(c *VectorMinusAssignmentContext)

	// EnterVectorPlusAssignment is called when entering the VectorPlusAssignment production.
	EnterVectorPlusAssignment(c *VectorPlusAssignmentContext)

	// EnterIfElseStmt is called when entering the IfElseStmt production.
	EnterIfElseStmt(c *IfElseStmtContext)

	// EnterElseIfStmt is called when entering the ElseIfStmt production.
	EnterElseIfStmt(c *ElseIfStmtContext)

	// EnterSwitchStmt is called when entering the switchStmt production.
	EnterSwitchStmt(c *SwitchStmtContext)

	// EnterCaseBlock is called when entering the caseBlock production.
	EnterCaseBlock(c *CaseBlockContext)

	// EnterDefaultBlock is called when entering the defaultBlock production.
	EnterDefaultBlock(c *DefaultBlockContext)

	// EnterWhileStmt is called when entering the whileStmt production.
	EnterWhileStmt(c *WhileStmtContext)

	// EnterForRangeExpr is called when entering the ForRangeExpr production.
	EnterForRangeExpr(c *ForRangeExprContext)

	// EnterForExpr is called when entering the ForExpr production.
	EnterForExpr(c *ForExprContext)

	// EnterForRange is called when entering the forRange production.
	EnterForRange(c *ForRangeContext)

	// EnterGuardStmt is called when entering the guardStmt production.
	EnterGuardStmt(c *GuardStmtContext)

	// EnterFunctionWithoutParams is called when entering the FunctionWithoutParams production.
	EnterFunctionWithoutParams(c *FunctionWithoutParamsContext)

	// EnterFunctionWithParams is called when entering the FunctionWithParams production.
	EnterFunctionWithParams(c *FunctionWithParamsContext)

	// EnterListFunctionParamsEI is called when entering the listFunctionParamsEI production.
	EnterListFunctionParamsEI(c *ListFunctionParamsEIContext)

	// EnterListFunctionParamsNEI is called when entering the listFunctionParamsNEI production.
	EnterListFunctionParamsNEI(c *ListFunctionParamsNEIContext)

	// EnterListFunctionParamsBEI is called when entering the listFunctionParamsBEI production.
	EnterListFunctionParamsBEI(c *ListFunctionParamsBEIContext)

	// EnterListFunctionParamsEIVector is called when entering the listFunctionParamsEIVector production.
	EnterListFunctionParamsEIVector(c *ListFunctionParamsEIVectorContext)

	// EnterListFunctionParamsNEIVector is called when entering the listFunctionParamsNEIVector production.
	EnterListFunctionParamsNEIVector(c *ListFunctionParamsNEIVectorContext)

	// EnterListFunctionParamsBEIVector is called when entering the listFunctionParamsBEIVector production.
	EnterListFunctionParamsBEIVector(c *ListFunctionParamsBEIVectorContext)

	// EnterCallFunctionWithoutParams is called when entering the CallFunctionWithoutParams production.
	EnterCallFunctionWithoutParams(c *CallFunctionWithoutParamsContext)

	// EnterCallFunctionWithParams is called when entering the CallFunctionWithParams production.
	EnterCallFunctionWithParams(c *CallFunctionWithParamsContext)

	// EnterListCallFunctionStmtEI is called when entering the listCallFunctionStmtEI production.
	EnterListCallFunctionStmtEI(c *ListCallFunctionStmtEIContext)

	// EnterListCallFunctionStmtNEI is called when entering the listCallFunctionStmtNEI production.
	EnterListCallFunctionStmtNEI(c *ListCallFunctionStmtNEIContext)

	// EnterAppendVector is called when entering the AppendVector production.
	EnterAppendVector(c *AppendVectorContext)

	// EnterRemoveLastVector is called when entering the RemoveLastVector production.
	EnterRemoveLastVector(c *RemoveLastVectorContext)

	// EnterRemoveAtVector is called when entering the RemoveAtVector production.
	EnterRemoveAtVector(c *RemoveAtVectorContext)

	// EnterIsEmptyVector is called when entering the IsEmptyVector production.
	EnterIsEmptyVector(c *IsEmptyVectorContext)

	// EnterCountVector is called when entering the CountVector production.
	EnterCountVector(c *CountVectorContext)

	// EnterAccessVector is called when entering the AccessVector production.
	EnterAccessVector(c *AccessVectorContext)

	// EnterSelfFunction is called when entering the SelfFunction production.
	EnterSelfFunction(c *SelfFunctionContext)

	// EnterStructAttribute is called when entering the StructAttribute production.
	EnterStructAttribute(c *StructAttributeContext)

	// EnterStructFunction is called when entering the StructFunction production.
	EnterStructFunction(c *StructFunctionContext)

	// EnterEmbbededFunc is called when entering the embbededFunc production.
	EnterEmbbededFunc(c *EmbbededFuncContext)

	// EnterPrintstmt is called when entering the printstmt production.
	EnterPrintstmt(c *PrintstmtContext)

	// EnterExprList is called when entering the exprList production.
	EnterExprList(c *ExprListContext)

	// EnterIntstmt is called when entering the intstmt production.
	EnterIntstmt(c *IntstmtContext)

	// EnterFloatstmt is called when entering the floatstmt production.
	EnterFloatstmt(c *FloatstmtContext)

	// EnterStringstmt is called when entering the stringstmt production.
	EnterStringstmt(c *StringstmtContext)

	// EnterStringExpr is called when entering the StringExpr production.
	EnterStringExpr(c *StringExprContext)

	// EnterEmbeddedFunctionExpr is called when entering the EmbeddedFunctionExpr production.
	EnterEmbeddedFunctionExpr(c *EmbeddedFunctionExprContext)

	// EnterNilExpr is called when entering the NilExpr production.
	EnterNilExpr(c *NilExprContext)

	// EnterIdExpr is called when entering the IdExpr production.
	EnterIdExpr(c *IdExprContext)

	// EnterCallBackExpr is called when entering the CallBackExpr production.
	EnterCallBackExpr(c *CallBackExprContext)

	// EnterLogicalOperationExpr is called when entering the LogicalOperationExpr production.
	EnterLogicalOperationExpr(c *LogicalOperationExprContext)

	// EnterNegExpr is called when entering the NegExpr production.
	EnterNegExpr(c *NegExprContext)

	// EnterComparationOperationExpr is called when entering the ComparationOperationExpr production.
	EnterComparationOperationExpr(c *ComparationOperationExprContext)

	// EnterArithmeticOperationExpr is called when entering the ArithmeticOperationExpr production.
	EnterArithmeticOperationExpr(c *ArithmeticOperationExprContext)

	// EnterStructCallExpr is called when entering the StructCallExpr production.
	EnterStructCallExpr(c *StructCallExprContext)

	// EnterRelationalOperationExpr is called when entering the RelationalOperationExpr production.
	EnterRelationalOperationExpr(c *RelationalOperationExprContext)

	// EnterDigitExpr is called when entering the DigitExpr production.
	EnterDigitExpr(c *DigitExprContext)

	// EnterNotExpr is called when entering the NotExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterParenExpr is called when entering the ParenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterCallFunctionExpr is called when entering the CallFunctionExpr production.
	EnterCallFunctionExpr(c *CallFunctionExprContext)

	// EnterBooleanExpr is called when entering the BooleanExpr production.
	EnterBooleanExpr(c *BooleanExprContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStmts is called when exiting the stmts production.
	ExitStmts(c *StmtsContext)

	// ExitBreakStmt is called when exiting the BreakStmt production.
	ExitBreakStmt(c *BreakStmtContext)

	// ExitContinueStmt is called when exiting the ContinueStmt production.
	ExitContinueStmt(c *ContinueStmtContext)

	// ExitReturnStmt is called when exiting the ReturnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitStructStmt is called when exiting the structStmt production.
	ExitStructStmt(c *StructStmtContext)

	// ExitStructBlock is called when exiting the structBlock production.
	ExitStructBlock(c *StructBlockContext)

	// ExitStructStmts is called when exiting the structStmts production.
	ExitStructStmts(c *StructStmtsContext)

	// ExitStructDeclarationWithValueAndType is called when exiting the StructDeclarationWithValueAndType production.
	ExitStructDeclarationWithValueAndType(c *StructDeclarationWithValueAndTypeContext)

	// ExitStructDeclarationWithoutValue is called when exiting the StructDeclarationWithoutValue production.
	ExitStructDeclarationWithoutValue(c *StructDeclarationWithoutValueContext)

	// ExitStructDeclarationImplicitValue is called when exiting the StructDeclarationImplicitValue production.
	ExitStructDeclarationImplicitValue(c *StructDeclarationImplicitValueContext)

	// ExitStructDeclarationVector is called when exiting the StructDeclarationVector production.
	ExitStructDeclarationVector(c *StructDeclarationVectorContext)

	// ExitStructFunctionWithoutParams is called when exiting the StructFunctionWithoutParams production.
	ExitStructFunctionWithoutParams(c *StructFunctionWithoutParamsContext)

	// ExitStructFunctionWithParams is called when exiting the StructFunctionWithParams production.
	ExitStructFunctionWithParams(c *StructFunctionWithParamsContext)

	// ExitStructCallList is called when exiting the structCallList production.
	ExitStructCallList(c *StructCallListContext)

	// ExitStructCreation is called when exiting the StructCreation production.
	ExitStructCreation(c *StructCreationContext)

	// ExitTypeValueDeclaration is called when exiting the TypeValueDeclaration production.
	ExitTypeValueDeclaration(c *TypeValueDeclarationContext)

	// ExitTypeOptionalValueDeclaration is called when exiting the TypeOptionalValueDeclaration production.
	ExitTypeOptionalValueDeclaration(c *TypeOptionalValueDeclarationContext)

	// ExitValueDeclaration is called when exiting the ValueDeclaration production.
	ExitValueDeclaration(c *ValueDeclarationContext)

	// ExitVectorDeclaration is called when exiting the VectorDeclaration production.
	ExitVectorDeclaration(c *VectorDeclarationContext)

	// ExitType_declaration is called when exiting the type_declaration production.
	ExitType_declaration(c *Type_declarationContext)

	// ExitValueAssignment is called when exiting the ValueAssignment production.
	ExitValueAssignment(c *ValueAssignmentContext)

	// ExitPlusAssignment is called when exiting the PlusAssignment production.
	ExitPlusAssignment(c *PlusAssignmentContext)

	// ExitMinusAssignment is called when exiting the MinusAssignment production.
	ExitMinusAssignment(c *MinusAssignmentContext)

	// ExitVectorAssignment is called when exiting the VectorAssignment production.
	ExitVectorAssignment(c *VectorAssignmentContext)

	// ExitVectorMinusAssignment is called when exiting the VectorMinusAssignment production.
	ExitVectorMinusAssignment(c *VectorMinusAssignmentContext)

	// ExitVectorPlusAssignment is called when exiting the VectorPlusAssignment production.
	ExitVectorPlusAssignment(c *VectorPlusAssignmentContext)

	// ExitIfElseStmt is called when exiting the IfElseStmt production.
	ExitIfElseStmt(c *IfElseStmtContext)

	// ExitElseIfStmt is called when exiting the ElseIfStmt production.
	ExitElseIfStmt(c *ElseIfStmtContext)

	// ExitSwitchStmt is called when exiting the switchStmt production.
	ExitSwitchStmt(c *SwitchStmtContext)

	// ExitCaseBlock is called when exiting the caseBlock production.
	ExitCaseBlock(c *CaseBlockContext)

	// ExitDefaultBlock is called when exiting the defaultBlock production.
	ExitDefaultBlock(c *DefaultBlockContext)

	// ExitWhileStmt is called when exiting the whileStmt production.
	ExitWhileStmt(c *WhileStmtContext)

	// ExitForRangeExpr is called when exiting the ForRangeExpr production.
	ExitForRangeExpr(c *ForRangeExprContext)

	// ExitForExpr is called when exiting the ForExpr production.
	ExitForExpr(c *ForExprContext)

	// ExitForRange is called when exiting the forRange production.
	ExitForRange(c *ForRangeContext)

	// ExitGuardStmt is called when exiting the guardStmt production.
	ExitGuardStmt(c *GuardStmtContext)

	// ExitFunctionWithoutParams is called when exiting the FunctionWithoutParams production.
	ExitFunctionWithoutParams(c *FunctionWithoutParamsContext)

	// ExitFunctionWithParams is called when exiting the FunctionWithParams production.
	ExitFunctionWithParams(c *FunctionWithParamsContext)

	// ExitListFunctionParamsEI is called when exiting the listFunctionParamsEI production.
	ExitListFunctionParamsEI(c *ListFunctionParamsEIContext)

	// ExitListFunctionParamsNEI is called when exiting the listFunctionParamsNEI production.
	ExitListFunctionParamsNEI(c *ListFunctionParamsNEIContext)

	// ExitListFunctionParamsBEI is called when exiting the listFunctionParamsBEI production.
	ExitListFunctionParamsBEI(c *ListFunctionParamsBEIContext)

	// ExitListFunctionParamsEIVector is called when exiting the listFunctionParamsEIVector production.
	ExitListFunctionParamsEIVector(c *ListFunctionParamsEIVectorContext)

	// ExitListFunctionParamsNEIVector is called when exiting the listFunctionParamsNEIVector production.
	ExitListFunctionParamsNEIVector(c *ListFunctionParamsNEIVectorContext)

	// ExitListFunctionParamsBEIVector is called when exiting the listFunctionParamsBEIVector production.
	ExitListFunctionParamsBEIVector(c *ListFunctionParamsBEIVectorContext)

	// ExitCallFunctionWithoutParams is called when exiting the CallFunctionWithoutParams production.
	ExitCallFunctionWithoutParams(c *CallFunctionWithoutParamsContext)

	// ExitCallFunctionWithParams is called when exiting the CallFunctionWithParams production.
	ExitCallFunctionWithParams(c *CallFunctionWithParamsContext)

	// ExitListCallFunctionStmtEI is called when exiting the listCallFunctionStmtEI production.
	ExitListCallFunctionStmtEI(c *ListCallFunctionStmtEIContext)

	// ExitListCallFunctionStmtNEI is called when exiting the listCallFunctionStmtNEI production.
	ExitListCallFunctionStmtNEI(c *ListCallFunctionStmtNEIContext)

	// ExitAppendVector is called when exiting the AppendVector production.
	ExitAppendVector(c *AppendVectorContext)

	// ExitRemoveLastVector is called when exiting the RemoveLastVector production.
	ExitRemoveLastVector(c *RemoveLastVectorContext)

	// ExitRemoveAtVector is called when exiting the RemoveAtVector production.
	ExitRemoveAtVector(c *RemoveAtVectorContext)

	// ExitIsEmptyVector is called when exiting the IsEmptyVector production.
	ExitIsEmptyVector(c *IsEmptyVectorContext)

	// ExitCountVector is called when exiting the CountVector production.
	ExitCountVector(c *CountVectorContext)

	// ExitAccessVector is called when exiting the AccessVector production.
	ExitAccessVector(c *AccessVectorContext)

	// ExitSelfFunction is called when exiting the SelfFunction production.
	ExitSelfFunction(c *SelfFunctionContext)

	// ExitStructAttribute is called when exiting the StructAttribute production.
	ExitStructAttribute(c *StructAttributeContext)

	// ExitStructFunction is called when exiting the StructFunction production.
	ExitStructFunction(c *StructFunctionContext)

	// ExitEmbbededFunc is called when exiting the embbededFunc production.
	ExitEmbbededFunc(c *EmbbededFuncContext)

	// ExitPrintstmt is called when exiting the printstmt production.
	ExitPrintstmt(c *PrintstmtContext)

	// ExitExprList is called when exiting the exprList production.
	ExitExprList(c *ExprListContext)

	// ExitIntstmt is called when exiting the intstmt production.
	ExitIntstmt(c *IntstmtContext)

	// ExitFloatstmt is called when exiting the floatstmt production.
	ExitFloatstmt(c *FloatstmtContext)

	// ExitStringstmt is called when exiting the stringstmt production.
	ExitStringstmt(c *StringstmtContext)

	// ExitStringExpr is called when exiting the StringExpr production.
	ExitStringExpr(c *StringExprContext)

	// ExitEmbeddedFunctionExpr is called when exiting the EmbeddedFunctionExpr production.
	ExitEmbeddedFunctionExpr(c *EmbeddedFunctionExprContext)

	// ExitNilExpr is called when exiting the NilExpr production.
	ExitNilExpr(c *NilExprContext)

	// ExitIdExpr is called when exiting the IdExpr production.
	ExitIdExpr(c *IdExprContext)

	// ExitCallBackExpr is called when exiting the CallBackExpr production.
	ExitCallBackExpr(c *CallBackExprContext)

	// ExitLogicalOperationExpr is called when exiting the LogicalOperationExpr production.
	ExitLogicalOperationExpr(c *LogicalOperationExprContext)

	// ExitNegExpr is called when exiting the NegExpr production.
	ExitNegExpr(c *NegExprContext)

	// ExitComparationOperationExpr is called when exiting the ComparationOperationExpr production.
	ExitComparationOperationExpr(c *ComparationOperationExprContext)

	// ExitArithmeticOperationExpr is called when exiting the ArithmeticOperationExpr production.
	ExitArithmeticOperationExpr(c *ArithmeticOperationExprContext)

	// ExitStructCallExpr is called when exiting the StructCallExpr production.
	ExitStructCallExpr(c *StructCallExprContext)

	// ExitRelationalOperationExpr is called when exiting the RelationalOperationExpr production.
	ExitRelationalOperationExpr(c *RelationalOperationExprContext)

	// ExitDigitExpr is called when exiting the DigitExpr production.
	ExitDigitExpr(c *DigitExprContext)

	// ExitNotExpr is called when exiting the NotExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitParenExpr is called when exiting the ParenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitCallFunctionExpr is called when exiting the CallFunctionExpr production.
	ExitCallFunctionExpr(c *CallFunctionExprContext)

	// ExitBooleanExpr is called when exiting the BooleanExpr production.
	ExitBooleanExpr(c *BooleanExprContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)
}
