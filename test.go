package main

import (
	"github.com/antlr4-go/antlr/v4"
  "kigo-som/parser"
	"os"
	"log"
	"reflect"
)

type Calling struct{
	name string
}

type Method struct{
	pattern string
	calling *Calling
}

type Classdef struct{
	name  string
	super string
	instance_variables []string
	instance_methods   []*Method
}

func MakeUnarySelector(ctx *parser.UnaryPatternContext) (*Calling) {
	name := ctx.UnarySelector().Identifier().GetText()
//	log.Println("Unary" , name , reflect.TypeOf(name))
	return &Calling{name}
}

func MakeBinarySelector(ctx *parser.BinaryPatternContext) (*Calling) {
	sel := ctx.BinarySelector()

	var binary antlr.TerminalNode
	var name string
	// seems i am missing something, silly way to do a switch, or just store the darn op
	if binary = sel.Or(); binary != nil {
		name = binary.GetText()
	} else if binary = sel.Comma(); binary != nil {
		name = binary.GetText()
	} else if binary = sel.Minus(); binary != nil {
		name = binary.GetText()
	} else if binary = sel.Equal(); binary != nil {
		name = binary.GetText()
	} else if binary = sel.Not(); binary != nil {
		name = binary.GetText()
	} else if binary = sel.And(); binary != nil {
		name = binary.GetText()
	} else if binary = sel.Star(); binary != nil {
		name = binary.GetText()
	} else if binary = sel.Div(); binary != nil {
		name = binary.GetText()
	} else if binary = sel.Mod(); binary != nil {
		name = binary.GetText()
	} else if binary = sel.Plus(); binary != nil {
		name = binary.GetText()
	} else if binary = sel.More(); binary != nil {
		name = binary.GetText()
	} else if binary = sel.Less(); binary != nil {
		name = binary.GetText()
	} else if binary = sel.At(); binary != nil {
		name = binary.GetText()
	} else if binary = sel.Per(); binary != nil {
		name = binary.GetText()
	} else if binary = sel.OperatorSequence(); binary != nil {
		panic( "double op not impemented")
		name = binary.GetText()
	} else {
		panic("Unknown operator type")
	}
//log.Println("Context Unary" , unary , reflect.TypeOf(unary))
	return &Calling{name}
}

func MakeKeywordSelector(ctx *parser.KeywordPatternContext) (*Calling) {
	all := ""
	for keyword_idx := range ctx.AllKeyword() {
		keyword_ctx := ctx.Keyword(keyword_idx)
		name := keyword_ctx.Keyword().GetText()
		all += name
	}
	//log.Println("Ctx" , all , reflect.TypeOf(all))
	return &Calling{all}
}

func MakeMethodFromContext(ctx *parser.MethodContext) (*Method)  {
	pattern := ctx.Pattern()
	if unary := pattern.UnaryPattern() ; unary != nil {
		MakeUnarySelector(unary.(*parser.UnaryPatternContext))
	} else if binary := pattern.BinaryPattern() ; binary != nil {
		MakeBinarySelector(binary.(*parser.BinaryPatternContext))
	} else if keyword := pattern.KeywordPattern() ; keyword != nil {
		//log.Println("Context Keyword" , keyword, reflect.TypeOf(keyword))
		MakeKeywordSelector(keyword.(*parser.KeywordPatternContext))
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
		clazz.instance_methods = append(clazz.instance_methods , method)
  }
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

func main() {
  ClassdefFromFile(os.Args[1])
  return
}
