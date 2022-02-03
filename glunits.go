package main

import (
    "fmt"
    "os"
    "regexp"
    "strings"
)

var re = regexp.MustCompile("([0-9.]+)([a-z/]+)")

func lowercase(array []string) {
	for i := 0; i < len(array); i++ {
		array[i] = strings.ToLower(array[i])
	}
	return
}


func main() {

	//Command-line Arguments
	args := os.Args[1:]
	
	//Pass to lower case
	lowercase(args)

	//Check the number of arguments
	if len(args) < 4 {
		matches := re.FindStringSubmatch(args[0])
		fmt.Println(matches[1])
		fmt.Println(matches[2])
		//fmt.Println(args[0])
	}
}
