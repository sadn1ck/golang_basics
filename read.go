package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

// since max len is 20, getting max len of each first name and last name
func getMaxLen(s string) string {
	runes := []rune(s)
	return string(runes[0:20])
}

func main() {
	var filename string
	nameSli := make([]name, 0)
	var nameObj name
	// nameobj is the object holding the read names
	fmt.Print("Enter file name: ")
	fmt.Scan(&filename)
	file, err := os.Open(filename)
	// opening file
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	// scanning line by line
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		// splitting line into space separated values
		if len(s[0]) > 20 {
			s[0] = getMaxLen(s[0])
		}
		if len(s[1]) > 20 {
			s[1] = getMaxLen(s[1])
		}
		// fmt.Println(s)
		nameObj.fname, nameObj.lname = s[0], s[1]
		// setting nameobj values from split strings
		nameSli = append(nameSli, nameObj)
	}

	for _, v := range nameSli {
		fmt.Println(v.fname, v.lname)
	}
	file.Close()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
