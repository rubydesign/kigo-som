package cst

import (
    "github.com/antlr4-go/antlr/v4"
    "kigo-som/parser"
    "fmt"
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
  class_variables  []string
}

func MakeInstances( ctx *parser.InstanceFieldsContext ) ([]string) {
  instances := make([]string, 0, 3)
  variables  := ctx.AllVariable()
  for value := range instances {
    inst  := variables[value].(*parser.VariableContext)
		name := inst.GetText()
    instances = append(instances , name)
  }
  return instances
}

func MakeMethods( ctx *parser.ClassdefContext ) ([]*Method){
	var methods_ctx []parser.IMethodContext = ctx.AllMethod()
  methods := make([]*Method, 0, 3)
	fmt.Println("methods" , len(methods_ctx))
	for value := range methods_ctx {
		method_ctx  := methods_ctx[value].(*parser.MethodContext)
		method := MakeMethod(method_ctx)
		methods = append(methods , method)
  }
  return methods
}

func MakeClassfields(ctx *parser.ClassFieldsContext) ([]string){
  class_vars := make([]string, 0, 3)
  variables  := ctx.AllVariable()
  for value := range class_vars {
    inst  := variables[value].(*parser.VariableContext)
    name := inst.GetText()
    class_vars = append(class_vars , name)
  }
  return class_vars

}

func  MakeClassdef(ctx *parser.ClassdefContext) (*Classdef )  {
  name := ctx.Identifier().GetText()
  superclazz := ctx.Superclass().Identifier()
  super_name := ""
  if superclazz != nil { super_name = superclazz.GetText() }
  instances := MakeInstances(ctx.InstanceFields().( *parser.InstanceFieldsContext ) )
  methods := MakeMethods(ctx)
  class_vars := make([]string , 0 ,3)
  if classfieds_ctx := ctx.ClassFields() ; classfieds_ctx != nil {
    class_vars = MakeClassfields(classfieds_ctx.(*parser.ClassFieldsContext))
  }
  return  &Classdef{name , super_name ,  instances , methods , class_vars}
}

func ClassdefFromFile(file_name string)(*Classdef){
  input, _ := antlr.NewFileStream(file_name)
  lexer := parser.NewSomLexer(input)
	stream := antlr.NewCommonTokenStream(lexer,0)
  p := parser.NewSomParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
  ctx := p.Classdef().(*parser.ClassdefContext)
  return MakeClassdef(ctx)
}
