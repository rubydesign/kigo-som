package sool

import (
  "kigo-som/cst"
  "fmt"
  // "strings"
)

// type Literal struct {
//   ltype     int    // 1 array , 2 sym , 3 string, 4 int , 5 double
//   array []*Literal
//   symbol   string // also string
//   number   int
//   decimal  float64
// }
type Literal interface{
  toGo() string
}
type IntLiteral struct {
  value int
}
type StringLiteral struct {
  value string
}

func (lit IntLiteral) toGo() string {
  return "1"
}
func (lit StringLiteral) toGo() string {
  return lit.value
}

// assign the literal to tmp var,  returned variable
func MakeAssignmentForLiteral( literal *cst.Literal)([]Statement , string){
  switch literal.Ltype {
  case 1:
    fmt.Println( "LiteralArray:" )
    for _,lit := range literal.Array {
      fmt.Println( lit)
    }
  case 2:
    fmt.Println( "LiteralSymbol:" , literal.Symbol  )
  case 3:
    fmt.Println( "LiteralString:" , literal.Symbol  )
  case 4:
    fmt.Println( "LiteralInt:" , literal.Number  )
  case 5:
    fmt.Println( "LiteralFloat:" , literal.Decimal  )
  }
  return nil , ""

}
