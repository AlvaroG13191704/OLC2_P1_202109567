package routes

import (
	"fmt"
	"server/parserInterpreter/interpreter"
	"server/parserInterpreter/parser"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gofiber/fiber/v2"
)

func AnalyzeAndParseCode() fiber.Handler {
	return func(c *fiber.Ctx) error {

		// get the code from the request body
		code := string(c.Body())
		fmt.Println(code)

		input := antlr.NewInputStream(code)                                    // convert string to stream
		lexer := parser.NewSFGrammarLexer(input)                               // create lexer
		tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel) // create tokens
		p := parser.NewSFGrammarParser(tokens)                                 // create parser

		p.BuildParseTrees = true // tell the parser to build parse trees
		tree := p.Start_()       // parse the input

		// create the visitor
		visitor := interpreter.NewVisitor()
		visitor.Visit(tree) // visit the tree

		output := visitor.Outputs

		// join the output array and separate by new line
		out := ""
		for _, v := range output {
			out += v + "\n"
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Code parsed successfully",
			"result":  out,
		})
	}
}
