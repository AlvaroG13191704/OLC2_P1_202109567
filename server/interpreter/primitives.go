package interpreter

import (
	"fmt"
	"server/parserInterpreter/parser"
	"strconv"
)

// visit digitexpr
func (v *Visitor) VisitDigitExpr(ctx *parser.DigitExprContext) interface{} {
	i, _ := strconv.ParseInt(ctx.GetText(), 10, 64) // convert the string to int
	fmt.Println("Nativo: ", i)
	return i // return the int
}
