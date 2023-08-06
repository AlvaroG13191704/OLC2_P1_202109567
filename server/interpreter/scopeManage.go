package interpreter

import "server/parserInterpreter/parser"

// create the visitor struct
type Visitor struct {
	parser.BaseSFGrammarVisitor
	Memory map[string]interface{}
	// symbolStack []map[string]interface{}
}
