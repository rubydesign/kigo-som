package sool

import (
  "kigo-som/cst"
  "fmt"
  "strings"
)


type Block struct{
  Locals []string
  Statements []Statement
}

func PrintBlock(block *Block){
  if len(block.Locals) > 0 {
    fmt.Println("    Locals:" , strings.Join(block.Locals , " " ) )
  }
}

func MakeStatementsFromBlockBody(body *cst.BlockBody) []Statement {
  if body.Result != nil {
    return MakeReturnStatement( body.Result )
  } else {
    statements := MakeStatementFromExpression( body.Main )
    if body.Code != nil {
      more := MakeStatementsFromBlockBody( body.Code )
      statements = append(statements , more...)
    }
    return statements
  }
}

func MakeBlock(block *cst.MethodBlock) (*Block){
  locals := block.BlockContents.Locals
  statements := MakeStatementsFromBlockBody(block.BlockContents.BlockBody)
  return &Block{ locals , statements}
}
