package interpreter

import (
	"fmt"
	"server/parserInterpreter/parser"
)

type SymbolTable struct {
	Id           string // the name of the variable
	TypeSymbol   string // the type of the symbol variable or function
	TypeVariable string // the type of the variable -> var or let
	TypeData     string // the type of the data -> Int, Float, String, Boolean, Character
	Value        interface{}
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

// create the visitor struct based on the SymbolTable struct
type Visitor struct {
	parser.BaseSFGrammarVisitor
	symbolStack []map[string]SymbolTable
	Outputs     []string
	Errors      []Error
	// manage loop context
	loopContexts []LoopContext
}

func (v *Visitor) pushScope() {
	v.symbolStack = append(v.symbolStack, make(map[string]SymbolTable))
}

func (v *Visitor) popScope() {
	if len(v.symbolStack) > 1 {
		v.symbolStack = v.symbolStack[:len(v.symbolStack)-1]
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
	return v.symbolStack[len(v.symbolStack)-1]
}

// udpate a variable in the scope
func (v *Visitor) UpdateVariable(varName string, value interface{}) {
	for i := len(v.symbolStack) - 1; i >= 0; i-- {
		scope := v.symbolStack[i]
		if _, ok := scope[varName]; ok {
			scope[varName] = value.(SymbolTable)
			return
		}
	}
}

// VerifyScope verify if the variable is in the scope
func (v *Visitor) VerifyScope(varName string) (interface{}, bool) {

	for i := len(v.symbolStack) - 1; i >= 0; i-- {
		scope := v.symbolStack[i]
		if val, ok := scope[varName]; ok {
			return val, true
		}
	}
	return nil, false
}

// VerifyVariableScope verify if the variable is already declared in the current scope, if not navigate to the others scopes
func (v *Visitor) VerifyVariableScope(varName string) bool {

	for i := len(v.symbolStack) - 1; i >= 0; i-- {
		scope := v.symbolStack[i]
		if _, ok := scope[varName]; ok {
			return true
		}
	}
	return false
}

// VerifyVariableCurrentScope verify if the variable is already declared in the current scope
func (v *Visitor) VerifyVariableCurrentScope(varName string) bool {

	scope := v.symbolStack[len(v.symbolStack)-1]
	fmt.Println("Current scope ->", scope)
	if _, ok := scope[varName]; ok {
		return true
	}
	return false
}
