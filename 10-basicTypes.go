package main

import (
	"fmt"
	"math/cmplx"
)

var (
	rahul bool       = false
	kumar int        = 23
	munda complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type of the variable is %T and its value is %v\n", rahul, rahul)
	fmt.Printf("Type of the variable is %T and its value is %v\n", kumar, kumar)
	fmt.Printf("Type of the variable is %T and its value is %v\n", munda, munda)
}
