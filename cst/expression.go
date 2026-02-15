package cst

import (
    "kigo-som/parser"
    "fmt"
    "strings"
    // "reflect"
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


type Assignation struct {
  assignments []string
  evaluation *Evaluation
}
type Evaluation struct {
  primary  *Primary
  messages []Message
}
type Expression struct {
  assignation *Assignation  // OR
  evaluation  *Evaluation
}

func PrintEvaluation(pre string , evaluation *Evaluation) {
  PrintPrimary( pre , evaluation.primary)
  PrintMessages( pre , evaluation.messages)
}

func PrintAssignation(pre string , assignation *Assignation) {
  fmt.Println(pre , strings.Join(assignation.assignments , " = ") , " = "  )
  PrintEvaluation("  " + pre , assignation.evaluation)
}

func PrintExpression(pre string , expression *Expression) {
  if expression.assignation != nil {
    PrintAssignation( pre , expression.assignation)
  } else {
    PrintEvaluation( pre , expression.evaluation)
  }
}

func  MakeAssignments(ctx *parser.AssignmentsContext) ([]string) {
  assignments := make([]string, 0,1) // usually just one assignment
  assignments_ctx := ctx.AllAssignment()
  for i := range  assignments_ctx{
    assignment := assignments_ctx[i]
    variable := assignment.Variable().GetText()
//    fmt.Println("variable" , variable , reflect.TypeOf(variable))
    assignments = append( assignments , variable)
  }
  return assignments
}

func MakeAssignation(ctx *parser.AssignationContext) (*Assignation){
  assignments_ctx := ctx.Assignments()// IAssignmentsContext
  evaluation_ctx := ctx.Evaluation() //IEvaluationContext
  assignments := MakeAssignments(assignments_ctx.(*parser.AssignmentsContext) )
  evaluation := MakeEvaluation(evaluation_ctx.(*parser.EvaluationContext))
  return &Assignation{ assignments , evaluation }
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
  //fmt.Println("ctx" , ctx , reflect.TypeOf(ctx))
  ctx.Evaluation()
  assignation_ctx := ctx.Assignation()
  if  assignation_ctx != nil {
    //fmt.Println("assignation_ctx" , assignation_ctx , reflect.TypeOf(assignation_ctx))
    assignation = MakeAssignation(assignation_ctx.(*parser.AssignationContext))
  }
  evaluation_ctx := ctx.Evaluation()
  if  evaluation_ctx != nil {
    //fmt.Println("evaluation_ctx" , evaluation_ctx , reflect.TypeOf(evaluation_ctx))
    evaluation = MakeEvaluation(evaluation_ctx.(*parser.EvaluationContext))
  }
  return &Expression{assignation,evaluation}
}
