package main

import "fmt"

func main() {
    var t string
	var j, tsd, tsmp, tsma int

	fmt.Scan(&t, &j)
	for t != "x" && j != 0 {
		if t == "sd" {
			tsd = tsd + j
		} else if t == "smp" {
			tsmp = tsmp + j
		} else if t == "sma" {
			tsma = tsma + j
		}
		fmt.Scan(&t, &j)
	}
	fmt.Print(tsd, tsmp, tsma)
}	