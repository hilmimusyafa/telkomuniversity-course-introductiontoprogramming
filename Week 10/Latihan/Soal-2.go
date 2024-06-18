package main

import "fmt"

func main() {
    var p1, p2, p3, p4, p5 int

	fmt.Scan(&p1, &p2, &p3, &p4, &p5)
	if ((p1 + p2 + p3) == (p3 + p4 + p5)) && ((p2 + p3 + p4) % (p1 + p5) == 0) {
		fmt.Println("cashback")
		fmt.Println("diskon")
	} else if (p1 + p2 + p3) == (p3 + p4 + p5) {
		fmt.Print("cashback")
	} else if ((p2 + p3 + p4) % (p1 + p5) == 0) {
		fmt.Print("diskon")
	} else {
		fmt.Print(" ")
	}
}