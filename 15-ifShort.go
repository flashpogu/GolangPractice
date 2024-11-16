package main

import (
	"fmt"
	"math"
)

func power(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v > lim {
		return v
	} else {
		return lim
	}
}
func main() {
	fmt.Println(
		power(3, 2, 5),
		power(9, 2, 90),
	)
}
