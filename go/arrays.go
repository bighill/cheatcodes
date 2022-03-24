package main

import (
    "fmt"
)

func printarray(a [3][2]string) {  
    for _, v1 := range a {
        for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
        }
        fmt.Printf("\n")
    }
}

func main() {

    // int array with length 3
    var a [3]int
    fmt.Println( a )

    a[0] = 11
    a[1] = 22
    a[2] = 33
    fmt.Println( a )

    // short hand declaration
    b := [3]int{1, 2, 3}
    fmt.Println( b )

    // partial short hand declaration
    c := [3]int{7}
    fmt.Println( c )

    // make compilier determine length
    d := [...]int{5, 6, 7, 8}
    fmt.Println( d )

    // the following doesn't work because [3]int and [5]int are distinct types
    /*
    a := [3]int{5, 78, 8}
    var b [5]int
    b = a
    */

    // when arrays are assigned new variable, a copy of the original array is assigned to the new variable
    color := [...]string{"red", "green", "blue"}
    colorCopy := color
    colorCopy[0] = "orange"
    fmt.Println( "color[0] is", color[0] )
    fmt.Println( "colorCopy[0] is", colorCopy[0] )

    // get length of array
    arr := [...]float64{67.7, 89.8, 21, 78}
    fmt.Println("length of arr is",len(arr))

    // loop over array
    loopArr := [...]float64{67.7, 89.8, 21, 78}
    for i := 0; i < len(loopArr); i++ { //looping from 0 to the length of the array
        fmt.Printf("%d th element of loopArr is %.2f\n", i, loopArr[i])
    }

    // loop array with range (better)
    x := [...]string{"hey", "hi", "hello"}
    for i, v := range x {
        fmt.Printf( "index is %d and value is %s\n", i, v )
    }

    // multidimensional arrays
    animals := [3][2]string{
        {"lion", "tiger"},
        {"cat", "dog"},
        {"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
    }
    printarray( animals )

    // alternate multidimensional arrays
    var biz [3][2]string
    biz[0][0] = "apple"
    biz[0][1] = "samsung"
    biz[1][0] = "microsoft"
    biz[1][1] = "google"
    biz[2][0] = "AT&T"
    biz[2][1] = "T-Mobile"
    fmt.Printf("\n")
    printarray( biz )
}
