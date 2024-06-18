package main

import "fmt"

func main() {
    var Temp int
    fmt.Scan(&Temp)
    if Temp < 0 {
        fmt.Print("Freezing weather")
    } else if Temp >= 0 && Temp <= 9{
        fmt.Print("Very Cold weather")
    } else if Temp >= 10 && Temp <= 19{
        fmt.Print("Cold weather")
    } else if Temp >= 20 && Temp <= 29{
        fmt.Print("Normal in Temp")
    } else if Temp >= 30 && Temp <= 39 {
        fmt.Print("It's Hot")
    } else {
        fmt.Print("It's Very Hot")
    }
}