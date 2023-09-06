package interpreter

import (
	"fmt"
	"log"
	"reflect"
	"server/parserInterpreter/interpreter/values"
	"server/parserInterpreter/parser"
)

// VisitAccessVectorStruct
func (v *Visitor) VisitAccessVectorStruct(ctx *parser.AccessVectorStructContext) interface{} {
	// get the vector name
	vectorName := ctx.AllID_PRIMITIVE()[0].GetText()

	// evaluate if the vector exists
	value, ok := v.VerifyScope(vectorName)
	if !ok {
		log.Printf("Error: Variable '%s' not declared \n", vectorName)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Variable '%s' not declared", vectorName),
			Type:   "Semantic",
		})
		return nil
	}
	symbolValue := value.(SymbolTable)

	// get the index
	index := v.Visit(ctx.Expr()).(values.PRIMITIVE)
	// evaluate if the index is a number
	if index.GetType() != values.IntType {
		log.Printf("Error: Index '%s' is not a number \n", index.GetValue())
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Index '%s' is not a number", index.GetValue()),
			Type:   "Semantic",
		})
		return nil
	}
	// convert the index from int64 to int
	numINdex := int(index.GetValue().(int64))

	// evaluate if the index is out of bounds
	if numINdex < 0 || numINdex >= len(symbolValue.Value.([]map[string]SymbolTable)) {
		log.Printf("Error: Index '%d' is out of bounds \n", index.GetValue())
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Index '%d' is out of bounds", index.GetValue()),
			Type:   "Semantic",
		})
		return nil
	}

	// get the struct value
	structValue := symbolValue.Value.([]map[string]SymbolTable)[numINdex]

	return structValue[ctx.AllID_PRIMITIVE()[1].GetText()].Value.(values.PRIMITIVE)

}

// VisitAccessVector
func (v *Visitor) VisitAccessVector(ctx *parser.AccessVectorContext) interface{} {
	// get the vector name
	vectorName := ctx.ID_PRIMITIVE().GetText()

	// evaluate if the vector exists
	value, ok := v.VerifyScope(vectorName)
	if !ok {
		log.Printf("Error: Variable '%s' not declared \n", vectorName)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Variable '%s' not declared", vectorName),
			Type:   "Semantic",
		})
		return nil
	}
	symbolValue := value.(SymbolTable)

	// get the index
	index := v.Visit(ctx.Expr()).(values.PRIMITIVE)

	// evaluate if the index is a number
	if index.GetType() != values.IntType {
		log.Printf("Error: Index '%s' is not a number \n", index.GetValue())
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Index '%s' is not a number", index.GetValue()),
			Type:   "Semantic",
		})
		return nil
	}

	// convert the index from int64 to int
	numINdex := int(index.GetValue().(int64))

	// evaluate if the index is out of bounds
	if numINdex < 0 || numINdex >= len(symbolValue.Value.([]values.PRIMITIVE)) {
		log.Printf("Error: Index '%d' is out of bounds \n", index.GetValue())
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Index '%d' is out of bounds", index.GetValue()),
			Type:   "Semantic",
		})
		return nil
	}

	return symbolValue.Value.([]values.PRIMITIVE)[numINdex]
}

// VisitAppendVector -> append an expression to an existing vector
func (v *Visitor) VisitAppendVector(ctx *parser.AppendVectorContext) interface{} {
	// get the vector name
	vectorName := ctx.ID_PRIMITIVE().GetText()

	// evaluate if the vector exists
	value, ok := v.VerifyScope(vectorName)
	if !ok {
		log.Printf("Error: Variable '%s' not declared \n", vectorName)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Variable '%s' not declared", vectorName),
			Type:   "Semantic",
		})
		return nil
	}
	symbolValue := value.(SymbolTable)

	assert := v.Visit(ctx.Expr())

	// assert
	if reflect.TypeOf(assert).Kind() == reflect.Map {
		fmt.Println("assert is a map")
		expr := assert.(map[string]SymbolTable)

		fmt.Println("expr", expr)
		// append
		symbolValue.Value = append(symbolValue.Value.([]map[string]SymbolTable), expr)
		// update the symbol table
		v.UpdateVariable(vectorName, symbolValue)
		return nil
	}
	// get the expression
	expr := assert.(values.PRIMITIVE)

	// evaluate if the expr type is the same as the symbol value
	if symbolValue.TypeData != expr.GetType() {
		log.Printf("Error: Type '%s' is not assignable to type '%s' \n", expr.GetType(), symbolValue.TypeData)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Type '%s' is not assignable to type '%s'", expr.GetType(), symbolValue.TypeData),
			Type:   "Semantic",
		})
		return nil
	}

	// append the value to the vector
	symbolValue.Value = append(symbolValue.Value.([]values.PRIMITIVE), expr)

	// update the symbol table
	v.UpdateVariable(vectorName, symbolValue)

	return nil
}

// VisitRemoveLastVector
func (v *Visitor) VisitRemoveLastVector(ctx *parser.RemoveLastVectorContext) interface{} {
	// get the vector name
	vectorName := ctx.ID_PRIMITIVE().GetText()

	// evaluate if the vector exists
	value, ok := v.VerifyScope(vectorName)
	if !ok {
		log.Printf("Error: Variable '%s' not declared \n", vectorName)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Variable '%s' not declared", vectorName),
			Type:   "Semantic",
		})
		return nil
	}
	symbolValue := value.(SymbolTable)

	// evaluate if the vector is empty
	if len(symbolValue.Value.([]values.PRIMITIVE)) == 0 {
		log.Printf("Error: Vector '%s' is empty \n", vectorName)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Vector '%s' is empty", vectorName),
			Type:   "Semantic",
		})
		return nil
	}

	fmt.Println("before removeLast", symbolValue.Value)
	// remove the last element
	symbolValue.Value = symbolValue.Value.([]values.PRIMITIVE)[:len(symbolValue.Value.([]values.PRIMITIVE))-1]

	fmt.Println("after removeLast", symbolValue.Value)
	// update the symbol table
	v.UpdateVariable(vectorName, symbolValue)

	return nil
}

// VisitRemoveAtVector
func (v *Visitor) VisitRemoveAtVector(ctx *parser.RemoveAtVectorContext) interface{} {
	// get the vector name
	vectorName := ctx.ID_PRIMITIVE().GetText()

	// evaluate if the vector exists
	value, ok := v.VerifyScope(vectorName)
	if !ok {
		log.Printf("Error: Variable '%s' not declared \n", vectorName)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Variable '%s' not declared", vectorName),
			Type:   "Semantic",
		})
		return nil
	}
	symbolValue := value.(SymbolTable)

	// get the index
	index := v.Visit(ctx.Expr()).(values.PRIMITIVE)

	// evaluate if the index is a number
	if index.GetType() != values.IntType {
		log.Printf("Error: Index '%s' is not a number \n", index.GetValue())
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Index '%s' is not a number", index.GetValue()),
			Type:   "Semantic",
		})
		return nil
	}

	// convert the index from int64 to int
	numINdex := int(index.GetValue().(int64))

	// evaluate if the index is out of bounds
	if numINdex < 0 || numINdex >= len(symbolValue.Value.([]values.PRIMITIVE)) {
		log.Printf("Error: Index '%d' is out of bounds \n", index.GetValue())
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Index '%d' is out of bounds", index.GetValue()),
			Type:   "Semantic",
		})
		return nil
	}

	fmt.Println("before removeAt", symbolValue.Value)
	// remove the element at the index
	symbolValue.Value = append(symbolValue.Value.([]values.PRIMITIVE)[:index.GetValue().(int64)], symbolValue.Value.([]values.PRIMITIVE)[index.GetValue().(int64)+1:]...)

	fmt.Println("after removeAt", symbolValue.Value)
	// update the symbol table
	v.UpdateVariable(vectorName, symbolValue)

	return nil

}

// VisitIsEmptyVector
func (v *Visitor) VisitIsEmptyVector(ctx *parser.IsEmptyVectorContext) interface{} {
	// get the vector name
	vectorName := ctx.ID_PRIMITIVE().GetText()

	// evaluate if the vector exists
	value, ok := v.VerifyScope(vectorName)
	if !ok {
		log.Printf("Error: Variable '%s' not declared \n", vectorName)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Variable '%s' not declared", vectorName),
			Type:   "Semantic",
		})
		return nil
	}
	symbolValue := value.(SymbolTable)

	// evaluate if the vector is empty
	if len(symbolValue.Value.([]values.PRIMITIVE)) == 0 {
		return &values.Boolean{Value: true}
	}

	return &values.Boolean{Value: false}
}

// VisitCountVector
func (v *Visitor) VisitCountVector(ctx *parser.CountVectorContext) interface{} {
	// get the vector name
	vectorName := ctx.ID_PRIMITIVE().GetText()

	// evaluate if the vector exists
	value, ok := v.VerifyScope(vectorName)
	if !ok {
		log.Printf("Error: Variable '%s' not declared \n", vectorName)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Variable '%s' not declared", vectorName),
			Type:   "Semantic",
		})
		return nil
	}
	symbolValue := value.(SymbolTable)

	return &values.Integer{Value: int64(len(symbolValue.Value.([]values.PRIMITIVE)))}
}
