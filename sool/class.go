package sool

import (
  "kigo-som/cst"
    "fmt"
    // "strings"
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
  // instance_variables []string
  // instance_methods   []*Method
  // meta   *Class
}

type Method struct {

}

func PrintClass(class *Class)  {
  fmt.Println("Class:" , class.Name , " < " , class.Super)
  // fmt.Println(pre , "-", "@  " , strings.Join(classdef.instance_variables , " "))
  // fmt.Println(pre , "@@ " , strings.Join(classdef.class_variables , " "))

}
func MakeClass(classdef *cst.Classdef) (*Class){
  class := &Class{ classdef.Name , classdef.Super}
  return class
}
