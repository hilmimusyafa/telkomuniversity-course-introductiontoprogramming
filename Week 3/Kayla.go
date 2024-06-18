package main

import "fmt"

func main() {

	var x, y byte
	var hasil bool

	fmt.Scanf("%c %c", &x, &y)
	hasil = (x <= 123 && x == y || x+32 == y || x-32 == y || y+32 == x || y-32 == x)
	fmt.Println(hasil)
}
 