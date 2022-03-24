package main

import (
    "fmt"
)

// declare a structure
type Employee struct {
    firstName string
    lastName  string
    age       int
    salary    int

    // or like this
    /*
    firstName, lastName string
    age, salary         int
    */
}

// delare structs for nesting
type Address struct {
    city, state string
}
type Person struct {
    name string
    age int
    address Address
}

func main() {

    // creating structure using field names (okay to assign out of order)
    emp1 := Employee{
        firstName: "Sam",
        age:       25,
        salary:    500,
        lastName:  "Anderson",
    }

    // creating structure without using field names (just need to be sure to add in order)
    emp2 := Employee{"Thomas", "Paul", 29, 800}

    // anonymous structure for one-time use
    emp3 := struct {
        firstName, lastName string
        age, salary         int
    }{
        firstName: "Andreah",
        lastName:  "Nikola",
        age:       31,
        salary:    5000,
    }

    fmt.Println("Employee 1", emp1)
    fmt.Println("Employee 2", emp2)
    fmt.Println("Employee 3", emp3)

    // accessing individual fields
    emp6 := Employee{"Sam", "Anderson", 55, 6000}
    fmt.Println("First Name:", emp6.firstName)
    fmt.Println("Last Name:", emp6.lastName)
    fmt.Println("Age:", emp6.age)
    fmt.Printf("Salary: $%d\n", emp6.salary)

    // nested structs
    var p Person
    p.name = "Naveen"
    p.age = 50
    p.address = Address {
        city: "Chicago",
        state: "Illinois",
    }
    fmt.Println("Name:", p.name)
    fmt.Println("Age:",p.age)
    fmt.Println("City:",p.address.city)
    fmt.Println("State:",p.address.state)
}
