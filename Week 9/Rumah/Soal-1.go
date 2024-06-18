package main

import (
    "fmt"
    "math"
)

func main() {
    var x1, x2, x3 float64
    
    fmt.Scan(&x1, &x2, &x3)
    if math.Pow(x1, x2) == x3 {
        fmt.Print("benar")
    }else {
        fmt.Print("salah")
    } 
}