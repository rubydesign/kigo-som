package main

import (
	"github.com/antlr4-go/antlr/v4"
  "kigo-som/parser"
	"os"
  "log"
)

type Classdef struct{
  Name  string
//  Super string
}

func  MakeClassdef(ctx *parser.ClassdefContext) (*Classdef )  {
  //superclazz := ctx.Superclass()
  name := ctx.Identifier().GetText()
  log.Println("Classdef Make: " , name)
  clazz := &Classdef{name}
  return clazz
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
  lexer := parser.NewSomLexer(input)
	stream := antlr.NewCommonTokenStream(lexer,0)
  p := parser.NewSomParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

  ret := MakeClassdef(p.Classdef().(*parser.ClassdefContext))

  log.Println("return = " , ret)
  return
}
