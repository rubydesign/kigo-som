package cst

import (
    "kigo-som/parser"
    "log"
    "reflect"
)

// expression:
//     assignation | evaluation;
//
// assignation:
//     assignments evaluation;
//
// assignments:
//     assignment+;
//
// assignment:
//     variable Assign;
//
// evaluation:
//     primary messages?;
//
// primary:
//     variable | nestedTerm | nestedBlock | literal;


type Assignment struct {
  variable string
}
type Assignments struct {
  assignments []*Assignment
}
type Assignation struct {
  assignments *Assignments
  evaluation *Evaluation
}
type Primary struct {

}
type Message  struct {

}
type Evaluation struct {
  primary  *Primary
  messages []*Message
}

type Expression struct {
  assignation *Assignation  // OR
  evaluation  *Evaluation
}
func  MakeAssignments(ctx *parser.AssignmentsContext) (*Assignments) {
  return &Assignments{nil}
}

func MakeAssignation(ctx *parser.AssignationContext) (*Assignation){
  assignments_ctx := ctx.Assignments()// IAssignmentsContext
  evaluation_ctx := ctx.Evaluation() //IEvaluationContext
  assignments := MakeAssignments(assignments_ctx.(*parser.AssignmentsContext) )
  evaluation := MakeEvaluation(evaluation_ctx.(*parser.EvaluationContext))
  return &Assignation{ assignments , evaluation }
}
func MakePrimary(ctx *parser.PrimaryContext) (*Primary){
  return &Primary{}
}
func MakeMessages(ctx parser.IMessagesContext) ([]*Message){
  messages := make([]*Message , 0 ,  3)
  if ctx == nil { return messages }
  //message_ctx := ctx.(*parser.MessagesContext)
  return messages
}

func MakeEvaluation(ctx *parser.EvaluationContext) (*Evaluation){
  primary_ctx := ctx.Primary()
  primary := MakePrimary(primary_ctx.(*parser.PrimaryContext))
  messages_ctx := ctx.Messages()
  messages := MakeMessages(messages_ctx)
  return &Evaluation{ primary , messages}
}

func MakeExpression(ctx *parser.ExpressionContext) (*Expression) {
  var evaluation *Evaluation = nil
  var assignation *Assignation = nil
  log.Println("ctx" , ctx , reflect.TypeOf(ctx))
  ctx.Evaluation()
  assignation_ctx := ctx.Assignation()
  if  assignation_ctx != nil {
    log.Println("assignation_ctx" , assignation_ctx , reflect.TypeOf(assignation_ctx))
    assignation = MakeAssignation(assignation_ctx.(*parser.AssignationContext))
  }
  evaluation_ctx := ctx.Evaluation()
  if  evaluation_ctx != nil {
    log.Println("evaluation_ctx" , evaluation_ctx , reflect.TypeOf(evaluation_ctx))
    evaluation = MakeEvaluation(evaluation_ctx.(*parser.EvaluationContext))
  }
  return &Expression{assignation,evaluation}
}
