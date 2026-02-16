package cst

import (
    "kigo-som/parser"
    "fmt"
    "strings"
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

type UnaryMessage struct {
  message []string  // selectors and args (alternate)
}
type BinaryMessage struct {
  selector string
  operand *BinaryOperand
}
type KeywordMessage struct {
  keywords []string
  formuli  []*Formula
}
type Message struct {
  typ int  // 1 unary , 2 binary , 3 keyword
  message []string
  operand *BinaryOperand  // only for binary
  formuli  []*Formula // only for keyword
}
type BinaryOperand struct {
  primary *Primary
  unary   []string
}
type Formula struct{
  operand *BinaryOperand
  messages []*Message; // binary
}

func PrintFormula(pre string , formula *Formula){
  fmt.Println(pre , "Formula:" )
  PrintBinaryOperand( "|  " + pre, formula.operand)
  for i := range formula.messages {
    message := formula.messages[i]
    fmt.Println("|  " + pre + " binary: ", message.message   )
    PrintBinaryOperand( "|  " + pre, message.operand)
  }
}

func PrintBinaryOperand( pre string, operand *BinaryOperand){
  fmt.Println(pre  , "BinaryOperand: " )
  PrintPrimary( "|  " + pre , operand.primary)
  if len(operand.unary) > 0 {
    fmt.Println(pre + "  unaries: ", strings.Join(operand.unary , ".") )
  }
}

func PrintMessage( pre string, message *Message){
  switch message.typ {
  case 1:
    fmt.Println(pre , "Message Unary: ", strings.Join(message.message , ".")   )
  case 2:
    fmt.Println(pre  ,"MessageBinary: ", message.message   )
    PrintBinaryOperand( pre , message.operand)
  case 3:
    fmt.Println(pre , "MessageKeywords: ", strings.Join(message.message , ".")   )
    for i := range message.formuli {
      formula := message.formuli[i]
      PrintFormula("|  " + pre , formula)
    }
  default: panic("message type " )
  }//switch
}

func PrintMessages( pre string, messages []*Message){
  if len( messages) == 0 { return }
  fmt.Println(pre , "Messages:" )
  for i := range messages {
    message := messages[i]
    PrintMessage("|  " + pre , message)
  }
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

func MakeBinaryMessage(ctx *parser.BinaryMessageContext) (*Message){
  selector_ctx := ctx.BinarySelector() //IBinarySelectorContext {
  selector := MakeBinarySelector(selector_ctx.(*parser.BinarySelectorContext))
  operarand_ctx := ctx.BinaryOperand() // {
  operand := MakeBinaryOperand(operarand_ctx.(*parser.BinaryOperandContext))
  messages := []string{selector}
  return &Message{ 2 , messages , operand , nil}
}

func MakeFormula(ctx *parser.FormulaContext) (*Formula)  {
  operarand_ctx := ctx.BinaryOperand()
  operand := MakeBinaryOperand(operarand_ctx.(*parser.BinaryOperandContext))
  messages := make([]*Message , 0 ,  3)
  binaries_ctx:= ctx.AllBinaryMessage()
  for bi := range binaries_ctx {
    binary_ctx := binaries_ctx[bi]
    message := MakeBinaryMessage( binary_ctx.(*parser.BinaryMessageContext) )
    //fmt.Println("binary" , message , reflect.TypeOf(message))
    messages = append( messages , message )
  }
  return &Formula{operand , messages}
}
func MakeKeywordMessage(ctx *parser.KeywordMessageContext) (*Message){
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
  return &Message{ 3 , keywords , nil , formuli}
}

func MakeMessages(ctx parser.IMessagesContext) ([]*Message){
  messages := make([]*Message , 0 ,  3)
  if ctx == nil { return messages }
  unaries_ctx := ctx.AllUnaryMessage()
  for i := range unaries_ctx {
    unary_ctx := unaries_ctx[i]
    method_name := unary_ctx.UnarySelector().Identifier().GetText()
    //fmt.Println("unary" , method_name , reflect.TypeOf(method_name))
    messages = append( messages , &Message{ 1 , []string{method_name} , nil,nil } )
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
