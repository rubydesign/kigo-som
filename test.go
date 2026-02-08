package main

import (
	"github.com/antlr4-go/antlr/v4"
  "kigo-som/parser"
	"os"
	"log"
	"reflect"
)

type Method struct{
	Pattern string
}

type Classdef struct{
  Name  string
  Super string
  Instance_variables []string
	Instance_methods   []*Method
}

func MakeMethodFromContext(ctx *parser.MethodContext) (*Method)  {
	pattern := ctx.Pattern()
	if unary := pattern.UnaryPattern() ; unary != nil {
		log.Println("Unary" , unary , reflect.TypeOf(unary))
	} else if keyword := pattern.KeywordPattern() ; keyword != nil {
		log.Println("Keyword" , keyword, reflect.TypeOf(keyword))
	} else if binary := pattern.BinaryPattern() ; binary != nil {
		log.Println("Binary" , binary , reflect.TypeOf(binary))
	} else {
		log.Println("Unknown type" , pattern , reflect.TypeOf(pattern))
		panic(1)
	}

	return nil
}

func AddMethods( ctx *parser.ClassdefContext , clazz *Classdef){
	var methods_ctx []parser.IMethodContext = ctx.AllMethod()
	log.Println("No of methods" , len(methods_ctx))
	for value := range methods_ctx {
		method_ctx  := methods_ctx[value].(*parser.MethodContext)
		method := MakeMethodFromContext(method_ctx)
		clazz.Instance_methods = append(clazz.Instance_methods , method)
  }
}

func AddInstanceNames( ctx *parser.ClassdefContext , clazz *Classdef){
  instances  := ctx.InstanceFields().AllVariable()
  for value := range instances {
    inst  := instances[value].(*parser.VariableContext)
    name := inst.GetText()
    clazz.Instance_variables = append(clazz.Instance_variables , name)
  }
}

func  ClassdefFromContext(ctx *parser.ClassdefContext) (*Classdef )  {
  name := ctx.Identifier().GetText()
  superclazz := ctx.Superclass().Identifier()
  super_name := ""
  if superclazz != nil { super_name = superclazz.GetText() }
	clazz := &Classdef{name , super_name ,  make([]string, 0, 3) , make([]*Method, 0, 3)}
  AddInstanceNames(ctx , clazz)
	AddMethods(ctx , clazz)
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
