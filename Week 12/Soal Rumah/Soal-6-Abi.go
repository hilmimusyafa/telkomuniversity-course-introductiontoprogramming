package main
import (
    "fmt"
    "math"
)

func main() {
    var i, n, m, a, x float64
    
    fmt.Scan(&n, &m)
    if n < -1 || m < -1 {
		for i = n; i <= m; i++ {
			if math.Mod(i , 2) == 0 {
				x = i
			} else if math.Mod(i , 2) != 0 {
				x = i * -1
			}
			a = math.Pow(2, x)
			fmt.Printf("%.2f ", a)
		}
	} else if n >= -1 || m >= -1 {
		for i = n; i <= m; i++ {
			if i >= -1 {
				if math.Mod(i , 2) == 0 {
					x = i * -1
				} else if math.Mod(i , 2) != 0 {
					x = i
				}
			}
			if i < -1 {
				if math.Mod(i , 2) == 0 {
					x = i
				} else if math.Mod(i , 2) != 0 {
					x = i * -1
				}

			}
			a = math.Pow(2, x)
			fmt.Printf("%.2f ", a)
		}
	}
}
	