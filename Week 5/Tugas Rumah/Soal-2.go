package main

import "fmt"

func main() {
	var iter, n int
	var str string

	fmt.Scan(&n)
	fmt.Scan(&str)
	for iter = 1; iter <= n; iter++ {
		fmt.Println(str)
	}
}
