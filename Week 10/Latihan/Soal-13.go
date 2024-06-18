package main

import "fmt"

func main() {
    var r1, r2, r3, r4 int
    var r float64
    
    fmt.Scan(&r1, &r2, &r3, &r4)
    r = (float64(r1) + float64(r2) + float64(r3) + float64(r4)) / 4.0
    if r >= 3.50 {
        fmt.Print("bagus")
    } else if  r <= 1.50{
        fmt.Print("kurang")
    } else {
        fmt.Print("sedang")
    }
}