package cst

import (
    "kigo-som/parser"
    //"log"
    //"reflect"
)

// method:
//    pattern Equal ( Primitive | methodBlock );

type MethodBlock struct {
  primitive bool
  block_contents *BlockContents
}

type Method struct{
	calling string
  method_block *MethodBlock
}

func MakeBlockOrPrimitive(ctx *parser.MethodContext) (*MethodBlock) {

  if primitive := ctx.Primitive() ; primitive != nil {
//    log.Println("Primitive" , primitive.GetText() , reflect.TypeOf(primitive))
    return &MethodBlock{ true , nil }
  }
  method_block := ctx.MethodBlock()
  block_contents := MakeBlockContents( method_block.(*parser.MethodBlockContext) )
  return &MethodBlock{ false , block_contents }
}

func MakeMethodFromContext(ctx *parser.MethodContext) (*Method)  {
	calling := MakeSelector(ctx)
  method_block := MakeBlockOrPrimitive(ctx)
  return &Method{ calling , method_block}
}
