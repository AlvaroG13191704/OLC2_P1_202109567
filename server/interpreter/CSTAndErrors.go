package interpreter

import (
	"github.com/antlr4-go/antlr/v4"
)

type NewCustomErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []Error
}

func (c *NewCustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.Errors = append(c.Errors, Error{Line: line, Column: column, Msg: msg, Type: "SyntaxError"})
}

// var Counter int = 0
// var Dot string = "graph CST {" + "\n"

// type Node struct {
// 	Id    int
// 	Label string
// 	Left  *Node
// 	Right *Node
// }

// type Tree struct {
// 	Root *Node
// }

// func (t Tree) Postorder(tmp *Node) {
// 	if tmp != nil {
// 		t.Postorder(tmp.Left)
// 		t.Postorder(tmp.Right)
// 		// fmt.Print(tmp.Label, " ")
// 		Dot += strconv.Itoa(tmp.Id) + "[label=\"" + tmp.Label + "\"];\n"
// 		if tmp.Left != nil {
// 			Dot += strconv.Itoa(tmp.Id) + "--" + strconv.Itoa(tmp.Left.Id) + ";\n"
// 		}
// 		if tmp.Right != nil {
// 			Dot += strconv.Itoa(tmp.Id) + "--" + strconv.Itoa(tmp.Right.Id) + ";\n"
// 		}
// 	}
// }
