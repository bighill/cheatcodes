package main

import (
    "fmt"
)

func printBytes(s string) {
    for i:= 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
}

func printChars(s string) {
    runes := []rune(s) // rune in this manner will handle special characters
    for i:= 0; i < len(runes); i++ {
        fmt.Printf("%c ",runes[i])
    }
}

func main() {
    name := "Hello World"

    fmt.Println( name )

    printBytes( name )
    fmt.Printf( "\n" )

    printChars( name )
    fmt.Printf( "\n" )

    name = "Señor"
    printChars( name )
    fmt.Printf( "\n" )
}
