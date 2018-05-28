package main

import "fmt"

func main() {
    defer fmt.Println("Defer function is called, untill the surrounding functions returns")
    fmt.Println("Execute me first!!")
}
