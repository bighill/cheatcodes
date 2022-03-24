package main

import "fmt"

/*
func examplefunction( paramName paramType) returnType {
}
*/

/*
  sum1() and sum2() are effectively the same
  ...no need to declare type if all params are of the same type
*/
func sum1( x int, y int ) int {
  return x + y
}
func sum2( x, y int) int {
  return x + y
}

/*
  multiple return values. rad.
*/
func rectangle( length, width float64 )( float64, float64) {
  var area      = length * width
  var perimeter = (length + width) * 2

  return area, perimeter
}

func main() {
  fmt.Println( sum1(1,2), sum2(1,2) )

  area, perimeter := rectangle( 11.1, 22.2 )
  fmt.Printf( "area is %f and perimeter is %f\n", area, perimeter )

  // blank identifier _ use to discard a return value
  justarea, _ := rectangle( 33.3, 44.4 )
  _, justperimeter := rectangle( 55.5, 66.6 )
  fmt.Println( justarea )
  fmt.Println( justperimeter )
}
