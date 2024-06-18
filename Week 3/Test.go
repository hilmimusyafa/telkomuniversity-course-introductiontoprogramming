package main

import "fmt"

func main() {
	var d1, d2, d3, d4 int
	var inp int
	var diskon, cashback, voucher bool

	fmt.Scan(&inp)
	d1 = inp / 1000
	inp = inp % 1000
	d2 = inp / 100
	inp = inp % 100
	d3 = inp / 10
	inp = inp % 10
	d4 = inp

	diskon = ((d2*10)+d3)%2 == 0
	cashback = (d1+d3)%d4 == 0
	voucher = (diskon || cashback) && !(diskon && cashback)

	fmt.Println(diskon, cashback, voucher)
}
