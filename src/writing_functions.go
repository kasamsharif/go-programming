package main

import (
    "fmt"
)

func add(a int, b int) int {
    return a + b
}

func subtract(a,b int) int {
    return a - b
}

func main() {
    fmt.Println(add(11,23))
    fmt.Println(subtract(61,43))
}
