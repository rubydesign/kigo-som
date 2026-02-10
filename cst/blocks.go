package cst

import (
    "kigo-som/parser"
    "log"
    "reflect"
)

type Expression struct {
  exp string
}

// either result / return   OR
// main with optional code
// ie a (possibly empty) list of expression followed by a return
type BlockBody struct {
  result *Expression
  main *Expression
  code *BlockBody
}

type BlockContents struct {
  locals []string
  block_body *BlockBody
}

func MakeExpression(ctx *parser.ExpressionContext) (*Expression) {
  return nil
}

func MakeBlockBody(block_ctx *parser.BlockBodyContext) (*BlockBody) {
  result_ctx := block_ctx.Result()
  var result_exp *Expression
  if result_ctx != nil {
    expression_ctx := result_ctx.Expression()
    result_exp = MakeExpression( expression_ctx.(*parser.ExpressionContext) )
    log.Println("result_exp" , result_exp , reflect.TypeOf(result_exp))
    return &BlockBody{ result_exp , nil , nil}
  } else {
    expression_ctx := block_ctx.Expression()
    main_expression := MakeExpression( expression_ctx.(*parser.ExpressionContext) )

    return &BlockBody{ nil , main_expression , nil}
  }
}

func MakeBlockContents(ctx *parser.MethodBlockContext) (*BlockContents) {
  content_ctx := ctx.BlockContents()
  locals_ctx := content_ctx.LocalDefs()
  locals := MakeLocals(locals_ctx)
  block_ctx := content_ctx.BlockBody()
  block_body := MakeBlockBody(block_ctx.(*parser.BlockBodyContext))
  return &BlockContents{ locals , block_body }
}
