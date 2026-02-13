package cst

import (
    "kigo-som/parser"
    "log"
    "reflect"
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
  selector []string
  operand *BinaryOperand
}

func MakeBinaryOperand(ctx *parser.BinaryOperandContext) (*BinaryOperand){
  primary_ctx := ctx.Primary()
  primary := MakePrimary(primary_ctx.(*parser.PrimaryContext))
  unary := make([]string , 0 ,  3)
  for i := range ctx.AllUnaryMessage() {
    unary_ctx := ctx.UnaryMessage(i)
    method_name := unary_ctx.UnarySelector().Identifier().GetText()
    log.Println("method_name" , method_name , reflect.TypeOf(method_name))
    unary = append( unary , method_name )
  }
  return &BinaryOperand{primary , unary}
}

func MakeBinaryMessage(ctx *parser.BinaryMessageContext) (Message){
  selector_ctx := ctx.BinarySelector() //IBinarySelectorContext {
  selector := MakeBinarySelector(selector_ctx.(*parser.BinarySelectorContext))
  operarand_ctx := ctx.BinaryOperand() // {
  operand := MakeBinaryOperand(operarand_ctx.(*parser.BinaryOperandContext))
  message := &BinaryMessage{ selector , operand}
  return message
}

func MakeMessages(ctx parser.IMessagesContext) ([]Message){
  messages := make([]Message , 0 ,  3)
  if ctx == nil { return messages }
  for i := range ctx.AllUnaryMessage() {
    unary_ctx := ctx.UnaryMessage(i)
    method_name := unary_ctx.UnarySelector().Identifier().GetText()
    log.Println("method_name" , method_name , reflect.TypeOf(method_name))
    messages = append( messages , &UnaryMessage{[]string{method_name} } )
  }
  for bi := range ctx.AllBinaryMessage() {
    binary_ctx := ctx.BinaryMessage(bi)
    message := MakeBinaryMessage( binary_ctx.(*parser.BinaryMessageContext) )
    log.Println("message" , message , reflect.TypeOf(message))
    messages = append( messages , message )
  }
  return messages
}
