package interpreter

import (
	"fmt"
	"log"
	"server/parserInterpreter/parser"
)

type SymbolTable struct {
	Id           string
	TypeSymbol   string
	TypeVariable string
	TypeData     string
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

// create the visitor struct based on the SymbolTable struct
type Visitor struct {
	parser.BaseSFGrammarVisitor
	memory      map[string]SymbolTable
	symbolStack []map[string]SymbolTable
	Outputs     []string
	Errors      []Error
}

func (v *Visitor) pushScope() {
	v.symbolStack = append(v.symbolStack, make(map[string]SymbolTable))
}

func (v *Visitor) popScope() {
	if len(v.symbolStack) > 1 {
		v.symbolStack = v.symbolStack[:len(v.symbolStack)-1]
	}
}

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
	log.Fatalf("Error: Variable '%s' not declared", varName)
	// add the error to the errors
	v.Errors = append(v.Errors, Error{
		Line:   0,
		Column: 0,
		Msg:    fmt.Sprintf("Error: Variable '%s' not declared", varName),
		Type:   "Semantic",
	})
	return nil, false
}

// VerifyVariableScope verify if the variable is already declared in the scope
func (v *Visitor) VerifyVariableScope(varName string) bool {

	for i := len(v.symbolStack) - 1; i >= 0; i-- {
		scope := v.symbolStack[i]
		if _, ok := scope[varName]; ok {
			return true
		}
	}
	return false
}
