package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func readCsv(path string) ([][]string, error) {
	var dataFile *os.File
	var err error

	dataFile, err = os.OpenFile(path, os.O_RDONLY, 0666)
	Check(err)
	defer dataFile.Close()
	if err == nil {
		var buf []byte
		var rerr error
		// Reading text from file
		buf, rerr = io.ReadAll(dataFile)
		Check(rerr)
		// Parsing from comma-separated text
		r := csv.NewReader(strings.NewReader(string(buf)))
		// if there are commas not semicolons comment out lines 41,42
		r.Comma = ';'
		r.LazyQuotes = true
		records, err := r.ReadAll()
		Check(err)
		return records, nil
	}
	return nil, err
}

func convertStrArrayToJson(records [][]string) string {
	// Converting from array of string to JSON
	jsonData := ""
	paramsLength := len(records[0]) - 1
	/*
			Replace "\n" and "\t" in description with spaces.
			If you need to add other characters to replace just add new values in the function  NewReplacer() input.
		 	Odd values are what you want to replace. Even what do you want to receive on output.
	*/
	var replacer = strings.NewReplacer("\n", " ", "\t", " ")

	for _, record := range records {
		wrongStr := false
		if len(record) < paramsLength || len(record) > paramsLength+1 {
			fmt.Println("Wrong parameters count")
			wrongStr = true
		}
		// CSV params
		name := record[0]
		description := replacer.Replace(record[1])
		tags := strings.Replace(record[2], ";", " \",\"", -1)
		speakers := strings.Replace(record[3], ";", "\",\"", -1)
		presentation := record[4]
		title := record[5]
		complexity := record[6]
		language := record[7]

		if !wrongStr {
			jsonData += "\"" + name + "\" : { \"description\": \"" + description + "\" ,  \"tags\": [ \"" + tags + "\" ],  \"speakers\": [ \"" + speakers + "\" ], \"presentation\": \"" + presentation + "\", \"title\": \"" + title + "\", \"complexity\": \"" + complexity + "\", \"language\": \"" + language + "\" },"
		}
	}
	jsonData = "{" + jsonData[:len(jsonData)-1] + "}"
	return jsonData
}

func writeToJson(records [][]string) {
	f, err := os.Create("sessions.json")

	Check(err)
	defer f.Close()

	_, err2 := f.WriteString(convertStrArrayToJson(records))

	Check(err2)
}
