package main

import (
	"kigo-som/cst"
	"kigo-som/sool"
	"os"
)




func main() {
	classdef := cst.ClassdefFromFile(os.Args[1])
	cst.PrintClassdef(classdef)
	class := sool.MakeClass(classdef)
	sool.PrintClass(class)
  return
}
