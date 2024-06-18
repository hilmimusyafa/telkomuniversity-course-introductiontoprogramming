package main

import "fmt"

func main() {
	var number1, number2 int
	var x1, x2 int
	var y1, y2 int
	var convertedDigit1, convertedDigit2 int

	fmt.Scan(&number1, &number2)
	x1 = number1 % 100
	x2 = number1 / 100

	y1 = number2 / 10
	y2 = number2 % 10

	convertedDigit1 = (y2 * 100) + x1
	convertedDigit2 = (y1 * 10) + x2

	fmt.Println(convertedDigit1, convertedDigit2)
}
