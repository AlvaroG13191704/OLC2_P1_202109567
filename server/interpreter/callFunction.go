package interpreter

import (
	"fmt"
	"log"
	"reflect"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

// CallFunctionExpr
func (v *Visitor) VisitCallFunctionExpr(ctx *parser.CallFunctionExprContext) interface{} {
	v.Visit(ctx.CallFunctionStmt())
	// visit the expression
	// change the return value
	v.IsReturn = false
	fmt.Println("VisitCallFunctionExpr", v.ReturnValue)

	return v.ReturnValue
}

// VisitCallFunctionWithoutParams
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

// VisitCallFunctionWithParamsEI
func (v *Visitor) VisitCallFunctionWithParams(ctx *parser.CallFunctionWithParamsContext) interface{} {
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

	// get list params
	listParams := function.ListParams.(map[string][]SymbolTable)
	// get list of arguments
	listArguments := v.Visit(ctx.ListCallFunctionStmt())

	// if both list has external and internal keys then validate
	if listParams["external"] != nil && listArguments.(map[string][]SymbolTable)["external"] != nil && listParams["internal"] != nil && listArguments.(map[string][]SymbolTable)["internal"] != nil {
		fmt.Println("params with external and internal --> _ value")
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
			// value := listArguments.(map[string][]SymbolTable)["external"][i].Value.(values.PRIMITIVE)
			value := listArguments.(map[string][]SymbolTable)["external"][i].Value

			// verifiy if the value is a slice
			if reflect.TypeOf(value).Kind() == reflect.Slice {
				value = listArguments.(map[string][]SymbolTable)["external"][i].Value.([]values.PRIMITIVE)
				typeValue := listArguments.(map[string][]SymbolTable)["external"][i].TypeData
				// evaluate if the values are the same type
				if param.TypeData != typeValue {
					// add error
					v.Errors = append(v.Errors, Error{
						Line:   ctx.GetStart().GetLine(),
						Column: ctx.GetStart().GetColumn(),
						Msg:    fmt.Sprintf("function %s expected %s type, got %s", functionName, param.TypeData, typeValue),
						Type:   "Semantic",
					})
					log.Printf("function %s expected %s type, got %s", functionName, param.TypeData, typeValue)
					return nil
				}

			} else {
				valueP := listArguments.(map[string][]SymbolTable)["external"][i].Value.(values.PRIMITIVE)
				// evaluate if the values are the same type
				if param.TypeData != valueP.GetType() {
					// add error
					v.Errors = append(v.Errors, Error{
						Line:   ctx.GetStart().GetLine(),
						Column: ctx.GetStart().GetColumn(),
						Msg:    fmt.Sprintf("function %s expected %s type, got %s", functionName, param.TypeData, valueP.GetType()),
						Type:   "Semantic",
					})
					log.Printf("function %s expected %s type, got %s", functionName, param.TypeData, valueP.GetType())
					return nil
				}
			}

		}

		// asign the values to the internal params
		for i, param := range listParams["internal"] {
			// get the value of the expression
			value := listArguments.(map[string][]SymbolTable)["external"][i].Value

			// asign the value to the symbol table
			param.Value = value

			// asign the symbol table to the internal params
			listParams["internal"][i] = param
		}

	} else if listParams["internal"] != nil && listArguments.(map[string][]SymbolTable)["internal"] != nil {
		fmt.Println("params with only internal, not external --> _ value")
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
			value := listArguments.(map[string][]SymbolTable)["internal"][i].Value

			// verifiy if the value is a slice
			if reflect.TypeOf(value).Kind() == reflect.Slice {
				value = listArguments.(map[string][]SymbolTable)["internal"][i].Value.([]values.PRIMITIVE)
				typeValue := listArguments.(map[string][]SymbolTable)["internal"][i].TypeData
				// evaluate if the values are the same type
				if param.TypeData != typeValue {
					// add error
					v.Errors = append(v.Errors, Error{
						Line:   ctx.GetStart().GetLine(),
						Column: ctx.GetStart().GetColumn(),
						Msg:    fmt.Sprintf("function %s expected %s type, got %s", functionName, param.TypeData, typeValue),
						Type:   "Semantic",
					})
					log.Printf("function %s expected %s type, got %s", functionName, param.TypeData, typeValue)
					return nil
				}

			} else {
				value = listArguments.(map[string][]SymbolTable)["internal"][i].Value.(values.PRIMITIVE)
				valueP := listArguments.(map[string][]SymbolTable)["internal"][i].Value.(values.PRIMITIVE)
				// evaluate if the values are the same type
				if param.TypeData != valueP.GetType() {
					// add error
					v.Errors = append(v.Errors, Error{
						Line:   ctx.GetStart().GetLine(),
						Column: ctx.GetStart().GetColumn(),
						Msg:    fmt.Sprintf("function %s expected %s type, got %s", functionName, param.TypeData, valueP.GetType()),
						Type:   "Semantic",
					})
					log.Printf("function %s expected %s type, got %s", functionName, param.TypeData, valueP.GetType())
					return nil
				}
			}

			// asign the value to the symbol table
			param.Value = value

			// asign the symbol table to the internal params
			listParams["internal"][i] = param
		}

	} else if listParams["internal"] != nil && listArguments.(map[string][]SymbolTable)["external"] != nil {
		fmt.Println("params with both internal and external, same name")
		fmt.Println("listParams ->", listParams)
		fmt.Println("listArguments ->", listArguments)
		// validate external params with external arguments
		if len(listParams["internal"]) != len(listArguments.(map[string][]SymbolTable)["external"]) {
			// add error
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("function %s expected %d arguments, got %d", functionName, len(listParams["internal"]), len(listArguments.(map[string][]SymbolTable)["external"])),
				Type:   "Semantic",
			})
			log.Printf("function %s expected %d arguments, got %d", functionName, len(listParams["internal"]), len(listArguments.(map[string][]SymbolTable)["external"]))
			return nil
		}

		// evaluate if the values are the same type in external params and external arguments
		for i, param := range listParams["internal"] {
			// get the value of the expression
			value := listArguments.(map[string][]SymbolTable)["external"][i].Value

			// verifiy if the value is a slice
			if reflect.TypeOf(value).Kind() == reflect.Slice {
				value = listArguments.(map[string][]SymbolTable)["external"][i].Value.([]values.PRIMITIVE)
				typeValue := listArguments.(map[string][]SymbolTable)["external"][i].TypeData
				// evaluate if the values are the same type
				if param.TypeData != typeValue {
					// add error
					v.Errors = append(v.Errors, Error{
						Line:   ctx.GetStart().GetLine(),
						Column: ctx.GetStart().GetColumn(),
						Msg:    fmt.Sprintf("function %s expected %s type, got %s", functionName, param.TypeData, typeValue),
						Type:   "Semantic",
					})
					log.Printf("function %s expected %s type, got %s", functionName, param.TypeData, typeValue)
					return nil
				}
			} else {
				valueP := listArguments.(map[string][]SymbolTable)["external"][i].Value.(values.PRIMITIVE)
				// evaluate if the values are the same type
				if param.TypeData != valueP.GetType() {
					// add error
					v.Errors = append(v.Errors, Error{
						Line:   ctx.GetStart().GetLine(),
						Column: ctx.GetStart().GetColumn(),
						Msg:    fmt.Sprintf("function %s expected %s type, got %s", functionName, param.TypeData, valueP.GetType()),
						Type:   "Semantic",
					})
					log.Printf("function %s expected %s type, got %s", functionName, param.TypeData, valueP.GetType())
					return nil
				}
			}

		}

		// asign the values to the internal params
		for i, param := range listParams["internal"] {
			// get the value of the expression
			value := listArguments.(map[string][]SymbolTable)["external"][i].Value

			// asign the value to the symbol table
			param.Value = value

			// asign the symbol table to the internal params
			listParams["internal"][i] = param

		}
	}

	// // execute the function
	v.pushScope()
	defer v.popScope()
	for _, param := range listParams["internal"] {
		v.getCurrentScope()[param.Id] = param
	}
	// create the internal values
	v.Visit(function.Value.(*parser.BlockContext))

	if v.ReturnValue != nil {
		if function.TypeVariable != v.ReturnValue.(values.PRIMITIVE).GetType() {
			// add error
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("function %s return type is %s, expected %s", functionName, function.TypeVariable, v.ReturnValue.(values.PRIMITIVE).GetType()),
				Type:   "Semantic",
			})
			log.Printf("function %s return type is %s, expected %s", functionName, function.TypeVariable, v.ReturnValue.(values.PRIMITIVE).GetType())
			return values.Nil{}
		}
	}

	// fmt.Println("----------------------------------------------------")
	// fmt.Println("Current scope or symbol table ->", v.getCurrentScope())
	// fmt.Println("Global scope or symbol table ->", v.symbolStack)
	// fmt.Println("----------------------------------------------------")

	return values.Nil{}
}

// listCallFunctionStmtEI
func (v *Visitor) VisitListCallFunctionStmtEI(ctx *parser.ListCallFunctionStmtEIContext) interface{} {
	// create a map to save the params
	params := make(map[string][]SymbolTable)

	// create two keys, external and internal
	params["external"] = []SymbolTable{}

	listIds := ctx.AllID_PRIMITIVE()
	listExpr := ctx.AllExpr()

	for i, id := range listIds {
		value := v.Visit(listExpr[i])

		// get the id
		idValue := id.GetText()

		if reflect.TypeOf(value).Kind() == reflect.Slice {
			fmt.Println("is a slice")
			// create a symbol table
			primitives := value.([]values.PRIMITIVE)
			symbolTable := SymbolTable{
				Id:           idValue,
				TypeSymbol:   values.Type_Vector,
				TypeVariable: "let",
				TypeData:     primitives[0].GetType(),
				Value:        primitives,
				ListParams:   nil,
				Line:         ctx.GetStart().GetLine(),
				Column:       ctx.GetStart().GetColumn(),
			}

			// append the symbol table to the params
			params["external"] = append(params["external"], symbolTable)

		} else {
			primitiveValue := value.(values.PRIMITIVE)
			// create a symbol table
			symbolTable := SymbolTable{
				Id:           idValue,
				TypeSymbol:   values.Type_Variable,
				TypeVariable: "let",
				TypeData:     primitiveValue.GetType(),
				Value:        primitiveValue,
				ListParams:   nil,
				Line:         ctx.GetStart().GetLine(),
				Column:       ctx.GetStart().GetColumn(),
			}

			// append the symbol table to the params
			params["external"] = append(params["external"], symbolTable)
		}
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
		value := v.Visit(expr)

		// create a symbol table
		if reflect.TypeOf(value).Kind() == reflect.Slice {
			fmt.Println("is a slice")
			// create a symbol table
			primitives := value.([]values.PRIMITIVE)
			symbolTable := SymbolTable{
				Id:           expr.GetText(),
				TypeSymbol:   values.Type_Vector,
				TypeVariable: "let",
				TypeData:     primitives[0].GetType(),
				Value:        primitives,
				ListParams:   nil,
				Line:         ctx.GetStart().GetLine(),
				Column:       ctx.GetStart().GetColumn(),
			}

			// append the symbol table to the params
			params["internal"] = append(params["internal"], symbolTable)

		} else {
			primitiveValue := value.(values.PRIMITIVE)
			// create a symbol table
			symbolTable := SymbolTable{
				Id:           expr.GetText(),
				TypeSymbol:   values.Type_Variable,
				TypeVariable: "let",
				TypeData:     primitiveValue.GetType(),
				Value:        primitiveValue,
				ListParams:   nil,
				Line:         ctx.GetStart().GetLine(),
				Column:       ctx.GetStart().GetColumn(),
			}
			// append the symbol table to the params
			params["internal"] = append(params["internal"], symbolTable)
		}

	}

	return params
}
