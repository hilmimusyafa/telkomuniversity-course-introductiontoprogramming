package main
import "fmt"

func main() {
    var hasil, iter int
    
    for iter = 2; iter<=50; iter++{
        hasil = hasil + iter
    }
    fmt.Println(hasil)
}
