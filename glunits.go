package main

import (
    "fmt"
    "os"
    "regexp"
    "strings"
    "strconv"
)

type dataC struct {
	from string
	to   string
	val  string
}

var c = []dataC{
	// Length
	{"yard", "foot", "3"},
	{"yard", "centimeter", "91.44"},
	{"yard", "millimeter", "914"},
	{"yard", "micrometer", "914400"},
	{"yard", "nanometer", "914400000"},
	{"yard", "inch", "36"},
	
	{"foot", "centimeter", "30.48"},
	{"foot", "millimeter", "305"},
	{"foot", "micrometer", "304800"},
	{"foot", "nanometer", "304800000"},
	{"foot", "inch", "12"},
	
	{"mile", "foot", "5280"},
	{"mile", "kilometer", "1.609"},
	{"mile", "meter", "1609"},
	{"mile", "centimeter", "160934"},
	{"mile", "millimeter", "1609000"},
	{"mile", "micrometer", "1609000000"},
	{"mile", "nanometer", "1609000000000"},
	{"mile", "yard", "1760"},
	{"mile", "inch", "63360"},
	
	{"nautical mile", "foot", "6076"},
	{"nautical mile", "kilometer", "1.852"},
	{"nautical mile", "meter", "1852"},
	{"nautical mile", "centimeter", "185200"},
	{"nautical mile", "millimeter", "1852000"},
	{"nautical mile", "micrometer", "1852000000"},
	{"nautical mile", "nanometer", "1852000000000"},
	{"nautical mile", "mile", "1.151"},
	{"nautical mile", "yard", "2025"},
	
	{"kilometer", "meter", "1000"},
	{"kilometer", "foot", "3281"},
	{"kilometer", "centimeter", "100000"},
	{"kilometer", "millimeter", "1000000"},
	{"kilometer", "micrometer", "1000000000"},
	{"kilometer", "nanometer", "1000000000000"},
	{"kilometer", "yard", "1094"},
	{"kilometer", "inch", "39370"},
	
	{"meter", "foot", "3.281"},
	{"meter", "centimeter", "100"},
	{"meter", "millimeter", "1000"},
	{"meter", "micrometer", "1000000"},
	{"meter", "nanometer", "1000000000"},
	{"meter", "yard", "1.094"},
	{"meter", "inch", "39.37"},
	
	{"centimeter", "millimeter", "10"},
	{"centimeter", "micrometer", "10000"},
	{"centimeter", "nanometer", "10000000"},
	
	{"inch", "centimeter", "2.54"},
	{"inch", "millimeter", "25.4"},
	{"inch", "micrometer", "25400"},
	{"inch", "nanometer", "25400000"},
	
	{"millimeter", "micrometer", "1000"},
	{"millimeter", "nanometer", "1000000"},
	
	{"micrometer", "nanometer", "1000"},
}

#reg
var re = regexp.MustCompile("([0-9.]+)([a-z/]+)")

func lowercase(array []string) {
	for i := 0; i < len(array); i++ {
		array[i] = strings.ToLower(array[i])
	}
	return
}


func main() {

	// Command-line Arguments
	args := os.Args[1:]
	
	// Pass to lower case
	lowercase(args)

	// Check the number of arguments
	if len(args) == 3 {
		matches := re.FindStringSubmatch(args[0])
		if len(matches) > 1 {
			// Number match!
			// matches[1]: is de number; matches[2]: from; args[1]: in or to; args[2]: final unit 
			mord(matches[2], args[2], matches[1])
		} else {
			fmt.Println("The first argument is not a number")
		}
	}
}

func stringTofloat(strVarf string) float64 {
	floatVar, err := strconv.ParseFloat(strVarf, 8)
	if err != nil {
		fmt.Println("Error, in try parser to float")
	}
	return floatVar
}


func mord(from string, to string, value string) {
	lengthStruct := len(c)-1
	for i, cNames := range c {
		if cNames.from == from && cNames.to == to {
			fmt.Println(stringTofloat(value) * stringTofloat(cNames.val), to, i)
			i = i + 1
			break
		} else if cNames.from == to && cNames.to == from {
			fmt.Println(stringTofloat(value) / stringTofloat(cNames.val), to, i)
			i = i + 1
			break
		} else if i == lengthStruct {
			fmt.Println("Unknown unit. From:", from, "To:", to)
		}
		
	}
	return
}
