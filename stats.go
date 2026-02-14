package main

import (
	"kigo-som/cst"
	"os"
	"fmt"
)




func main() {
	dir := "./SOM/Smalltalk/"
	entries, _ := os.ReadDir(dir)
	mapp := make(map[string]*cst.Classdef)

	 for _, e := range entries {
		 file := dir + e.Name()
		 clazz := cst.ClassdefFromFile(file)
		 mapp[file] = clazz
		 fmt.Println(e.Name())
 	}

	 for key, _ := range mapp {
		 fmt.Println(key)
	 }

	 return
}
