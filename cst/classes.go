package cst

import (
    "github.com/antlr4-go/antlr/v4"
    "kigo-som/parser"
    "log"
)

// classdef:
//     Identifier Equal superclass
//     instanceFields method*
//     ( Separator classFields method* )?
//     EndTerm;
//
// superclass:
//     Identifier? NewTerm;
//
// instanceFields:
//     ( Or variable* Or )?;
//
// classFields:
//     ( Or variable* Or )?;


type Classdef struct{
	name  string
	super string
	instance_variables []string
	instance_methods   []*Method
}

func AddInstanceNames( ctx *parser.ClassdefContext , clazz *Classdef){
  instances  := ctx.InstanceFields().AllVariable()
  for value := range instances {
		inst  := instances[value].(*parser.VariableContext)
		name := inst.GetText()
		clazz.instance_variables = append(clazz.instance_variables , name)
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

func AddMethods( ctx *parser.ClassdefContext , clazz *Classdef){
	var methods_ctx []parser.IMethodContext = ctx.AllMethod()
	log.Println("No of methods" , len(methods_ctx))
	for value := range methods_ctx {
		method_ctx  := methods_ctx[value].(*parser.MethodContext)
		method := MakeMethod(method_ctx)
		clazz.instance_methods = append(clazz.instance_methods , method)
  }
}
