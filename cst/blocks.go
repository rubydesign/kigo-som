package cst

import (
    "kigo-som/parser"
    "log"
    "reflect"
)
//
// blockContents:
//     ( Or localDefs Or )?
//     blockBody;
//
// localDefs:
//     variable*;
//
// blockBody:
//       Exit result
//     | expression ( Period blockBody? )?;
//
// nestedBlock:
//     NewBlock blockPattern? blockContents? EndBlock;
//
// blockPattern:
//     blockArguments Or;
//
// blockArguments:
//     ( Colon argument )+;
//
// argument:
//     variable;
//
// variable:
//     identifier;
//
// identifier:
//     Primitive | Identifier;

type NestedBlock struct {
  pattern  []string
  contents *BlockContents
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

func MakeBlockPattern(pattern_ctx *parser.BlockPatternContext) ([]string) {
  pattern := make( []string , 0 , 3)
  arguments_ctx := pattern_ctx.BlockArguments()
  for i := range arguments_ctx.AllArgument() {
    argument_ctx := arguments_ctx.Argument(i)
    variable_ctx := argument_ctx.Variable()
    variable := variable_ctx.GetText()
    log.Println("variable" , variable , reflect.TypeOf(variable))
    pattern = append(pattern , variable)
  }
  return pattern
}

func MakeNestedBlock(ctx *parser.NestedBlockContext) (*NestedBlock){
  var pattern []string
  var contents *BlockContents
  if pattern_ctx := ctx.BlockPattern() ; pattern_ctx != nil {
    pattern = MakeBlockPattern(pattern_ctx.(*parser.BlockPatternContext))
  }
  if contents_ctx := ctx.BlockContents() ; contents_ctx != nil {
    contents = MakeBlockContents(contents_ctx.(*parser.BlockContentsContext))
  }
  return &NestedBlock{ pattern , contents}
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

func MakeBlockContents(ctx *parser.BlockContentsContext) (*BlockContents) {
  locals_ctx := ctx.LocalDefs()
  locals := MakeLocals(locals_ctx)
  block_ctx := ctx.BlockBody()
  block_body := MakeBlockBody(block_ctx.(*parser.BlockBodyContext))
  return &BlockContents{ locals , block_body }
}

func MakeBlockBody(block_ctx *parser.BlockBodyContext) (*BlockBody) {
  result_ctx := block_ctx.Result()
  var result_exp *Expression
  if result_ctx != nil {
    expression_ctx := result_ctx.Expression()
    result_exp = MakeExpression( expression_ctx.(*parser.ExpressionContext) )
    //log.Println("result_exp" , result_exp , reflect.TypeOf(result_exp))
    return &BlockBody{ result_exp , nil , nil}
  } else {
    expression_ctx := block_ctx.Expression()
    main_expression := MakeExpression( expression_ctx.(*parser.ExpressionContext) )
    var inner_block *BlockBody = nil
    inner_block_ctx := block_ctx.BlockBody()
    //log.Println("inner_block_ctx" , inner_block_ctx , reflect.TypeOf(inner_block_ctx))
    if inner_block_ctx != nil {
      inner_block = MakeBlockBody( inner_block_ctx.(*parser.BlockBodyContext) )
    }
    return &BlockBody{ nil , main_expression , inner_block}
  }
}
