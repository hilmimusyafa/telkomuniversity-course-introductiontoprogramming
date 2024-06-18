package main

import "fmt"

func main() {
	var n, m int

	fmt.Scan(&n, &m)
	n = n - 1
	for n < m {
		n = n + 1
		fmt.Println(n)
	}
}
