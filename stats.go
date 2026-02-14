package main

import (
	"kigo-som/cst"
	"os"
	"fmt"
	"strings"
)




func main() {
	dir := os.Args[1]
	entries, _ := os.ReadDir(dir)
	mapp := make(map[string]*cst.Classdef)

	 for _, e := range entries {
		 if !strings.HasSuffix(e.Name() , ".som") {continue}
		 file := dir + e.Name()
		 fmt.Print(e.Name() , ": ")
		 clazz := cst.ClassdefFromFile(file)
		 mapp[file] = clazz
 	}

	 // for key, _ := range mapp {
		//  fmt.Println(key)
	 // }

	 return
}
