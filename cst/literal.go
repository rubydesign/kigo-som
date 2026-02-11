package cst

import (
    "kigo-som/parser"
     "log"
     "reflect"
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
  return nil
}
func MakeLiteralSymbol(ctx *parser.LiteralSymbolContext) (string){
  return ""
}

func MakeLiteralString(ctx *parser.LiteralStringContext) (string){
  return ""
}

func MakeLiteralDecimal(decimal_ctx *parser.LiteralDecimalContext) (int , int , float64){
  if literal_ctx := decimal_ctx.LiteralInteger() ; literal_ctx != nil {
    int_str := literal_ctx.(*parser.LiteralIntegerContext).Integer().GetText()
    int, _ := strconv.Atoi(int_str)
    log.Println("lit int " , int , reflect.TypeOf(int))
    return 4, int , 0.0
  } else {
    double_ctx := decimal_ctx.LiteralDouble().(*parser.LiteralDoubleContext)
    double_str := double_ctx.Double().GetText()
    log.Println("lit float " , double_str , reflect.TypeOf(double_str))
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
    log.Println("array " , array , reflect.TypeOf(array))
    return &Literal{1 , array , "" , 0, 0}
  }
  if symbol_ctx := ctx.LiteralSymbol() ; symbol_ctx != nil {
    symbol := MakeLiteralSymbol( symbol_ctx.(*parser.LiteralSymbolContext) )
    log.Println("symbol " , symbol , reflect.TypeOf(symbol))
    return &Literal{2 , nil ,symbol  , 0, 0}
  }
  if string_ctx := ctx.LiteralString() ; string_ctx != nil {
    string := MakeLiteralString( string_ctx.(*parser.LiteralStringContext) )
    log.Println("string " , string , reflect.TypeOf(string))
    return &Literal{2 , nil ,string  , 0, 0}
  }
  number_ctx := ctx.LiteralNumber()
  typ , int , float := MakeLiteralNumber( number_ctx.(*parser.LiteralNumberContext) )
  return &Literal{typ , nil , "" , int, float}
}
