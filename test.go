package main

import (
	"github.com/antlr4-go/antlr/v4"
  "kigo-som/parser"
	"os"
  "log"
)

type SomVisitor struct {
    parser.BaseSomVisitor
}

func (v *SomVisitor) VisitClassdef(ctx *parser.ClassdefContext) interface{} {
  log.Println("Classdef visited")
  return v.VisitChildren(ctx)
}
func (v *SomVisitor) VisitSuperclass(ctx *parser.SuperclassContext) interface{} {
  log.Println("Superclass visited")
  return v.VisitChildren(ctx)
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

  tree.Accept(visitor)

  return
}
