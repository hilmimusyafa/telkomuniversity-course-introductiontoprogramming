package main

import "fmt"

func main() {
	var n, ta, tb int
	var iter int

	fmt.Scan(&ta, &n)
	for iter = 1; iter <= n; iter++ {
		tb += ta
		ta += 2500
	}
	fmt.Println(tb)
}
