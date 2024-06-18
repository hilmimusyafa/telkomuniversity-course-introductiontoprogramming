package main

import "fmt"

func main() {
	var u1, u2, u3 int
	fmt.Scan(&u1, &u2, &u3)
	if u1 % 2 == 0 && u2 % 2 == 0 && u3 % 2 == 0 {
	    fmt.Print("genap")
	} else if u1 % 2 == 1 && u2 % 2 == 1 && u3 % 2 == 1 {
	    fmt.Print("ganjil")
	} else {
	    fmt.Print("bukan ganjil atau genap")
	}
}