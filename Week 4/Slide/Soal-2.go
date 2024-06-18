package main

import "fmt"

func main() {
	var iter int
	var n, l, k int
	var s int

	fmt.Scan(&n)
	for iter = 1; iter <= n; iter++ {
		fmt.Scan(&s)
		l = s * s
		k = 4 * s
		fmt.Println(l, k)
	}
}
