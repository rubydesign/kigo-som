package sool

import (
  "kigo-som/cst"
  // "fmt"
  // "strings"
)


type Statement interface{
  using() string
}

type ReturnStatement struct {
  value string
}

func (ret ReturnStatement) using() string {
  return ret.value
}

func MakeReturnStatement( exp *cst.Expression ) ([]Statement){
  statements , variable := ReduceExpressionToVariable(exp)
  return append(statements , ReturnStatement{ variable } )
}

// assign the block to tmp var, extracting any recursion on the way
// returned variable is the one the block is assigned to
func MakeAssignmentForNestedBlock( nestedTerm *cst.NestedBlock)([]Statement , string){
  panic("MakeAssignmentForNestedBlock")
  return nil , ""
}

func MakeAssignmentForAssignation( ass * cst.Assignation ) ([]Statement){
  panic("MakeAssignmentForNestedBlock")
  return nil
}

func ReduceEvaluationToVariable(eval *cst.Evaluation ) ([]Statement , string) {
  statements := make([]Statement , 0 , 3)
  primary := eval.Primary
  if len(eval.Messages) == 0 && primary.NestedTerm == nil &&
      primary.NestedBlock == nil && primary.Literal == nil {
    return statements ,  primary.Variable
  }
  if primary.NestedTerm != nil {
    return ReduceExpressionToVariable( primary.NestedTerm )
  } else if primary.NestedBlock != nil {
    // deal with the assignment variable, ie extract + return
    return MakeAssignmentForNestedBlock( primary.NestedBlock )
  } else if primary.Literal != nil {
    return MakeAssignmentForLiteral( primary.Literal )
  }
  cst.PrintEvaluation("panic " , eval)
  panic( "Message and variable case")
}

func ReduceExpressionToVariable(exp *cst.Expression ) ([]Statement , string) {
  if eval := exp.Evaluation ; eval != nil  {
      return ReduceEvaluationToVariable( eval )
  } else {
    panic( "deal with assignation")
    MakeAssignmentForAssignation( exp.Assignation )
  }
  return nil , ""
}

func MakeStatementFromExpression(expression *cst.Expression) []Statement {
  if expression.Assignation != nil {
    return MakeAssignmentForAssignation( expression.Assignation )
  } else {
    statements , _ := ReduceExpressionToVariable(expression)
    // remove last ?
    return statements
  }
  panic( "MakeStatementFromExpression")

  return nil
}
