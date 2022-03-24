package main

import (
    "fmt"
)

// map is kind of like an object or python dict

func main() {

    // make a map like this
    personSalary := make(map[string]int)
    personSalary["steve"] = 12000
    personSalary["jamie"] = 15000
    personSalary["mike"] = 9000
    fmt.Println("personSalary map contents:", personSalary)

    // or make a map like this
    personAge := map[string]int {
        "lola" : 33,
        "mary" : 44,
    }
    personAge["mike"] = 55
    fmt.Println( "personAge:", personAge )

    // accessing map items
    employee := "mike"
    fmt.Println("Salary of", employee, "is", personSalary[employee], "and their age is", personAge[employee])

    // test whether key is present
    newEmployee := "betty"
    value, ok := personSalary[newEmployee]
    if ok == true {
        fmt.Println("Salary of", newEmployee, "is", value)
    } else {
        fmt.Println(newEmployee,"not found")
    }

    // loop over map
    for k, v := range personSalary {
        fmt.Printf( "personSalary[%s] = %d\n", k, v )
    }

    // delete items from map
    delete( personSalary, "steve" )
    fmt.Println( "map after deletion", personSalary )

    // length of map
    fmt.Println("length of personSalary is", len(personSalary))
}
