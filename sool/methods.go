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
  Block      *Block
}

func PrintMethod(method *Method){
  fmt.Print("  Method:" , method.Name + "(" , strings.Join(method.Arguments , " , " ) , ")")
  if method.Primitive {
    fmt.Println("  (primitive)")
  }else{
    fmt.Println("")
    PrintBlock(method.Block)
  }
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
  var block *Block
  if !primitive {
    block = MakeBlock( meth.MethodBlock )
  }
  return &Method{ name , args , primitive , block}
}
