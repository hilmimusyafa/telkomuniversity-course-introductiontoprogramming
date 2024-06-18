	package main

	import "fmt"

	func main() {
		var a, b, x1, x2, x3, x4 int
		
		fmt.Scan(&x1, &x2, &x3, &x4)
		if x1 != x2 && x2 != x3 && x3 != x4{
			if x1 > x2 && x1 > x3 && x1 > x4 && x2 < x1 && x2 < x3 && x2 < x4{
				a = x1
				b = x2
			}else if x1 > x2 && x1 > x3 && x1 > x4 && x3 < x1 && x3 < x2 && x2 < x3{
				a = x1
				b = x3
			}else if x1 > x2 && x1 > x3 && x1 > x4 && x4 < x1 && x4 < x3 && x4 < x3{
				a = x1
				b = x4
			}else if x2 > x1 && x2 > x3 && x2 > x4 && x1 < x2 && x1 < x3 && x1 < x4{
				a = x2
				b = x1
			}else if x2 > x1 && x2 > x3 && x2 > x4 && x3 < x1 && x3 < x2 && x3 < x4{
				a = x2
				b = x3
			}else if x2 > x1 && x2 > x3 && x2 > x4 && x4 < x1 && x4 < x2 && x4 < x3{
				a = x2
				b = x4
			}else if x3 > x1 && x3 > x2 && x3 > x4 && x1 < x2 && x1 < x3 && x1 < x4{
				a = x3
				b = x1
			}else if x3 > x1 && x3 > x2 && x3 > x4 && x2 < x1 && x2 < x3 && x2 < x4{
				a = x3
				b = x2
			}else if x3 > x1 && x3 > x2 && x3 > x4 && x4 < x1 && x4 < x2 && x4 < x3{
				a = x3
				b = x4
			}else if x4 > x1 && x4 > x2 && x4 > x3 && x1 < x2 && x1 < x3 && x1 < x4{
				a = x4
				b = x1
			}else if x4 > x1 && x4 > x2 && x4 > x3 && x2 < x1 && x2 < x3 && x2 < x4{
				a = x4
				b = x2
			}else{
				a = x4
				b = x3
			}
		}else{
			a = x1
			b = x4
		}
		fmt.Print(a, b)
	}