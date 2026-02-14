package cst

import (
    "github.com/antlr4-go/antlr/v4"
    "kigo-som/parser"
    "fmt"
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
//
// unarySelector:
//     identifier;
//
// binarySelector:
//     Or | Comma | Minus | Equal | Not | And | Star | Div | Mod | Plus | More |
//     Less | At | Per | OperatorSequence;

// Selectors represented as strings for now

// Helper method for argument extraction
// Get the variable from an argument and addd the text of it to the slice
// return the same slice
func AddVariable(argument_ctx *parser.ArgumentContext, in []string) ([]string){
  variable_ctx := argument_ctx.Variable()
  variable := variable_ctx.GetText()
//  fmt.Println("variable" , variable , reflect.TypeOf(variable))
  return append(in , variable)
}

func MakeUnaryPattern(ctx *parser.UnaryPatternContext) (string) {
	name := ctx.UnarySelector().Identifier().GetText()
//	fmt.Println("Unary" , name , reflect.TypeOf(name))
	return name
}

func MakeBinarySelector(sel *parser.BinarySelectorContext) (string) {
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
  return name
}

func MakeKeywordPattern(ctx *parser.KeywordPatternContext) ([]string) {
  calling := make( [] string , 0 , 3)
  keywords_ctx := ctx.AllKeyword()
  for idx := range keywords_ctx {
    keyword_ctx := keywords_ctx[idx]
		name := keyword_ctx.Keyword().GetText()
    calling = append(calling , name )
    calling = AddVariable( ctx.Argument(idx).(*parser.ArgumentContext) , calling)
	}
  return calling
}

func MakePattern(pattern_ctx *parser.PatternContext) ([]string) {
  if keyword := pattern_ctx.KeywordPattern() ; keyword != nil {
    //fmt.Println("Context Keyword" , keyword, reflect.TypeOf(keyword))
    return MakeKeywordPattern(keyword.(*parser.KeywordPatternContext))
  }
  if unary := pattern_ctx.UnaryPattern() ; unary != nil {
    calling := MakeUnaryPattern(unary.(*parser.UnaryPatternContext))
    return []string{calling}
  } else if binary := pattern_ctx.BinaryPattern() ; binary != nil {
    selector := binary.BinarySelector()
    name := MakeBinarySelector(selector.(*parser.BinarySelectorContext))
    calling := make([]string , 0 , 3)
    calling = append(calling , name)
    return  AddVariable( binary.Argument().(*parser.ArgumentContext) , calling)
  } else {
    fmt.Println("Unknown type" , pattern_ctx , reflect.TypeOf(pattern_ctx))
		panic(1)
	}
}
func MakeKeywordSelector(ctx *parser.KeywordSelectorContext)(string){
  if keyword := ctx.Keyword() ; keyword != nil {
    return keyword.GetText()
  }
  sequence := ctx.KeywordSequence()
  name := sequence.GetText()
  fmt.Println("KEYWORD sequence" , name , reflect.TypeOf(name))
  return name
}
func MakeSelector(selector_ctx *parser.SelectorContext) (string) {
  if keyword := selector_ctx.KeywordSelector() ; keyword != nil {
    return MakeKeywordSelector(keyword.(*parser.KeywordSelectorContext))
  }
  if unary := selector_ctx.UnarySelector() ; unary != nil {
    name := unary.Identifier().GetText()
    return name
  } else if binary := selector_ctx.BinarySelector() ; binary != nil {
    return MakeBinarySelector(binary.(*parser.BinarySelectorContext))
  } else {
    fmt.Println("Unknown type" , selector_ctx , reflect.TypeOf(selector_ctx))
		panic(1)
	}
}
