package main

import "fmt"

func main() {
    var f, p, c bool

	fmt.Scan(&f, &p, &c)
	if f == true && p == true && c == true {
		fmt.Print("berkemah")
	} else {
		fmt.Print("tidak berkemah")
	}
}