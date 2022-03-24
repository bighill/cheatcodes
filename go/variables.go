package main

import (
  "fmt"
  "math"
)

func main() {
  /* var */
  var v = 11
  v = 33 // reassignment is allowed
  fmt.Println("var v =", v)

  /* const */
  const c = 11
  // c = 33 // reassignment not allowed
  fmt.Println("const c =", c)

  /* values from functions */
  var vv = math.Sqrt(81)
  // const cc = math.Sqrt(81) // bad, constants need to be known at compile time
  fmt.Println("var vv =", vv)

  /* bool */
  var bool0 bool = true
  bool1 := false
  boolX, boolY := bool0 && bool1, bool0 || bool1
  fmt.Println("bool boolX =", boolX)
  fmt.Println("bool boolY =", boolY)

  /* int */
  var int0 int = 99
  int1 := 55
  int2 := int0+int1
  fmt.Println("int2 =", int2)
  // int8, int16, int32, and int64 are available
  // but probably only need int

  /* floating point */
  float0, float1 := 5.54, 8.2
  fmt.Printf( "float0 is type %T and float1 is type %T\n", float0, float1 )

  /* strings */
  first := "john"
  last := "doe"
  name := first + " " + last
  fmt.Println("name =", name)

  /* type conversion */
  i := 55     // int
  j := 33.1   // float64
  //sum := i + j // bad
  sum := float64(i) + j //good
  fmt.Println("sum =", sum)
}
