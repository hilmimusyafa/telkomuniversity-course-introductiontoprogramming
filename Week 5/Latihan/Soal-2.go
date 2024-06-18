package main
import "fmt"

func main() {
    var hasil, iter int
    
    for iter = 1; iter<=1000; iter++{
        hasil = hasil + iter
    }
    fmt.Println(hasil)
}
