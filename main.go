package main

import (
	"flag"
	"fmt"
)

var (
	input_file_session   = "input/sessions.json"
	input_file_speackers = "input/speakers.json"
	input_file_schedule  = "input/schedule.json"
	operation_type       string
	path                 string
)

func Check(e error) {
	if e != nil {
		fmt.Print(e)
	}
}

func init() {
	flag.StringVar(&operation_type, "t", operation_type, "Type of a file to convert.")
	flag.StringVar(&path, "p", path, "Path to file. Format : input/file_name.json/csv")
}

func main() {
	flag.Parse()
	convertJSONtoCSV(path)
	// records, err := readCsv(path)
	// writeToJson(records)
	// fmt.Print(err)
}
