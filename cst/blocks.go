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
