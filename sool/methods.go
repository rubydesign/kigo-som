package sool

import (
  "kigo-som/cst"
  "fmt"
    // "strings"
)


type Method struct{
  Name      string
  Primitive  bool
  // Locals []string
  // Statements []*Statement
}

func PrintMethod(method *Method){
}

func PrintMethods(methods []*Method){
  fmt.Println("  " , "Methods:" , len(methods))
  for i := range methods {
    method := methods[i]
    PrintMethod(method)
  }
}

func MakeMethod(meth *cst.Method) (*Method){
  return nil
}
