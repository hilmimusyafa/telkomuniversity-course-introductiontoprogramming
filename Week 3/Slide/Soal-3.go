package main

import "fmt"

func main() {
	var k string
	//var h bool

	fmt.Scan(&k)
	//fmt.Println(k)

	h := (k >= string('A') && k <= string('Z'))
	fmt.Println(h)
}
