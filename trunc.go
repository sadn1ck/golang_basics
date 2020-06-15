package main

import (
	"fmt"
	"math"
)

func main2() {
	fmt.Println("Enter a floating point number: ")
	var inp float64
	fmt.Scanf("%f", &inp)
	inp = math.Trunc(inp)
	fmt.Println(inp)
}
