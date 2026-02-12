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

type Message  struct {
  typ int // 1 unary , 2 binary , 3 keyword
  message []string  // selectors and args (alternate)
}

func MakeBinaryMessage(ctx *parser.BinaryMessageContext) (*Message){
  selector_ctx := ctx.BinarySelector() //IBinarySelectorContext {
  selector := MakeBinarySelector(selector_ctx.(*parser.BinarySelectorContext))
  message := &Message{2 , selector}
  operarand_ctx := ctx.BinaryOperand() //IBinaryOperandContext {
  primary_ctx := operarand_ctx.Primary()
  //primary := 
  MakePrimary(primary_ctx.(*parser.PrimaryContext))
  return message
}

func MakeMessages(ctx parser.IMessagesContext) ([]*Message){
  messages := make([]*Message , 0 ,  3)
  if ctx == nil { return messages }
  for i := range ctx.AllUnaryMessage() {
    unary_ctx := ctx.UnaryMessage(i)
    method_name := unary_ctx.UnarySelector().Identifier().GetText()
    log.Println("method_name" , method_name , reflect.TypeOf(method_name))
    messages = append( messages , &Message{1 , []string{method_name} } )
  }
  for bi := range ctx.AllBinaryMessage() {
    binary_ctx := ctx.BinaryMessage(bi)
    message := MakeBinaryMessage( binary_ctx.(*parser.BinaryMessageContext) )
    log.Println("message" , message , reflect.TypeOf(message))
    messages = append( messages , message )
  }
  return messages
}
