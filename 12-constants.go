package main

import "fmt"

const pi = 3.14

func main() {
	const world = "Ni HAO"
	fmt.Println("Hello", world)
	fmt.Println("Happy", pi, "Day")

	const truth = true
	fmt.Println("Go rules?", truth)
}
