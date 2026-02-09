package cst

import (
    "github.com/antlr4-go/antlr/v4"
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


func MakeUnarySelector(ctx *parser.UnaryPatternContext) (string) {
	name := ctx.UnarySelector().Identifier().GetText()
//	log.Println("Unary" , name , reflect.TypeOf(name))
	return name
}

func MakeBinarySelector(ctx *parser.BinaryPatternContext) (string) {
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
	return name
}

func MakeKeywordSelector(ctx *parser.KeywordPatternContext) (string) {
	all := ""
	for keyword_idx := range ctx.AllKeyword() {
		keyword_ctx := ctx.Keyword(keyword_idx)
		name := keyword_ctx.Keyword().GetText()
		all += name
	}
	//log.Println("Ctx" , all , reflect.TypeOf(all))
	return all
}

func MakeSelector(ctx *parser.MethodContext) (string) {
	calling := ""
	pattern := ctx.Pattern()
	if unary := pattern.UnaryPattern() ; unary != nil {
		calling = MakeUnarySelector(unary.(*parser.UnaryPatternContext))
	} else if binary := pattern.BinaryPattern() ; binary != nil {
		calling = MakeBinarySelector(binary.(*parser.BinaryPatternContext))
	} else if keyword := pattern.KeywordPattern() ; keyword != nil {
		//log.Println("Context Keyword" , keyword, reflect.TypeOf(keyword))
		calling = MakeKeywordSelector(keyword.(*parser.KeywordPatternContext))
	} else {
		log.Println("Unknown type" , pattern , reflect.TypeOf(pattern))
		panic(1)
	}
	return calling
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
