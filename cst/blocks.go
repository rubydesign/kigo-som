package cst

import (
    "kigo-som/parser"
    "fmt"
    "strings"
    // "reflect"
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
  Locals []string
  block_body *BlockBody
}

func PrintNestedBlock(pre string ,  block *NestedBlock ){
  if len( block.pattern ) > 0 {
    fmt.Println(pre , "NestedBlock pattern:" , strings.Join(block.pattern , " ") )
  }
  PrintBlockContents(  pre , block.contents)
}

func PrintBlockBody( pre string , block_body *BlockBody)  {
  fmt.Println(pre , "BlockBody:" )
  if block_body.result != nil {
    fmt.Println("|  " + pre , "return:" )
    PrintExpression("|  " + pre , block_body.result)
  } else {
    PrintExpression("|  " + pre,block_body.main)
    if block_body.code != nil { PrintBlockBody("|  " + pre , block_body.code) }
  }
}

func PrintBlockContents( pre string , block_contents *BlockContents)  {
  if len( block_contents.Locals ) > 0 {
    fmt.Println(pre , "BlockContents" , "locals:" , strings.Join(block_contents.Locals , " ") )
  }
  PrintBlockBody( pre ,  block_contents.block_body)
}

func MakeBlockPattern(pattern_ctx *parser.BlockPatternContext) ([]string) {
  pattern := make( []string , 0 , 3)
  block_ctx := pattern_ctx.BlockArguments()
  arguments_ctx := block_ctx.AllArgument()
  for _, argument_ctx := range arguments_ctx {
    variable_ctx := argument_ctx.Variable()
    variable := variable_ctx.GetText()
    //fmt.Println("block variable" , variable , reflect.TypeOf(variable))
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
    variables_ctx := ctx.AllVariable()
    for _, variable := range variables_ctx  {
      locals = append( locals , variable.GetText() )
//      fmt.Println("variable" , variable , reflect.TypeOf(variable))
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
    //fmt.Println("result_exp" , result_exp , reflect.TypeOf(result_exp))
    return &BlockBody{ result_exp , nil , nil}
  } else {
    expression_ctx := block_ctx.Expression()
    main_expression := MakeExpression( expression_ctx.(*parser.ExpressionContext) )
    var inner_block *BlockBody = nil
    inner_block_ctx := block_ctx.BlockBody()
    //fmt.Println("inner_block_ctx" , inner_block_ctx , reflect.TypeOf(inner_block_ctx))
    if inner_block_ctx != nil {
      inner_block = MakeBlockBody( inner_block_ctx.(*parser.BlockBodyContext) )
    }
    return &BlockBody{ nil , main_expression , inner_block}
  }
}
