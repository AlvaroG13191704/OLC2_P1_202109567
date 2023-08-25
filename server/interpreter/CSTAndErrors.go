package interpreter

import "github.com/antlr4-go/antlr/v4"

type NewCustomErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []Error
}

func (c *NewCustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.Errors = append(c.Errors, Error{Line: line, Column: column, Msg: msg, Type: "SyntaxError"})
}
