package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main3() {
	reader := bufio.NewReader(os.Stdin)
	// create reading buffer like in java
	fmt.Printf("Enter a string: ")
	s, _ := reader.ReadString('\n')
	// _ is for disregarding err

	str := strings.ToLower(s)
	// converted to lower

	b := strings.ContainsAny(str, "a")
	// checks if str contains any a

	if str[0] == 'i' && str[len(str)-2] == 'n' && b {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}

}
