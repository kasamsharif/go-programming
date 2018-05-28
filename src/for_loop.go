package main

import "fmt"

// init and post statement are optional

func main() {
    sum := 0
    for i:=0; i<10; i++ {
        sum += i
    }
    fmt.Println(sum)
}
