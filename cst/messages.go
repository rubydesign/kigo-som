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
  unaries_ctx := ctx.AllUnaryMessage()
  for i := range unaries_ctx {
    unary_ctx := unaries_ctx[i]
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
  binaries_ctx:= ctx.AllBinaryMessage()
  for bi := range binaries_ctx {
    binary_ctx := binaries_ctx[bi]
    message := MakeBinaryMessage( binary_ctx.(*parser.BinaryMessageContext) )
    //fmt.Println("binary" , message , reflect.TypeOf(message))
    messages = append( messages , message.(*BinaryMessage) )
  }
  return &Formula{operand , messages}
}
func MakeKeywordMessage(ctx *parser.KeywordMessageContext) (Message){
  keywords := make([]string , 0,3)
  keywords_ctx := ctx.AllKeyword()
  for idx := range keywords_ctx {
    keyword_ctx := keywords_ctx[idx]
    name := keyword_ctx.Keyword().GetText()
    keywords = append(keywords , name )
  }
  formuli  := make([]*Formula , 0,3)
  formuli_ctx := ctx.AllFormula()
  for fi := range formuli_ctx {
    formula_ctx := formuli_ctx[fi]
    formula := MakeFormula(formula_ctx.(*parser.FormulaContext))
    formuli = append(formuli , formula )
  }
  return &KeywordMessage{ keywords , formuli}
}

func MakeMessages(ctx parser.IMessagesContext) ([]Message){
  messages := make([]Message , 0 ,  3)
  if ctx == nil { return messages }
  unaries_ctx := ctx.AllUnaryMessage()
  for i := range unaries_ctx {
    unary_ctx := unaries_ctx[i]
    method_name := unary_ctx.UnarySelector().Identifier().GetText()
    //fmt.Println("unary" , method_name , reflect.TypeOf(method_name))
    messages = append( messages , &UnaryMessage{[]string{method_name} } )
  }
  binaries_ctx := ctx.AllBinaryMessage()
  for bi := range binaries_ctx {
    binary_ctx := binaries_ctx[bi]
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
