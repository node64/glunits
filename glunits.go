package main

import (
    "fmt"
    "os"
)

func main() {

	//Command-line Arguments
	args := os.Args[1:]
	if len(args) < 4 {
		fmt.Println(args)
	}
}
