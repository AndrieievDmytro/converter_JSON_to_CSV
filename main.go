package main

import (
	"flag"
	"fmt"
	"unicode"
)

var (
	operation_type string
	path           string
	file_name      string
)

func Check(e error) {
	if e != nil {
		fmt.Print(e)
	}
}

func IsUpper(r rune) bool {
	if !unicode.IsUpper(r) && unicode.IsLetter(r) {
		return false
	}
	return true
}

func init() {
	flag.StringVar(&operation_type, "t", operation_type, "Type of a file to convert.")
	flag.StringVar(&path, "p", path, "Path to file. Format : input/file_name.json/csv")
	flag.StringVar(&file_name, "f", file_name, "Name of a file to convert.")
}

func main() {
	flag.Parse()
	switch operation_type {
	case "csv":
		convertCSVtoJSON(path)
	case "json":
		convertJSONtoCSV(path)
	}
}
