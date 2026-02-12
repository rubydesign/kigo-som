package cst

import (
    "github.com/antlr4-go/antlr/v4"
    "kigo-som/parser"
    "log"
    "reflect"
)

// pattern:
//     unaryPattern | keywordPattern | binaryPattern;
//
// unaryPattern:
//     unarySelector;
//
// binaryPattern:
//     binarySelector argument;
//
// keywordPattern:
//     ( keyword argument )+;

// Selectors representet as strings for now

// Helper method for argument extraction
// Get the variable from an argument and addd the text of it to the slice
// return the same slice
func AddVariable(argument_ctx *parser.ArgumentContext, in []string) ([]string){
  variable_ctx := argument_ctx.Variable()
  variable := variable_ctx.GetText()
//  log.Println("variable" , variable , reflect.TypeOf(variable))
  return append(in , variable)
}

func MakeUnarySelector(ctx *parser.UnaryPatternContext) (string) {
	name := ctx.UnarySelector().Identifier().GetText()
//	log.Println("Unary" , name , reflect.TypeOf(name))
	return name
}

func MakeBinarySelector(ctx *parser.BinaryPatternContext) ([]string) {
	sel := ctx.BinarySelector()
  calling := make( [] string , 0 , 3)
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
    name = binary.GetText()
	} else {
		panic("Unknown operator type")
	}
  calling = append(calling , name)
  return AddVariable( ctx.Argument().(*parser.ArgumentContext) , calling)
}

func MakeKeywordSelector(ctx *parser.KeywordPatternContext) ([]string) {
  calling := make( [] string , 0 , 3)
	for idx := range ctx.AllKeyword() {
		keyword_ctx := ctx.Keyword(idx)
		name := keyword_ctx.Keyword().GetText()
    calling = append(calling , name )
    calling = AddVariable( ctx.Argument(idx).(*parser.ArgumentContext) , calling)
	}
  return calling
}

func MakeSelector(pattern_ctx *parser.PatternContext) ([]string) {
  if keyword := pattern_ctx.KeywordPattern() ; keyword != nil {
    //log.Println("Context Keyword" , keyword, reflect.TypeOf(keyword))
    return MakeKeywordSelector(keyword.(*parser.KeywordPatternContext))
  }
  if unary := pattern_ctx.UnaryPattern() ; unary != nil {
    calling := MakeUnarySelector(unary.(*parser.UnaryPatternContext))
    return []string{calling}
  } else if binary := pattern_ctx.BinaryPattern() ; binary != nil {
    return MakeBinarySelector(binary.(*parser.BinaryPatternContext))
  } else {
    log.Println("Unknown type" , pattern_ctx , reflect.TypeOf(pattern_ctx))
		panic(1)
	}
}
