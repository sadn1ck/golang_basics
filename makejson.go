package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main5() {

	mapv := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)
	/*
		// Alternative input method is using NewScanner
		in = bufio.NewScanner(os.Stdin)
		// for name
		in.Scan()
		name = in.Text()
		// etc
	*/
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Enter your address: ")
	address, _ := reader.ReadString('\n')
	// Removing newline character from input string
	mapv["name"] = name[:len(name)-1]
	mapv["address"] = address[:len(address)-1]

	// Create JSON object from the map
	myjson, err := json.Marshal(mapv)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(myjson))
}
