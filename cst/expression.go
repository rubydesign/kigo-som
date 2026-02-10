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


type Assignation struct {
  assign string
}
type Evaluation struct {
  eval string
}

type Expression struct {
  assignation *Assignation  // OR
  evaluation  *Evaluation
}

func MakeAssignation(ctx *parser.AssignationContext) (*Assignation){
  return nil
}
func MakeEvaluation(ctx *parser.EvaluationContext) (*Evaluation){
  return nil
}

func MakeExpression(ctx *parser.ExpressionContext) (*Expression) {
  var evaluation *Evaluation = nil
  var assignation *Assignation = nil
  log.Println("ctx" , ctx , reflect.TypeOf(ctx))
  ctx.Evaluation() //IEvaluationContext
  assignation_ctx := ctx.Assignation() // IAssignationContext
  if  assignation_ctx != nil {
    log.Println("assignation_ctx" , assignation_ctx , reflect.TypeOf(assignation_ctx))
    assignation = MakeAssignation(assignation_ctx.(*parser.AssignationContext))
  }
  evaluation_ctx := ctx.Evaluation() // IAssignationContext
  if  evaluation_ctx != nil {
    log.Println("evaluation_ctx" , evaluation_ctx , reflect.TypeOf(evaluation_ctx))
    evaluation = MakeEvaluation(evaluation_ctx.(*parser.EvaluationContext))
  }
  return &Expression{assignation,evaluation}
}
