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

func MakeStatement(expression *cst.Expression) []Statement {
  return nil
}
