package main

import "fmt"

func main() {
	var i string
	var k, t int

	fmt.Scan(&i)
	for i[k] != '.' {
    	if i[k] == 'a' || i[k] == 'i' || i[k] == 'u' || i[k] == 'e' || i[k] == 'o'{
        	t = t + 1
    	}
    	k = k + 1
	}
    fmt.Print(t)
}