package main

import "fmt"

func main() {
	var adik, kakak int
	var h bool

	fmt.Scan(&adik, &kakak)
	h = (kakak == adik || kakak == adik+1 || kakak == adik-1 || kakak == adik+5 || kakak == adik-5)
	fmt.Println(h)
}
