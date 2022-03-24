package main

import (
    "fmt"
    "log"
    "big-go-examples/packages-example/rectangle"

    // blank identifier can be used to ensure initialization works even though it isn't actually used
    // compile would fail if "math" was not used and without blank identifier
    _ "math"
)

// 1. package level variables are initialized first
var rectLen, rectWidth float64 = 6, 7

// 2. init function is called second
func init() {
    println("main package initialized")

    if rectLen < 0 {
        log.Fatal("length is less than zero")
    }
    if rectWidth < 0 {
        log.Fatal("width is less than zero")
    }
}

// 3. main function is called last
func main() {
    fmt.Println("Geometrical shape properties")

    fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
    fmt.Printf("diagonal of the rectangle %.2f\n",rectangle.Diagonal(rectLen, rectWidth))
}
