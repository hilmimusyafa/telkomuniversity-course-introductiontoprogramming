package main

import "fmt"

func main() {
	var n, m, x, y int
	var t int

	fmt.Scan(&n, &m, &x, &y)
	for n >= x && m >= y {
		n = n - x
		m = m - y
		//fmt.Println(n, ">", x)
		//fmt.Println(m, ">", y)
		t += 1
	}
	fmt.Println(t)
}
