package main

import (
  "fmt"
  "strconv"
)

func compareNumbers(x, y int) string {
  if x > y {
    return fmt.Sprintf("%s is greater than %s", strconv.Itoa(x), strconv.Itoa(y))
  } else {
    return fmt.Sprintf("%s is smaller than %s", strconv.Itoa(x), strconv.Itoa(y))
  }
}

func countTo(n int) {
    for i := 1; i <= n; i++ {
        fmt.Println( i )
    }
    return
}

func breakCountTo(n int) {
    for i := 1; i <= n; i++ {

        if i > 2 { break }

        fmt.Println( i )
    }
    return
}

func countOdd(n int) {
    for i := 1; i <= n; i++ {

        if i % 2 == 0 { continue }

        fmt.Println( i )
    }
    return
}

func main() {
  fmt.Println( compareNumbers(2,1) )
  fmt.Println( compareNumbers(2,5) )

  fmt.Println( "" )

  fmt.Println( "count to three..." )
  countTo(3)

  fmt.Println( "" )

  fmt.Println( "break count to ten..." )
  breakCountTo(10)

  fmt.Println( "" )

  fmt.Println( "count odd..." )
  countOdd(5)
}
