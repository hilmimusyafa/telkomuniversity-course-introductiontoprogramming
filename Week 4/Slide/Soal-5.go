package main

import "fmt"

func main() {
	var iter int
	var n int
	var d int
	var s bool

	fmt.Scan(&n)
	for iter = 1; iter <= n; iter++ {
		d = (n % iter)
		s = (d == 0)
		fmt.Println(iter, s)
	}
}
