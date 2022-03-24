package main

import (
    "fmt"
)

func find(num int, nums ...int) {

    // nums will represent the variable number of parameters as a slice

    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}
func main() {
    find(89, 89, 90, 95)
    find(45, 56, 67, 45, 90, 109)
    find(78, 38, 56, 98)
    find(87)

    //
    nums := []int{89, 90, 95}
    //find(89, nums) // can't pass slice to variadic func that expects params list
    find( 89, nums... ) // but this will work by expanding the slice
}
