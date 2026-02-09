package cst

import (
    "kigo-som/parser"
    "log"
    "reflect"
)

type Expression struct {
  exp string
}

type BlockBody struct {
  exit *Expression
}


type BlockContents struct {
  locals []string
  block_body *BlockBody
}

type MethodBlock struct {
  primitive bool
  block_contents *BlockContents
}

type Method struct{
	calling string
  method_block *MethodBlock
}


func MakeLocals(ctx parser.ILocalDefsContext) ([]string) {
  locals := make([]string , 0 , 3)
  if ctx != nil {
    for i := range ctx.AllVariable()  {
      variable := ctx.Variable(i).GetText()
      locals = append( locals , variable )
//      log.Println("variable" , variable , reflect.TypeOf(variable))
    }
  }
  return locals
}

func MakeExpression(ctx *parser.ExpressionContext) (*Expression) {
  return nil
}

func MakeBlockBody(ctx *parser.BlockContentsContext) (*BlockBody) {
  block_ctx := ctx.BlockBody()
  result_ctx := block_ctx.Result()
  var result_exp *Expression
  if result_ctx != nil {
    expression_ctx := result_ctx.Expression()
    result_exp = MakeExpression( expression_ctx.(*parser.ExpressionContext) )
    log.Println("result_exp" , result_exp , reflect.TypeOf(result_exp))
  }
  return &BlockBody{ result_exp }
}

func MakeBlockContents(ctx *parser.MethodBlockContext) (*BlockContents) {
  content_ctx := ctx.BlockContents()
  locals_ctx := content_ctx.LocalDefs()
  locals := MakeLocals(locals_ctx)
  block_body := MakeBlockBody(content_ctx.(*parser.BlockContentsContext))
  return &BlockContents{ locals , block_body }
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
