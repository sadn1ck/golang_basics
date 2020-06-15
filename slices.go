package main

import (
	"fmt"
)

func main() {
	fmt.Print("Slices\n")
	arr := [...]int{1, 2, 6, 3, 9, 10, 8, 11}
	s03 := arr[0:3]
	// s35 := arr[3:5]
	for i, elem := range s03 {
		fmt.Printf("%d %d\n", i, elem)
	}
	fmt.Println(cap(s03))
	fmt.Println("-------------------------")
	// slice is just a window,like a python shallow copy
	// change to slice changes array, as well as any other slice referecning the same element
	/*
		slice literals
		slice literal to create the array, slice points to start
		length is capacity
	*/
	sli := []int{1, 2, 3, 4}
	fmt.Println(len(sli))
	fmt.Println(cap(sli))
}
