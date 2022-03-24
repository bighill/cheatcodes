package main

import (
    "fmt"
)

func subtactTwo(numbers []int) {  
    for i := range numbers {
        numbers[i] -= 2
    }
}

func countries() []string {  
    countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
    neededCountries := countries[:len(countries)-2]
    countriesCpy := make([]string, len(neededCountries))
    copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
    return countriesCpy
}

func main() {

    // make a slice from an array
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4] //creates a slice from a[1] to a[3]
    fmt.Println(b)

    // straight to slice in one declaration
    c := []int{6, 7, 8} //creates and array and returns a slice reference
    fmt.Println(c)

    // modify slice and underlying array
    darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
    dslice := darr[2:5]
    fmt.Println("array before",darr)
    for i := range dslice {
        dslice[i]++
    }
    fmt.Println("array after",darr)

    // multiple slices can be used to edit an array
    numa := [3]int{78, 79 ,80}
    nums1 := numa[:] //creates a slice which contains all elements of the array
    nums2 := numa[:]
    fmt.Println("array before change 1",numa)
    nums1[0] = 100
    fmt.Println("array after modification to slice nums1", numa)
    nums2[1] = 101
    fmt.Println("array after modification to slice nums2", numa)

    // slice length vs slice capacity
    fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
    fruitslice := fruitarray[1:3]
    fmt.Printf("length is %d and capacity is %d\n", len(fruitslice), cap(fruitslice)) //length is 2 and capacity is 6

    fruitslice = fruitslice[:cap(fruitslice)] //re-slicing fruitslice till its capacity
    fmt.Println("After re-slicing length is",len(fruitslice), "and capacity is",cap(fruitslice))

    // appending to a slice
    cars := []string{"Ferrari", "Honda", "Ford"}
    fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
    cars = append(cars, "Toyota")
    fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

    // "nil" is a 0 length slice
    var names []string
    if names == nil {
        fmt.Println("append because slice is nil")
        names = append( names, "John", "Sebastian", "Vinay" )
        fmt.Println( "names contents:",names )
    }

    // merge multiple slices
    veggies := []string{ "carrot", "celery" }
    fruit := []string{ "apple", "plum" }
    food := append( veggies, fruit... )
    fmt.Println( "food:", food )

    // when splices are sent as params to a function
    // that function has the ability to change the splice
    nos := []int{8, 7, 6}
    fmt.Println("slice before function call", nos)
    subtactTwo(nos)                               //function modifies the slice
    fmt.Println("slice after function call", nos) //modifications are visible outside

    // multidimensional splices
    pls := [][]string {
            {"C", "C++"},
            {"JavaScript"},
            {"Go", "Rust"},
            }
    for _, v1 := range pls {
        for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
        }
        fmt.Printf("\n")
    }

    // arrays cannot be garbage collected if slice is still in memory
    // can use copy() to create a new array/slice and allow old array to be garbage collected
    countriesNeeded := countries()
    fmt.Println(countriesNeeded)
}
