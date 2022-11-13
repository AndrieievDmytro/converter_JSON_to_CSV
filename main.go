package main

import "fmt"

var (
	paths = []string{"input/sessions_test.json", "input/speakers_test.json", "input/schedule_test.json"}
)

func Check(e error) {
	if e != nil {
		fmt.Print(e)
	}
}

func main() {
	convertJSONtoCSV(paths)
	// records, err := readCsv(path)
	// writeToJson(records)
	// fmt.Print(err)
}
