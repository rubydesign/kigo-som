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
  fmt.Println("  Method:" , method.Name + "(" , strings.Join(method.Arguments , " , " ) , ")")
}

func ExtractName(pattern []string) (string){
  names := make([]string , 0 ,3)
  for i := range pattern {
    if (i % 2) == 0 { names = append(names , strings.Replace(pattern[i] , ":" , "" , 1))}
  }
  return strings.Join(names , "_")
}

func ExtractArgs(pattern []string) ([]string){
  args := make([]string , 0 ,3)
  for i := range pattern {
    if (i % 2) == 1 { args = append(args , pattern[i] ) }
  }
  return args
}

func MakeMethod(meth *cst.Method) (*Method){
//  fmt.Println(meth.Pattern)
  name := ExtractName(meth.Pattern)
  args := ExtractArgs(meth.Pattern)
  primitive := meth.MethodBlock.Primitive
  return &Method{ name , args , primitive}
}
