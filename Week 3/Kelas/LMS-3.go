package main

import "fmt"

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)
	a, b, c, d = b, c, d, a
	fmt.Println(a, b, c, d)
}
