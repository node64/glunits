package main

import (
    "fmt"
    "os"
    "regexp"
    "strings"
	"strconv"
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
		if len(matches) > 1 {
			fmt.Println(stringTofloat(matches[1])*12)
		} else {
			fmt.Println("The first argument is not a number")
		}
	}
}
/*
func convert(from, to, val, valt) {
	from * to
}*/


func stringTofloat(strVarf string) float64 {
	floatVar, err := strconv.ParseFloat(strVarf, 8)
	if err != nil {
		fmt.Println("Error, in try parser to float")
	}
	return floatVar
}
