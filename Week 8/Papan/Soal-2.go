package main

import "fmt"

func main() {
	var h int
    var d float64
        
    fmt.Scan(&h)
	if h >= 275000 {
		d = float64(h) * 5 / 100
	}
	d = float64(h) - d
	fmt.Print(d)
}