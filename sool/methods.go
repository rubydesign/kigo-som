package sool

import (
  "kigo-som/cst"
  "fmt"
  "strings"
)


type Method struct{
  Name      string
  Arguments []string
  Primitive  bool
  // Locals []string
  // Statements []*Statement
}

func PrintMethod(method *Method){
  fmt.Println("  Method:" , method.Name , "(" , strings.Join(method.Arguments , " , " ) , ")")
}

func ExtractName(pattern []string) (string){
  if len(pattern) == 1 { return pattern[0]}
  if len(pattern) == 0 { panic("Pattern length 0")}
  if len(pattern) == 2 { return pattern[0]}
  return pattern[0]
}
func ExtractArgs(pattern []string) ([]string){
  if len(pattern) == 0 { panic("Pattern length 0")}
  if len(pattern) == 1 { return make([]string , 0,0 )}
  args := make([]string , 0, 3 )
  if len(pattern) == 2 { args = append(args , pattern[1])}
  return args
}

func MakeMethod(meth *cst.Method) (*Method){
//  fmt.Println(meth.Pattern)
  name := ExtractName(meth.Pattern)
  args := ExtractArgs(meth.Pattern)
  primitive := meth.MethodBlock.Primitive
  return &Method{ name , args , primitive}
}
