package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	arr := make([]int, 0, 3)

	var inp string
	for {
		fmt.Print("Enter a number (x to quit): ")
		_, err := fmt.Scan(&inp)
		// necessary since scan alr inputting to variable, no need to point variable again
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		checkExit := strings.ToLower(inp)
		if checkExit == "x" {
			fmt.Println("x entered, quitting.....")
			os.Exit(0)
		}
		intFromInp, _ := strconv.Atoi(inp)
		arr = append(arr, intFromInp)
		sort.Ints(arr)
		fmt.Println(arr)
	}
}
