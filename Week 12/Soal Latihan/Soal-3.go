package main

import "fmt"

func main() {
    var e, m, t int
	
	fmt.Scan(&e, &m)
	for e != 0 && m != 0 {
		if e % 2 != 0 {
			t = t + m
		}
		fmt.Scan(&e, &m)
	}
	fmt.Print(t)
}
	