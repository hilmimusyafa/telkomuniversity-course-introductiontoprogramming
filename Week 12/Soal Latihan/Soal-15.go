package main

import "fmt"

func main() {
	var i string
	var k, tv, tk int

	fmt.Scan(&i)
	for i[k] != '.' {
    	if i[k] == 'a' || i[k] == 'i' || i[k] == 'u' || i[k] == 'e' || i[k] == 'o'{
        	tv = tv + 1
    	}
		if i[k] != 'a' && i[k] != 'i' && i[k] != 'u' && i[k] != 'e' && i[k] != 'o' {
        	tk = tk + 1
    	}
    	k = k + 1
	}
    fmt.Print(tv, tk)
}