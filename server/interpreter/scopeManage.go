package interpreter

import (
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

type SymbolTable struct {
	Id           string // the name of the variable
	TypeSymbol   string // the type of the symbol variable or function, vector, reference, struct
	TypeVariable string // the type of the variable -> var or let, struct
	TypeData     string // the type of the data -> Int, Float, String, Boolean, Character
	StructOf     string // the name of the struct that contains the variable
	Value        interface{}
	ListParams   interface{}
	Mutating     bool
	Line         int
	Column       int
}

type Error struct {
	Line   int
	Column int
	Msg    string
	Type   string
}

type LoopContext struct {
	TypeLoop      string
	ContinueFound bool
	BreakFound    bool
}

type SelfStruct struct {
	VarId        string
	StructOf     string
	FunctionName string
	FunctionCall bool
}

// create the visitor struct based on the SymbolTable struct
type Visitor struct {
	parser.BaseSFGrammarVisitor
	SymbolStack []map[string]SymbolTable
	TableSymbol []SymbolTable
	Outputs     []string
	Errors      []Error
	// manage loop context
	loopContexts    []LoopContext
	ReturnValue     interface{}
	IsReturn        bool
	FirstPass       bool
	FunctionContext []string
	// manage father structs to use it with the self keyword
	SelfStructs map[string]SelfStruct
}

func (v *Visitor) pushScope() {
	v.SymbolStack = append(v.SymbolStack, make(map[string]SymbolTable))
}

func (v *Visitor) popScope() {
	if len(v.SymbolStack) > 1 {
		v.SymbolStack = v.SymbolStack[:len(v.SymbolStack)-1]
	}
}

// PushLoopContext push a loop context
func (v *Visitor) PushLoopContext(typeLoop string) {
	v.loopContexts = append(v.loopContexts, LoopContext{TypeLoop: typeLoop, ContinueFound: false, BreakFound: false})
}

// PopLoopContext pop a loop context
func (v *Visitor) PopLoopContext() {
	if len(v.loopContexts) > 0 {
		v.loopContexts = v.loopContexts[:len(v.loopContexts)-1]
	}
}

// ExistsLoopContext check if a loop context exists
func (v *Visitor) ExistsLoopContext() bool {
	return len(v.loopContexts) > 0
}

// update the current loop context
func (v *Visitor) UpdateLoopContext(ctx LoopContext) {
	v.loopContexts[len(v.loopContexts)-1] = ctx
}

// GetLoopContext get the current loop context
func (v *Visitor) GetLoopContext() LoopContext {

	return v.loopContexts[len(v.loopContexts)-1]
}

// AssignVariable assign a variable to the current scope -> works for loops
func (v *Visitor) AssignVariable(varName string, value interface{}) {
	scope := v.getCurrentScope()
	scope[varName] = value.(SymbolTable)
}

// DeleteVariable delete a variable from the current scope -> works for loops
func (v *Visitor) DeleteVariable(varName string) {
	scope := v.getCurrentScope()
	delete(scope, varName)
}

// MANAGE SYMBOL TABLE
func (v *Visitor) getCurrentScope() map[string]SymbolTable {
	return v.SymbolStack[len(v.SymbolStack)-1]
}

// VerifySelfStruct
func (v *Visitor) VerifySelfStruct(attributeName string) (interface{}, bool, string) {
	// iterate over the self stack and find which is activate
	for _, i := range v.SelfStructs {
		if i.FunctionCall {
			// find the variable in the scope
			variable, ok := v.VerifyScope(i.VarId)
			if !ok {
				return nil, false, ""
			}
			// verify if the variable is the same type of the struct
			if variable.(SymbolTable).StructOf == i.StructOf {
				// find the attribute in the scope of the struct
				structScope := variable.(SymbolTable).Value.(map[string]SymbolTable)
				_, ok := v.VerifyStructScopeValue(structScope, attributeName)

				if !ok {
					return nil, false, ""
				}
				// return variable
				return variable.(SymbolTable), true, i.FunctionName
			}
		}
	}
	return nil, false, ""
}

// ActiveFunctionInSelfStack
func (v *Visitor) ActiveFunctionInSelfStack(idStruct string, functionName string) {
	// find the struct in the self stack
	for a, i := range v.SelfStructs {
		if a == idStruct {
			v.SelfStructs[a] = SelfStruct{VarId: i.VarId, StructOf: i.StructOf, FunctionCall: true, FunctionName: functionName}
		}
	}
}

// DeactiveFunctionInSelfStack
func (v *Visitor) DeactiveFunctionInSelfStack(idStruct string) {
	// find the struct in the self stack
	for a, i := range v.SelfStructs {
		if a == idStruct {
			v.SelfStructs[a] = SelfStruct{VarId: i.VarId, StructOf: i.StructOf, FunctionCall: false, FunctionName: ""}
		}
	}
}

// udpate a variable in the scope
func (v *Visitor) UpdateVariable(varName string, value interface{}) {
	for i := len(v.SymbolStack) - 1; i >= 0; i-- {
		scope := v.SymbolStack[i]
		if _, ok := scope[varName]; ok {
			scope[varName] = value.(SymbolTable)
			return
		}
	}
}

// VerifyStructScope verify if the id is in the scope of the struct
func (v *Visitor) VerifyStructScopeValue(scope map[string]SymbolTable, varName string) (interface{}, bool) {
	if val, ok := scope[varName]; ok {
		return val, true
	}
	return nil, false
}

// VerifyScope verify if the variable is in the scope
func (v *Visitor) VerifyScope(varName string) (interface{}, bool) {

	for i := len(v.SymbolStack) - 1; i >= 0; i-- {
		scope := v.SymbolStack[i]
		if val, ok := scope[varName]; ok {
			return val, true
		}
	}
	return nil, false
}

// verifySelf verify if the attribute is in the scope of the struct from the self

// VerifyVariableScope verify if the variable is already declared in the current scope, if not navigate to the others scopes
func (v *Visitor) VerifyVariableScope(varName string) bool {

	for i := len(v.SymbolStack) - 1; i >= 0; i-- {
		scope := v.SymbolStack[i]
		if _, ok := scope[varName]; ok {
			return true
		}
	}
	return false
}

// VerifyVariableCurrentScope verify if the variable is already declared in the current scope
func (v *Visitor) VerifyVariableCurrentScope(varName string) bool {

	scope := v.SymbolStack[len(v.SymbolStack)-1]
	// fmt.Println("Current scope ->", scope)
	if variable, ok := scope[varName]; ok {
		// evaluate if the variable is not a function
		return variable.TypeSymbol == values.Type_Variable

	}
	return false
}

// VerifyFunctionScope verify if the function is already declared in the current scope
func (v *Visitor) VerifyFunctionScope(varName string) bool {

	scope := v.SymbolStack[len(v.SymbolStack)-1]
	if variable, ok := scope[varName]; ok {
		// evaluate if the variable is a function
		return variable.TypeSymbol == values.Type_Function

	}
	return false
}

// GetFunction from the scope
func (v *Visitor) GetFunction(varName string) (SymbolTable, bool) {
	for i := len(v.SymbolStack) - 1; i >= 0; i-- {
		scope := v.SymbolStack[i]
		if val, ok := scope[varName]; ok {
			if val.TypeSymbol == values.Type_Function {
				return val, true
			}
		}
	}
	return SymbolTable{}, false
}

// ExistsFunctionContext check if a function context exists
func (v *Visitor) ExistsFunctionContext() bool {
	return len(v.FunctionContext) > 0
}

// PushFunctionContext push a function context
func (v *Visitor) PushFunctionContext(functionContext string) {
	v.FunctionContext = append(v.FunctionContext, functionContext)
}

// PopFunctionContext pop a function context
func (v *Visitor) PopFunctionContext() {
	if len(v.FunctionContext) > 0 {
		v.FunctionContext = v.FunctionContext[:len(v.FunctionContext)-1]
	}
}
