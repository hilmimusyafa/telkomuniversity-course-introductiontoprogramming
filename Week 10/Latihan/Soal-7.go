package main 
import "fmt"

func main() {
    var v, s int
    
    fmt.Scan(&v, &s)
    if v * 4 > s {
        fmt.Print("viking menang")
    } else if v * 4 < s {
        fmt.Print("saxon menang")
    } else {
        fmt.Print("imbang")
    }
}