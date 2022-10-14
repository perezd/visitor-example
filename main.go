package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/perezd/visitor-example/parser"
)

type jsonVisitorImpl struct {
	parser.BaseJSONVisitor
}

func (v *jsonVisitorImpl) VisitValue(ctx *parser.ValueContext) interface{} {
	fmt.Println("HERE")
	v.VisitChildren(ctx)
	return nil
}

func newJSONVisitor() parser.JSONVisitor {
	return new(jsonVisitorImpl)
}

func main() {
	input, err := antlr.NewFileStream("./input.json")
	if err != nil {
		panic(err)
	}

	lexer := parser.NewJSONLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewJSONParser(stream)
	parser.BuildParseTrees = true

	visitor := newJSONVisitor()
	visitor.Visit(parser.Json())
}
