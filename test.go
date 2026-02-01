package main

import (
	"github.com/antlr4-go/antlr/v4"
  "kigo-som/parser"
	"os"
)

type Visitor struct {
  *parser.BaseSomVisitor
}

type SomVisitor struct {
    parser.SomVisitor
}

func (v *SomVisitor) VisitClassdef(ctx *parser.ClassdefContext) interface{} {

    return ctx
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
  lexer := parser.NewSomLexer(input)
	stream := antlr.NewCommonTokenStream(lexer,0)
  p := parser.NewSomParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))


  tree := p.Classdef()

  // To use:
  visitor := &SomVisitor{}
  result := visitor.Visit(tree).(int)
  result = result - 1
  return
}
