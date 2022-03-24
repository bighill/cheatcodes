package rectangle

import (
    "fmt"
    "math"
)

// in go, all exported functions need to be capitalized (e.g. Area, Diagonal)

func init() {
    fmt.Println( "rectangle package initialized" )
}

func Area(length, width float64) float64 {
    area := length * width
    return area
}

func Diagonal(length, width float64) float64 {
    diagonal := math.Sqrt((length * length) + (width * width))
    return diagonal
}
