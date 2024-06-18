package main

import "fmt"

func main() {
    var i, j int

    for i = 1; i <= 99; i++ {
        j = j + i
        i = i + 1
    }
    fmt.Print(j)
}
