package main

import "fmt"

func main() {
    var n, s, x, tps int
    var s1, s2 int
    
    fmt.Scan(&n)
    for tps < n {
        for x <= 2 {
            fmt.Scan(&s)
            s1 = s1 + s
            x = x + 1
        } 
        for x <= 5 {
            fmt.Scan(&s)
            s2 = s2 + s
            x = x + 1
        }
		tps = tps + 1
		x = 0
    }
    if s1 > s2 {
        fmt.Print(s1, s2, " ", "A")
    } else if s1 < s2 {
        fmt.Print(s1, s2, " ", "B")
    } else {
        fmt.Print(s1, s2, 0 )
    }
}