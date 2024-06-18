package main

import "fmt"

func main() {
    var a, b, c, max int
    fmt.Scan(&a, &b, &c)
    
    max = a
    if a != b && b != c && a != c{
        if b > a && b > c{
            max = b
        } else if c > a && c > b{
            max = c
        }
    } else if a == b || b == c{
        if a == b && b > c {
            max = b
        } else if a == b && b < c {
            max = c
        } else if b == c && a > b {
            max = a
        } else if b == c && a < b {
            max = b 
        } else if a == c && a > b {
            max = a
        } else if a == c && a < b{
            max = b
        }
    }
    fmt.Print(max)
}