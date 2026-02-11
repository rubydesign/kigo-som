package cst

import (
    "kigo-som/parser"
    "log"
    "reflect"
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

func MakeBlockOrPrimitive(ctx *parser.MethodContext) (*MethodBlock) {
  if primitive := ctx.Primitive() ; primitive != nil {
//    log.Println("Primitive" , primitive.GetText() , reflect.TypeOf(primitive))
    return &MethodBlock{ true , nil }
  }
  method_block_ctx := ctx.MethodBlock()
  contents_ctx := method_block_ctx.BlockContents()
  block_contents := MakeBlockContents( contents_ctx.(*parser.BlockContentsContext) )
  return &MethodBlock{ false , block_contents }
}

func MakeMethod(ctx *parser.MethodContext) (*Method)  {
  pattern_ctx := ctx.Pattern()
  pattern := MakeSelector(pattern_ctx.(*parser.PatternContext))
  log.Println("pattern" , pattern , reflect.TypeOf(pattern))

  method_block := MakeBlockOrPrimitive(ctx)
  return &Method{ pattern , method_block}
}
