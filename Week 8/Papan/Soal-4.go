package main
import "fmt"

func main() {
	var n int
    
    fmt.Scan(&n)
	if n >= 0 && n <= 31 {
		if n == 0 {
		    fmt.Print("No Good")
		}
		if n >= 1 && n <= 15 {
		    fmt.Print("Decent")
	    }
	    if n >= 16 && n <= 25 {
		    fmt.Print("Decent")
	    }
	    if n >= 26 && n <= 29 {
		    fmt.Print("Very Good")
	    }
		if n == 30{
		    fmt.Print("Fantastic")
	    }
	    if n == 31{
	        fmt.Print("Best")
	    }
    }
}