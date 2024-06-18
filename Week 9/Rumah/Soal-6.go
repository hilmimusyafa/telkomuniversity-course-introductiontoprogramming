package main

import "fmt"

func main() {
	var a, b, c int
	var max, mid, min int

	fmt.Scan(&a, &b, &c)
	if a != b && a != c && b != c{
	    max = a
	    if b > a && b > c{
		max = b
	    }else if c > a && c > b{
		max = c
	    }
	    mid = b
	    if (a < b && a > c) || (c > a && b < a) {
		    mid = a
	    }else if (c < b && c > a) || (a > c && b < c) {
		    mid = c
	    }
	    min = c
	    if b < a && b < c{
		    min = b
	    }else if a < b && a < c{
		    min = a
	    }
	} else {
	    min = a
	    mid = b
	    max = c
	    if a == b && a < c{
	        min = a
	        mid = b
	        max = c
	    } else if a == b && a > c {
	        min = c
	        mid = a
	        max = b
        } else if b == c && a < b {
            min = a
            mid = b
            max = c
        } else if b == c && a > b {
            min = b
            mid = c
            max = a
        } else if a == c && a < b {
            min = a
            mid = c
            max = b
        } else if a == c && a > b {
            min = b
            mid = a
            max = c
	    }
	}
	fmt.Print(min, mid, max)
}