package main

import "fmt"

func main() {
    var x, y, z int
    fmt.Scan(&x, &y)
    if x % 2 == 0 && y % 2 == 0 {
        z = x * y
        fmt.Print(z)
    } else if x % 2 == 1 && y % 2 == 1 {
        z = x + y 
        fmt.Print(z)
    } else {
        z = 0
        fmt.Print(z)
    }
}