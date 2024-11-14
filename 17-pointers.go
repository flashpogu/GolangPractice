package main

import "fmt"

func main() {
	i, j := 23, 69
	p := &i
	*p = 239
	fmt.Println(*p, j)
}
