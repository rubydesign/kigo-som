package cst

import (
    "kigo-som/parser"
     // "fmt"
     // "reflect"
     "strconv"
)

// literal:
//     literalArray | literalSymbol | literalString | literalNumber;
//
// literalArray:
//     Pound NewTerm
//     literal*
//     EndTerm;
//
// literalNumber:
//     negativeDecimal | literalDecimal;
//
// literalDecimal:
//     literalInteger | literalDouble;
//
// negativeDecimal:
//     Minus literalDecimal;
//
// literalInteger:
//     Integer;
//
// literalDouble:
//     Double;
//
// literalSymbol:
//     Pound ( string | selector );
//
// literalString:
//     string;

type Literal struct {
  ltype     int    // 1 array , 2 sym , 3 string, 4 int , 5 double
  array []*Literal
  symbol   string // also string
  number   int
  decimal  float64
}

func MakeLiteralArray(ctx *parser.LiteralArrayContext) ([]*Literal){
  literals := make([]*Literal , 0, 3)
  literals_ctx := ctx.AllLiteral()
  for i := range literals_ctx {
    literal_ctx :=  literals_ctx[i].(*parser.LiteralContext)
    literal := MakeLiteral(literal_ctx)
    literals = append(literals , literal)
  }
  return literals
}

func MakeLiteralSymbol(ctx *parser.LiteralSymbolContext) (string){
  if string_ctx := ctx.String_() ; string_ctx != nil {
    st_ctx := string_ctx.STString()
    return st_ctx.GetText()
  }
  selector_ctx := ctx.Selector()
  selector := MakeSelector(selector_ctx.(*parser.SelectorContext))
  return selector
}

func MakeLiteralString(ctx *parser.LiteralStringContext) (string){
  string_ctx := ctx.String_()
  st_ctx := string_ctx.STString()
  return st_ctx.GetText()
}

func MakeLiteralDecimal(decimal_ctx *parser.LiteralDecimalContext) (int , int , float64){
  if literal_ctx := decimal_ctx.LiteralInteger() ; literal_ctx != nil {
    int_str := literal_ctx.(*parser.LiteralIntegerContext).Integer().GetText()
    int, _ := strconv.Atoi(int_str)
    //fmt.Println("lit int " , int , reflect.TypeOf(int))
    return 4, int , 0.0
  } else {
    double_ctx := decimal_ctx.LiteralDouble().(*parser.LiteralDoubleContext)
    double_str := double_ctx.Double().GetText()
    //fmt.Println("lit float " , double_str , reflect.TypeOf(double_str))
    double, _ := strconv.ParseFloat(double_str, 64)
    return 5, 0 , double
  }
}

func MakeLiteralNumber(ctx *parser.LiteralNumberContext) (int , int , float64){
  if negative_ctx := ctx.NegativeDecimal() ; negative_ctx != nil {
    decimal_ctx := negative_ctx.LiteralDecimal().(*parser.LiteralDecimalContext)
    typ , int , float := MakeLiteralDecimal(decimal_ctx)
    return typ , -int , -float
  } // else
  decimal_ctx := ctx.LiteralDecimal().(*parser.LiteralDecimalContext)
  return MakeLiteralDecimal( decimal_ctx)
}

func MakeLiteral(ctx *parser.LiteralContext) (*Literal){
  if array_ctx := ctx.LiteralArray() ; array_ctx != nil {
    array := MakeLiteralArray( array_ctx.(*parser.LiteralArrayContext) )
    //fmt.Println("array" , array , reflect.TypeOf(array))
    return &Literal{1 , array , "" , 0, 0}
  }
  if symbol_ctx := ctx.LiteralSymbol() ; symbol_ctx != nil {
    symbol := MakeLiteralSymbol( symbol_ctx.(*parser.LiteralSymbolContext) )
    // fmt.Println("symbol" , symbol , reflect.TypeOf(symbol))
    return &Literal{2 , nil ,symbol  , 0, 0}
  }
  if string_ctx := ctx.LiteralString() ; string_ctx != nil {
    string := MakeLiteralString( string_ctx.(*parser.LiteralStringContext) )
    //fmt.Println("string" , string , reflect.TypeOf(string))
    return &Literal{2 , nil ,string  , 0, 0}
  }
  number_ctx := ctx.LiteralNumber()
  typ , int , float := MakeLiteralNumber( number_ctx.(*parser.LiteralNumberContext) )
  return &Literal{typ , nil , "" , int, float}
}
