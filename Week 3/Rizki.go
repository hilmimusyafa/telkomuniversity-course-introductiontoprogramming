package main

import "fmt"

func main() {
	var j, m, d, x int
	var a int

	fmt.Scan(&x)
	j = x / 3600
	a = x % 3600

	m = a / 60
	d = a % 60

	fmt.Println(j, m, d)
}
