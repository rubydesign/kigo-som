package main

import (
	"kigo-som/cst"
	"kigo-som/sool"
	"os"
)




func main() {
	classdef := cst.ClassdefFromFile(os.Args[1])
	class := sool.MakeClass(classdef)
	sool.PrintClass(class)
  return
}
