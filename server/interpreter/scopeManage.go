package interpreter

import "server/parserInterpreter/parser"

type SymbolTable struct {
	Id         string
	TypeSymbol string
	TypeData   string
	Value      interface{}
	Line       int
	Column     int
}

// create the visitor struct based on the SymbolTable struct
type Visitor struct {
	parser.BaseSFGrammarVisitor
	memory      map[string]SymbolTable
	symbolStack []map[string]SymbolTable
	Outputs     []string
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

// type Visitor struct {
// 	parser.BaseSFGrammarVisitor
// 	memory      map[string]interface{}
// 	symbolStack []map[string]interface{}
// }

// func (v *Visitor) pushScope() {
// 	v.symbolStack = append(v.symbolStack, make(map[string]interface{}))
// }

// func (v *Visitor) popScope() {
// 	if len(v.symbolStack) > 1 {
// 		v.symbolStack = v.symbolStack[:len(v.symbolStack)-1]
// 	}
// }

// func (v *Visitor) getCurrentScope() map[string]interface{} {
// 	return v.symbolStack[len(v.symbolStack)-1]
// }
