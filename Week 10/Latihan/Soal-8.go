package main

import "fmt"

func main() {
    var w1, w2 string
    
    fmt.Scan(&w1, &w2)
    if (w1 == "merah" && w2 == "biru") || (w1 == "biru" && w2 == "merah") {
        fmt.Print("ungu")
    } else if (w1 == "kuning" && w2 == "merah") || (w1 == "merah" && w2 == "kuning") {
        fmt.Print("orange")
    } else if (w1 == "kuning" && w2 == "biru") || (w1 == "biru" && w2 == "kuning") {
        fmt.Print("hijau")
    }
}