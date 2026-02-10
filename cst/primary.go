package cst

import (
    "kigo-som/parser"
     "log"
     "reflect"
)

// primary:
//     variable | nestedTerm | nestedBlock | literal;


type Variable struct {
  variable string
}
type NestedTerm struct {

}
type NestedBlock struct {

}
type Primary struct {
  variable    *Variable
  nestedTerm  *NestedTerm
  nestedBlock *NestedBlock
  literal     *Literal
}

func MakeNestedTerm(ctx *parser.NestedTermContext) (*NestedTerm){
  return &NestedTerm{}
}
func MakeNestedBlock(ctx *parser.NestedBlockContext) (*NestedBlock){
  return &NestedBlock{}
}
func MakePrimary(ctx *parser.PrimaryContext) (*Primary){
  if variable_ctx := ctx.Variable() ; variable_ctx != nil {
    name := variable_ctx.(*parser.VariableContext).Identifier().GetText()
    return &Primary{&Variable{name} , nil , nil , nil}
  }
  if nestedTerm_ctx := ctx.NestedTerm() ; nestedTerm_ctx != nil {
    nestedTerm := MakeNestedTerm( nestedTerm_ctx.(*parser.NestedTermContext) )
    log.Println("nestedTerm " , nestedTerm , reflect.TypeOf(nestedTerm))
    return &Primary{nil , nestedTerm , nil , nil}
  }
  if nestedBlock_ctx := ctx.NestedBlock() ; nestedBlock_ctx != nil {
    nestedBlock := MakeNestedBlock( nestedBlock_ctx.(*parser.NestedBlockContext) )
    log.Println("nestedBlock " , nestedBlock , reflect.TypeOf(nestedBlock))
    return &Primary{nil , nil , nestedBlock , nil}
  }
  literal_ctx := ctx.Literal() // by now this can't be anything else
  literal := MakeLiteral( literal_ctx.(*parser.LiteralContext) )
  return &Primary{ nil , nil , nil , literal}

}
