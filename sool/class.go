package sool

import (
  "kigo-som/cst"
    "fmt"
    "strings"
)

// type Classdef struct{
// 	name  string
// 	super string
// 	instance_variables []string
// 	instance_methods   []*Method
//   class_variables  []string
//   class_methods    []*Method
// }

type Class struct {
  Name  string
  Super string
  InstanceVariables []string
  InstanceMethods   []*Method
  // meta   *Class
}


func PrintClass(class *Class)  {
  fmt.Println("Class:" , class.Name , " < " , class.Super)
  fmt.Println( "-", "@  " , strings.Join(class.InstanceVariables , " "))
  // fmt.Println(pre , "@@ " , strings.Join(classdef.class_variables , " "))
  fmt.Println("  " , "Methods:" , len(class.InstanceMethods))
  for value := range class.InstanceMethods {
    PrintMethod(class.InstanceMethods[value])
  }
}

func MakeClass(classdef *cst.Classdef) (*Class){
  methods := MakeMethods(classdef)
  class := &Class{ classdef.Name , classdef.Super , classdef.InstanceVariables , methods}
  return class
}

func MakeMethods( classdef *cst.Classdef ) ([]*Method){
  methods := make([]*Method, 0, 3)
  //fmt.Println("methods" , len(methods_ctx))
  for value := range classdef.InstanceMethods {
    meth  := classdef.InstanceMethods[value]
		method := MakeMethod(meth)
		methods = append(methods , method)
  }
  return methods
}
