package main

import "fmt"

func main() {
    var k string
	var t, tm, tb int

	fmt.Scan(&k, &t)
	for k != "x" && t != 0 {
		if k == "m" {
			tm = tm + t
		} else if k == "b" {
			tb = tb + t
		}
		fmt.Scan(&k, &t)
	}
	fmt.Print(tm, tb)
}	