package main

import (
	"github.com/antlr4-go/antlr/v4"
  "kigo-som/parser"
	"os"
  "log"
  "reflect"
)

type Classdef struct{
  Name  string
  Super string
  Instances []string
}

func AddInstanceNames( ctx *parser.ClassdefContext , clazz *Classdef){
  instances  := ctx.InstanceFields().AllVariable()
  for value := range instances {
    inst  := instances[value].(*parser.VariableContext)
    log.Println("Name + type" , inst , reflect.TypeOf(inst))
    name := inst.GetText()
    clazz.Instances = append(clazz.Instances , name)
  }
}
func  ClassdefFromContext(ctx *parser.ClassdefContext) (*Classdef )  {
  name := ctx.Identifier().GetText()
  superclazz := ctx.Superclass().Identifier()
  super_name := ""
  if superclazz != nil { super_name = superclazz.GetText() }
  clazz := &Classdef{name , super_name ,  make([]string, 0, 3) }
  AddInstanceNames(ctx , clazz)
  return clazz
}

func ClassdefFromFile(file_name string)(*Classdef){
  input, _ := antlr.NewFileStream(file_name)
  lexer := parser.NewSomLexer(input)
	stream := antlr.NewCommonTokenStream(lexer,0)
  p := parser.NewSomParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
  ctx := p.Classdef().(*parser.ClassdefContext)
  return ClassdefFromContext(ctx)
}

func main() {
  ClassdefFromFile(os.Args[1])
  return
}
