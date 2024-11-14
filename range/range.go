package main

import "fmt"

var pow = []string{"r", "a", "h", "u", "l"}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
