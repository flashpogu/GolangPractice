package main

import "fmt"

func addInt(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("Value is: ", addInt(22, 22))
}
