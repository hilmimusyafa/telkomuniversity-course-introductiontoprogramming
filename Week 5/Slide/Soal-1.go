package main

import "fmt"

func main() {
	var u int
	var x, b string
	var a, p string
	var stop bool

	u = 0
	x = "admin"
	b = "Login berhasil"
	stop = false
	for !stop {
		fmt.Scan(&a, &p)
		stop = (a == x && p == x)
		u = u + 1
	}
	fmt.Print(u-1, " ", b)
}
