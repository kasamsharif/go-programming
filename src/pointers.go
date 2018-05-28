package main

import "fmt"

func main() {
    x, y := 21, 16

    i := &x

    fmt.Println(*i) //x value

    *i = 41 // changing x value

    fmt.Println(x) // x value is now 41

    i = &y
    *i = *i / 4 // changing y 
    fmt.Println(y)
}
