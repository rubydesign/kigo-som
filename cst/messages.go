package cst

import (
    "kigo-som/parser"
    // "log"
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

type Message  struct {

}

func MakeMessages(ctx parser.IMessagesContext) ([]*Message){
  messages := make([]*Message , 0 ,  3)
  if ctx == nil { return messages }
  //message_ctx := ctx.(*parser.MessagesContext)
  return messages
}
