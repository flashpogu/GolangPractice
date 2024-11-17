package main

import "fmt"

func main() {
	var temp = []float64{
		-28.0, 32.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	set := make(map[float64]bool)
	for _, t := range temp {
		set[t] = true
	}

	if set[-28.0] {
		fmt.Println("set member")
	}
	fmt.Println(set)

}
