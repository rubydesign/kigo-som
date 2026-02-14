package cst

import (
    "kigo-som/parser"
     // "fmt"
     // "reflect"
)

// primary:
//     variable | nestedTerm | nestedBlock | literal;
// nestedTerm:
//     NewTerm expression EndTerm;

// collaping nestedTerm into Expression

type Primary struct {
  variable    string
  nestedTerm  *Expression
  nestedBlock *NestedBlock
  literal     *Literal
}

func MakePrimary(ctx *parser.PrimaryContext) (*Primary){
  if variable_ctx := ctx.Variable() ; variable_ctx != nil {
    name := variable_ctx.(*parser.VariableContext).Identifier().GetText()
    return &Primary{name , nil , nil , nil}
  }
  if nestedTerm_ctx := ctx.NestedTerm() ; nestedTerm_ctx != nil {
    nested_expression := nestedTerm_ctx.Expression()
    expression := MakeExpression(nested_expression.(*parser.ExpressionContext))
    //fmt.Println("nested_expression " , nested_expression , reflect.TypeOf(nested_expression))
    return &Primary{"" , expression , nil , nil}
  }
  if nestedBlock_ctx := ctx.NestedBlock() ; nestedBlock_ctx != nil {
    nestedBlock := MakeNestedBlock( nestedBlock_ctx.(*parser.NestedBlockContext) )
    //fmt.Println("nestedBlock " , nestedBlock , reflect.TypeOf(nestedBlock))
    return &Primary{"" , nil , nestedBlock , nil}
  }
  literal_ctx := ctx.Literal() // by now this can't be anything else
  literal := MakeLiteral( literal_ctx.(*parser.LiteralContext) )
  return &Primary{ "" , nil , nil , literal}

}
