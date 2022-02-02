package main

import (
    "fmt"
    "os"
	"regexp"
)

var re = regexp.MustCompile("([0-9]+)([A-Za-z]+)")

func main() {

	//Command-line Arguments
	args := os.Args[1:]
	if len(args) < 4 {
		matches := re.FindStringSubmatch(args[0])
		fmt.Println(matches[1])
		fmt.Println(matches[2])
		//fmt.Println(args[0])
	}
}
