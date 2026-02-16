package sool

import (
  "kigo-som/cst"
  "fmt"
  "strings"
)


type Block struct{
  Locals []string
  // Statements []*Statement
}

func PrintBlock(block *Block){
  if len(block.L  ocals) > 0 {
    fmt.Println("    Locals:" , strings.Join(block.Locals , " " ) )
  }
}
func MakeBlock(block *cst.MethodBlock) (*Block){
  locals := block.BlockContents.Locals
  return &Block{ locals }
}
