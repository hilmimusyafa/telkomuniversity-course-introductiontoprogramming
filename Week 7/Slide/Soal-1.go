package main

import "fmt"

func main() {
    var k byte
    var x bool

    fmt.Scanf("%c", &k)
    x = (k >= 48 && k <= 56) || (k >= 65 && k <= 90) || (k >= 97 && k <= 122)
    fmt.Print(x)
}