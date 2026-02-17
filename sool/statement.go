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
func MakeReturnStatement( exp *cst.Expression ) ([]Statement){
  statements , variable := ReduceExpressionToVariable(exp)
  return append(statements , ReturnStatement{ variable } )
}

// assign the block to tmp var, extracting any recursion on the way
// returned variable is the one the block is assigned to
func MakeAssignmentForNestedBlock( nestedTerm *cst.NestedBlock)([]Statement , string){
  return nil , ""
}
func ReduceExpressionToVariable(exp *cst.Expression ) ([]Statement , string) {
  statements := make([]Statement , 0 , 3)
  if eval := exp.Evaluation ; eval != nil  {
    primary := eval.Primary
    if len(eval.Messages) == 0 && primary.NestedTerm == nil &&
        primary.NestedBlock == nil && primary.Literal == nil {
      return statements ,  primary.Variable
    }
    if primary.NestedTerm != nil {
      return ReduceExpressionToVariable( primary.NestedTerm )
    } else if primary.NestedBlock != nil {
      cst.PrintExpression("return" , exp)
      return MakeAssignmentForNestedBlock( primary.NestedBlock )
    }
  } else {
    panic( "deal with assignation")
  }
  return nil , ""
}

func (ret ReturnStatement) using() string {
  return ret.value
}

func MakeStatementFromExpression(expression *cst.Expression) []Statement {
  return nil
}
