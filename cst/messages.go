package cst

import (
    "kigo-som/parser"
    // "fmt"
    // "reflect"
)

// messages:
//       unaryMessage+ binaryMessage* keywordMessage?
//     | binaryMessage+ keywordMessage?
//     | keywordMessage;
//
// unaryMessage:
//     unarySelector;
//
// binaryMessage:
//     binarySelector binaryOperand;
//
// binaryOperand:
//     primary unaryMessage*;
//
// keywordMessage:
//     ( keyword formula )+;
// formula:
//    binaryOperand binaryMessage*;

type Message  interface {
}

type UnaryMessage struct{
  message []string  // selectors and args (alternate)
}
type BinaryOperand struct {
  primary *Primary
  unary   []string
}
type BinaryMessage struct{
  selector string
  operand *BinaryOperand
}
type Formula struct{
  operand *BinaryOperand
  messages []*BinaryMessage;
}

type KeywordMessage struct {
  keywords []string
  formuli  []*Formula
}

func MakeBinaryOperand(ctx *parser.BinaryOperandContext) (*BinaryOperand){
  primary_ctx := ctx.Primary()
  primary := MakePrimary(primary_ctx.(*parser.PrimaryContext))
  unary := make([]string , 0 ,  3)
  for i := range ctx.AllUnaryMessage() {
    unary_ctx := ctx.UnaryMessage(i)
    method_name := unary_ctx.UnarySelector().Identifier().GetText()
    // fmt.Println("unary" , method_name , reflect.TypeOf(method_name))
    unary = append( unary , method_name )
  }
  return &BinaryOperand{primary , unary}
}

func MakeBinaryMessage(ctx *parser.BinaryMessageContext) (Message){
  selector_ctx := ctx.BinarySelector() //IBinarySelectorContext {
  selector := MakeBinarySelector(selector_ctx.(*parser.BinarySelectorContext))
  operarand_ctx := ctx.BinaryOperand() // {
  operand := MakeBinaryOperand(operarand_ctx.(*parser.BinaryOperandContext))
  return &BinaryMessage{ selector , operand}
}

func MakeFormula(ctx *parser.FormulaContext) (*Formula)  {
  operarand_ctx := ctx.BinaryOperand()
  operand := MakeBinaryOperand(operarand_ctx.(*parser.BinaryOperandContext))
  messages := make([]*BinaryMessage , 0 ,  3)
  for bi := range ctx.AllBinaryMessage() {
    binary_ctx := ctx.BinaryMessage(bi)
    message := MakeBinaryMessage( binary_ctx.(*parser.BinaryMessageContext) )
    //fmt.Println("binary" , message , reflect.TypeOf(message))
    messages = append( messages , message.(*BinaryMessage) )
  }
  return &Formula{operand , messages}
}
func MakeKeywordMessage(ctx *parser.KeywordMessageContext) (Message){
  keywords := make([]string , 0,3)
  for idx := range ctx.AllKeyword() {
    keyword_ctx := ctx.Keyword(idx)
    name := keyword_ctx.Keyword().GetText()
    keywords = append(keywords , name )
  }
  formuli  := make([]*Formula , 0,3)
  for fi := range ctx.AllFormula() {
    formula_ctx := ctx.Formula(fi)
    formula := MakeFormula(formula_ctx.(*parser.FormulaContext))
    formuli = append(formuli , formula )
  }
  return &KeywordMessage{ keywords , formuli}
}

func MakeMessages(ctx parser.IMessagesContext) ([]Message){
  messages := make([]Message , 0 ,  3)
  if ctx == nil { return messages }
  for i := range ctx.AllUnaryMessage() {
    unary_ctx := ctx.UnaryMessage(i)
    method_name := unary_ctx.UnarySelector().Identifier().GetText()
    //fmt.Println("unary" , method_name , reflect.TypeOf(method_name))
    messages = append( messages , &UnaryMessage{[]string{method_name} } )
  }
  for bi := range ctx.AllBinaryMessage() {
    binary_ctx := ctx.BinaryMessage(bi)
    message := MakeBinaryMessage( binary_ctx.(*parser.BinaryMessageContext) )
    //fmt.Println("binary" , message , reflect.TypeOf(message))
    messages = append( messages , message )
  }
  if keyword_ctx := ctx.KeywordMessage() ; keyword_ctx != nil {
    message := MakeKeywordMessage( keyword_ctx.(*parser.KeywordMessageContext) )
    //fmt.Println("keyword" , message , reflect.TypeOf(message))
    messages = append( messages , message )
  }
  return messages
}
