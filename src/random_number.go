package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main(){
    rand.Seed(time.Now().UnixNano())
    fmt.Println("Random number generated is ", rand.Intn(100))
}
