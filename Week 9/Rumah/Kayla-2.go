package main

import "fmt"

func main() {
	var u1, u2, u3 int
	var h string

	fmt.Scan(&u1, &u2, &u3)
	if (u2 - u1) == (u3 - u2) {
		h = "ya"
	} else {
		h = "tidak"
	}
	fmt.Print(h)
}
