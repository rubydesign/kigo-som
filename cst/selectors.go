package cst

import (
    "github.com/antlr4-go/antlr/v4"
    "kigo-som/parser"
    "log"
    "reflect"
)

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
