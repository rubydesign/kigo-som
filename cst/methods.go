package cst

import (
    "kigo-som/parser"
    "fmt"
    "strings"
    // "reflect"
)

// method:
//    pattern Equal ( Primitive | methodBlock );
// methodBlock:
//     NewTerm blockContents? EndTerm;

type MethodBlock struct {
  primitive bool  // OR MethodBlock, which is basically BlockContents
  block_contents *BlockContents
}

type Method struct{
  pattern []string
  method_block *MethodBlock
}

func PrintMethodBlock( pre string , method_block *MethodBlock)  {
  fmt.Println(pre ,  "MethodBlock:" )
  if method_block.primitive {
    fmt.Println( "|  " + pre,  "primitive")
  } else{
    PrintBlockContents(  "|  " + pre ,  method_block.block_contents)
  }
}

func PrintMethod( pre string , method *Method)  {
  fmt.Println(pre , "Method:(" , strings.Join(method.pattern , " ") , ")")
  PrintMethodBlock(pre , method.method_block)
}

func PrintMethods(pre  string , methods []*Method){
  fmt.Println(pre , "Methods:" , len(methods))
  for i := range methods {
    method := methods[i]
    PrintMethod(pre , method)
  }
}

func MakeBlockOrPrimitive(ctx *parser.MethodContext) (*MethodBlock) {
  if primitive := ctx.Primitive() ; primitive != nil {
//    fmt.Println("Primitive" , primitive.GetText() , reflect.TypeOf(primitive))
    return &MethodBlock{ true , nil }
  }
  method_block_ctx := ctx.MethodBlock()
  var block_contents *BlockContents
  if contents_ctx := method_block_ctx.BlockContents() ; contents_ctx != nil {
    block_contents = MakeBlockContents( contents_ctx.(*parser.BlockContentsContext) )
  }
  return &MethodBlock{ false , block_contents }
}

func MakeMethod(ctx *parser.MethodContext) (*Method)  {
  pattern_ctx := ctx.Pattern()
  pattern := MakePattern(pattern_ctx.(*parser.PatternContext))
  //fmt.Println("pattern" , pattern , reflect.TypeOf(pattern))
  method_block := MakeBlockOrPrimitive(ctx)
  return &Method{ pattern , method_block}
}
