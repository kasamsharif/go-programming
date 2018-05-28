package main

import "fmt"

// defer statements are executed in order last in first out

func main() {
    for i:=0; i<10; i++ {
        defer fmt.Println(i)
    }
}
