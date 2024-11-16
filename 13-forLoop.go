package main

import "fmt"

func main() {
	var sum int = 0
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		sum += i
	}
	fmt.Println("Hello World! LMao", sum)
}
