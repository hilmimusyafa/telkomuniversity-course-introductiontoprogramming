package main
import (
    "fmt"
    "math"
)

func main() {
    var i, n, m, t float64
    
    fmt.Scan(&n, &m)
    if n < m {
        for i = n ; i <= m ; i++ {
            if i < 0 {
				if math.Mod(-i, 2) == 1 {
					t = math.Pow(2, i)
					fmt.Printf("%0.2f ", t)
				}
				if math.Mod(i, 2) == 0 {
					t = math.Pow(2, -i)
					fmt.Printf("%0.2f ", t)
				}
			} else {
				if math.Mod(i, 2) == 1 {
					t = math.Pow(2, i)
					fmt.Printf("%0.2f ", t)
				}
				if math.Mod(i, 2) == 0 {
					t = math.Pow(2, i * -1)
					fmt.Printf("%0.2f ", t)
				}
			}
        }    
    } else {
        fmt.Print(" ")
    }
}