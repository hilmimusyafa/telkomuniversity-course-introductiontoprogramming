package main
import "fmt"

func main() {
    var a, m, j int
    var b string
    var t, i int
    
    fmt.Scan(&t)
    for i = 1 ; i <= t; i++ {
        fmt.Scan(&b)
        if b == "mangga"{
            m = m + 1
        }
        if b == "jeruk"{
            j = j + 1
        }
        if b == "apel"{
            a = a + 1
        }
    }
    fmt.Print(a, j, m)
}