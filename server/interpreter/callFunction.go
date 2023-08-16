package interpreter

import (
	"fmt"
	"log"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

func (v *Visitor) VisitCallFunctionWithoutParams(ctx *parser.CallFunctionWithoutParamsContext) interface{} {
	// push a new function context
	v.PushFunctionContext("function")
	defer v.PopFunctionContext()

	functionName := ctx.ID_PRIMITIVE().GetText()

	// get function from the scope
	function, ok := v.GetFunction(functionName)
	if !ok {
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("function %s not found", functionName),
			Type:   "Semantic",
		})
		log.Printf("function %s not found", functionName)
		return nil
	}

	// verifiy if the return value type is the same as the function return type
	if function.TypeVariable == "void" {
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("function '%s' is void", functionName),
			Type:   "Semantic",
		})
		log.Printf("function '%s' is void", functionName)
		return nil
	} else {
		// execute the function
		v.Visit(function.Value.(*parser.BlockContext))

		// evaluate if the return type is the same as the function return type
		if function.TypeVariable != v.ReturnValue.(values.PRIMITIVE).GetType() {
			// add error
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("function %s return type is %s, expected %s", functionName, function.TypeVariable, v.ReturnValue.(values.PRIMITIVE).GetType()),
				Type:   "Semantic",
			})
			log.Printf("function %s return type is %s, expected %s", functionName, function.TypeVariable, v.ReturnValue.(values.PRIMITIVE).GetType())
			return nil
		}

		return nil
	}

}

// CallFunctionExpr
func (v *Visitor) VisitCallFunctionExpr(ctx *parser.CallFunctionExprContext) interface{} {
	v.Visit(ctx.CallFunctionStmt())
	// visit the expression
	fmt.Println("VisitCallFunctionExpr", v.ReturnValue)

	return v.ReturnValue
}

// VisitCallFunctionWithParamsEI
func (v *Visitor) VisitCallFunctionWithParamsEI(ctx *parser.CallFunctionWithParamsEIContext) interface{} {
	// push a new function context
	v.PushFunctionContext("function")
	defer v.PopFunctionContext()

	functionName := ctx.ID_PRIMITIVE().GetText()

	// get function from the scope
	function, ok := v.GetFunction(functionName)
	if !ok {
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("function %s not found", functionName),
			Type:   "Semantic",
		})
		log.Printf("function %s not found", functionName)
		return nil
	}

	// verifiy if the return value type is the same as the function return type
	if function.TypeVariable == "void" {
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("function '%s' is void", functionName),
			Type:   "Semantic",
		})
		log.Printf("function '%s' is void", functionName)
		return nil
	}

	// fmt.Println("function", function)
	// get list params
	listParams := function.ListParams.(map[string][]SymbolTable)
	// get list of arguments
	listArguments := v.Visit(ctx.ListCallFunctionStmt())

	// if both list has external and internal keys then validate
	if listParams["external"] != nil && listArguments.(map[string][]SymbolTable)["external"] != nil && listParams["internal"] != nil && listArguments.(map[string][]SymbolTable)["internal"] != nil {
		fmt.Println("listParams ->", listParams)
		fmt.Println("listArguments ->", listArguments)
		// validate external params with external arguments
		if len(listParams["external"]) != len(listArguments.(map[string][]SymbolTable)["external"]) {
			// add error
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("function %s expected %d arguments, got %d", functionName, len(listParams["external"]), len(listArguments.(map[string][]SymbolTable)["external"])),
				Type:   "Semantic",
			})
			log.Printf("function %s expected %d arguments, got %d", functionName, len(listParams["external"]), len(listArguments.(map[string][]SymbolTable)["external"]))
			return nil
		}

		// evaluate if the values are the same type in external params and external arguments
		for i, param := range listParams["external"] {
			// get the value of the expression
			value := listArguments.(map[string][]SymbolTable)["external"][i].Value.(values.PRIMITIVE)

			// evaluate if the values are the same type
			if param.TypeData != value.GetType() {
				// add error
				v.Errors = append(v.Errors, Error{
					Line:   ctx.GetStart().GetLine(),
					Column: ctx.GetStart().GetColumn(),
					Msg:    fmt.Sprintf("function %s expected %s type, got %s", functionName, param.TypeData, value.GetType()),
					Type:   "Semantic",
				})
				log.Printf("function %s expected %s type, got %s", functionName, param.TypeData, value.GetType())
				return nil
			}
		}

		// asign the values to the internal params
		for i, param := range listParams["internal"] {
			// get the value of the expression
			value := listArguments.(map[string][]SymbolTable)["external"][i].Value.(values.PRIMITIVE)

			// evaluate if the values are the same type
			if param.TypeData != value.GetType() {
				// add error
				v.Errors = append(v.Errors, Error{
					Line:   ctx.GetStart().GetLine(),
					Column: ctx.GetStart().GetColumn(),
					Msg:    fmt.Sprintf("function %s expected %s type, got %s", functionName, param.TypeData, value.GetType()),
					Type:   "Semantic",
				})
				log.Printf("function %s expected %s type, got %s", functionName, param.TypeData, value.GetType())
				return nil
			}

			// asign the value to the symbol table
			param.Value = value

			// asign the symbol table to the internal params
			listParams["internal"][i] = param
		}

		// create the internal values
		v.pushScope()
		// defer v.popScope()
		for _, param := range listParams["internal"] {
			v.getCurrentScope()[param.Id] = param
		}

		// // execute the function
		v.Visit(function.Value.(*parser.BlockContext))
		// fmt.Println("----------------------------------------------------")
		// fmt.Println("Current scope or symbol table ->", v.getCurrentScope())
		// fmt.Println("Global scope or symbol table ->", v.symbolStack)
		// fmt.Println("----------------------------------------------------")
		// evaluate if the return type is the same as the function return type
		if function.TypeVariable != v.ReturnValue.(values.PRIMITIVE).GetType() {
			// add error
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("function %s return type is %s, expected %s", functionName, function.TypeVariable, v.ReturnValue.(values.PRIMITIVE).GetType()),
				Type:   "Semantic",
			})
			log.Printf("function %s return type is %s, expected %s", functionName, function.TypeVariable, v.ReturnValue.(values.PRIMITIVE).GetType())
			return nil
		}

	} else if listParams["internal"] != nil && listArguments.(map[string][]SymbolTable)["internal"] != nil {
		fmt.Println("listParams ->", listParams)
		fmt.Println("listArguments ->", listArguments)
		// validate internal params with internal arguments
		if len(listParams["internal"]) != len(listArguments.(map[string][]SymbolTable)["internal"]) {
			// add error
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("function %s expected %d arguments, got %d", functionName, len(listParams["internal"]), len(listArguments.(map[string][]SymbolTable)["internal"])),
				Type:   "Semantic",
			})
			log.Printf("function %s expected %d arguments, got %d", functionName, len(listParams["internal"]), len(listArguments.(map[string][]SymbolTable)["internal"]))
			return nil
		}

		// evaluate if the values are the same type in internal params and internal arguments
		for i, param := range listParams["internal"] {
			// get the value of the expression
			value := listArguments.(map[string][]SymbolTable)["internal"][i].Value.(values.PRIMITIVE)

			// evaluate if the values are the same type
			if param.TypeData != value.GetType() {
				// add error
				v.Errors = append(v.Errors, Error{
					Line:   ctx.GetStart().GetLine(),
					Column: ctx.GetStart().GetColumn(),
					Msg:    fmt.Sprintf("function %s expected %s type, got %s", functionName, param.TypeData, value.GetType()),
					Type:   "Semantic",
				})
				log.Printf("function %s expected %s type, got %s", functionName, param.TypeData, value.GetType())
				return nil
			}

			// asign the value to the symbol table
			param.Value = value

			// asign the symbol table to the internal params
			listParams["internal"][i] = param
		}

		// create the internal values
		v.pushScope()
		// defer v.popScope()
		for _, param := range listParams["internal"] {
			v.getCurrentScope()[param.Id] = param
		}

		// // execute the function
		v.Visit(function.Value.(*parser.BlockContext))
		// fmt.Println("----------------------------------------------------")
		// fmt.Println("Current scope or symbol table ->", v.getCurrentScope())
		// fmt.Println("Global scope or symbol table ->", v.symbolStack)
		// fmt.Println("----------------------------------------------------")
		// evaluate if the return type is the same as the function return type
		if function.TypeVariable != v.ReturnValue.(values.PRIMITIVE).GetType() {
			// add error
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("function %s return type is %s, expected %s", functionName, function.TypeVariable, v.ReturnValue.(values.PRIMITIVE).GetType()),
				Type:   "Semantic",
			})
			log.Printf("function %s return type is %s, expected %s", functionName, function.TypeVariable, v.ReturnValue.(values.PRIMITIVE).GetType())
			return nil
		}

	} else if listParams["EI"] != nil && listArguments.(map[string][]SymbolTable)["external"] != nil {
		fmt.Println("listParams ->", listParams)
		fmt.Println("listArguments ->", listArguments)
		// validate external params with external arguments
		if len(listParams["EI"]) != len(listArguments.(map[string][]SymbolTable)["external"]) {
			// add error
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("function %s expected %d arguments, got %d", functionName, len(listParams["EI"]), len(listArguments.(map[string][]SymbolTable)["external"])),
				Type:   "Semantic",
			})
			log.Printf("function %s expected %d arguments, got %d", functionName, len(listParams["EI"]), len(listArguments.(map[string][]SymbolTable)["external"]))
			return nil
		}

		// evaluate if the values are the same type in external params and external arguments
		for i, param := range listParams["EI"] {
			// get the value of the expression
			value := listArguments.(map[string][]SymbolTable)["external"][i].Value.(values.PRIMITIVE)

			// evaluate if the values are the same type
			if param.TypeData != value.GetType() {
				// add error
				v.Errors = append(v.Errors, Error{
					Line:   ctx.GetStart().GetLine(),
					Column: ctx.GetStart().GetColumn(),
					Msg:    fmt.Sprintf("function %s expected %s type, got %s", functionName, param.TypeData, value.GetType()),
					Type:   "Semantic",
				})
				log.Printf("function %s expected %s type, got %s", functionName, param.TypeData, value.GetType())
				return nil
			}
		}

		// asign the values to the internal params
		for i, param := range listParams["EI"] {
			// get the value of the expression
			value := listArguments.(map[string][]SymbolTable)["external"][i].Value.(values.PRIMITIVE)

			// evaluate if the values are the same type
			if param.TypeData != value.GetType() {
				// add error
				v.Errors = append(v.Errors, Error{
					Line: ctx.GetStart().GetLine(),
					Column: ctx.GetStart().
						GetColumn(),
					Msg:  fmt.Sprintf("function %s expected %s type, got %s", functionName, param.TypeData, value.GetType()),
					Type: "Semantic",
				})
				log.Printf("function %s expected %s type, got %s", functionName, param.TypeData, value.GetType())
				return nil
			}

			// asign the value to the symbol table
			param.Value = value

			// asign the symbol table to the internal params
			listParams["EI"][i] = param

		}

		// create the internal values
		v.pushScope()
		// defer v.popScope()
		for _, param := range listParams["EI"] {
			v.getCurrentScope()[param.Id] = param
		}

		// // execute the function
		v.Visit(function.Value.(*parser.BlockContext))
		// fmt.Println("----------------------------------------------------")
		// fmt.Println("Current scope or symbol table ->", v.getCurrentScope())
		// fmt.Println("Global scope or symbol table ->", v.symbolStack)
		// fmt.Println("----------------------------------------------------")

		// evaluate if the return type is the same as the function return type
		if function.TypeVariable != v.ReturnValue.(values.PRIMITIVE).GetType() {
			// add error
			v.Errors = append(v.Errors, Error{
				Line: ctx.GetStart().GetLine(),
				Column: ctx.GetStart().
					GetColumn(),
				Msg:  fmt.Sprintf("function %s return type is %s, expected %s", functionName, function.TypeVariable, v.ReturnValue.(values.PRIMITIVE).GetType()),
				Type: "Semantic",
			})
			log.Printf("function %s return type is %s, expected %s", functionName, function.TypeVariable, v.ReturnValue.(values.PRIMITIVE).GetType())
			return nil
		}

	}

	return nil
}

// listCallFunctionStmtEI
func (v *Visitor) VisitListCallFunctionStmtEI(ctx *parser.ListCallFunctionStmtEIContext) interface{} {
	// create a map to save the params
	params := make(map[string][]SymbolTable)

	// create two keys, external and internal
	params["external"] = []SymbolTable{}

	listIds := ctx.AllID_PRIMITIVE()
	listExpr := ctx.AllExpr()

	// iterate over the list of ids and save the values
	for i, id := range listIds {
		// get the value of the expression
		value := v.Visit(listExpr[i]).(values.PRIMITIVE)

		// get the id
		idValue := id.GetText()

		// create a symbol table
		symbolTable := SymbolTable{
			Id:           idValue,
			TypeSymbol:   values.Type_Variable,
			TypeVariable: "let",
			TypeData:     value.GetType(),
			Value:        value,
			ListParams:   nil,
			Line:         ctx.GetStart().GetLine(),
			Column:       ctx.GetStart().GetColumn(),
		}

		// append the symbol table to the params
		params["external"] = append(params["external"], symbolTable)
	}

	return params
}

// VisitListCallFunctionStmtNEI
func (v *Visitor) VisitListCallFunctionStmtNEI(ctx *parser.ListCallFunctionStmtNEIContext) interface{} {
	// create a map to save the params
	params := make(map[string][]SymbolTable)

	// create two keys, external and internal
	params["internal"] = []SymbolTable{}

	listExpr := ctx.AllExpr()

	// iterate over the list of ids and save the values
	for _, expr := range listExpr {
		// get the value of the expression
		value := v.Visit(expr).(values.PRIMITIVE)

		// create a symbol table
		symbolTable := SymbolTable{
			Id:           expr.GetText(),
			TypeSymbol:   values.Type_Variable,
			TypeVariable: "let",
			TypeData:     value.GetType(),
			Value:        value,
			ListParams:   nil,
			Line:         ctx.GetStart().GetLine(),
			Column:       ctx.GetStart().GetColumn(),
		}

		// append the symbol table to the params
		params["internal"] = append(params["internal"], symbolTable)
	}

	return params
}
