package main

import (
	"kigo-som/cst"
	"os"
)




func main() {
	classdef := cst.ClassdefFromFile(os.Args[1])
	cst.PrintClassdef(classdef)
  return
}
