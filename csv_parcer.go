package main

import (
	"encoding/csv"
	"encoding/json"
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
		records, err := r.ReadAll()
		Check(err)

		return records, nil
	}
	return nil, err
}

func convertSessionsToJSON(records [][]string) string {
	jsonData := "{"
	var column_names []string
	var replacer = strings.NewReplacer("\n", " ", "\t", " ")

	for i, record := range records {

		if i == 0 {
			column_names = append(column_names, record...)
			continue
		}

		for j, value := range record {
			if column_names[j] == "Session_id" {
				jsonData += "\"" + value + "\":{"
				continue
			}
			if column_names[j] == "Description" {
				value = replacer.Replace(record[j])
			}
			if column_names[j] == "Tags" {
				value = strings.Replace(record[j], ";", " \",\"", -1)
				jsonData += "\"" + column_names[j] + "\": [\"" + value + "\"],"
				continue
			}
			if column_names[j] == "Speakers" {
				value = strings.Replace(record[j], ";", " \",\"", -1)
				jsonData += "\"" + column_names[j] + "\": [\"" + value + "\"],"
				continue
			}
			if j != len(record)-1 {
				jsonData += "\"" + column_names[j] + "\":\"" + value + "\","
			} else if i != len(records)-1 {
				jsonData += "\"" + column_names[j] + "\":\"" + value + "\"},"
			} else {
				jsonData += "\"" + column_names[j] + "\":\"" + value + "\"}"
			}
		}
	}
	jsonData += "}"
	return jsonData
}

func convertSpeakersToJSON(records [][]string) string {
	jsonData := "{"
	var replacer = strings.NewReplacer("\n", " ", "\t", " ")
	var column_names []string

	for i, record := range records {
		// var str string
		if i == 0 {
			column_names = append(column_names, record...)
			continue
		}
		var social_records []Social
		var badges_records []Badge
		for j, value := range record {
			if column_names[j] == "Name_tag" {
				jsonData += "\"" + value + "\":{"
				continue
			}
			if column_names[j] == "Bio" {
				value = replacer.Replace(value)
			}
			if column_names[j] == "Badges" {
				if len(records[i][j]) == 0 {
					b := &Badge{}
					badges_records = append(badges_records, *b)
					badges_json, err := json.Marshal(badges_records)
					value = string(badges_json)
					Check(err)
					if i != len(records)-1 && j == len(record)-1 {
						jsonData += "\"" + column_names[j] + "\":" + value + "},"
					} else if i == len(records)-1 && j == len(record)-1 {
						jsonData += "\"" + column_names[j] + "\":" + value + "}"
					} else {
						jsonData += "\"" + column_names[j] + "\":" + value + ","
					}
					continue
				}
				badges := strings.Split(records[i][j], ";")
				for i := 0; i < len(badges)-1; i += 3 {
					b := &Badge{Name: badges[0], Link: badges[2], Description: badges[1]}
					badges_records = append(badges_records, *b)
					badges_json, err := json.Marshal(badges_records)
					Check(err)
					value = string(badges_json)
				}
				if i != len(records)-1 && j == len(record)-1 { // Value is the last column in a raw
					jsonData += "\"" + column_names[j] + "\":" + value + "},"
				} else if i == len(records)-1 && j == len(record)-1 { // Last raw and last column
					jsonData += "\"" + column_names[j] + "\":" + value + "}"
				} else {
					jsonData += "\"" + column_names[j] + "\":" + value + ","
				}
				continue
			}
			if column_names[j] == "Socials" {
				if len(records[i][j]) == 0 {
					s := &Social{}
					social_records = append(social_records, *s)
					socials_json, err := json.Marshal(social_records)
					Check(err)
					value = string(socials_json)
					jsonData += "\"" + column_names[j] + "\":" + value + ","
					continue
				}
				socials := strings.Split(records[i][j], ";")
				for i := 0; i < len(socials)-1; i += 3 {

					s := &Social{Link: socials[i+1], Icon: socials[i], Name: socials[i+2]}
					social_records = append(social_records, *s)
					socials_json, err := json.Marshal(social_records)
					Check(err)
					value = string(socials_json)
				}
				jsonData += "\"" + column_names[j] + "\":" + value + ","
				continue
			}
			if j != len(record)-1 {
				jsonData += "\"" + column_names[j] + "\":\"" + value + "\","
			} else if i != len(records)-1 {
				jsonData += "\"" + column_names[j] + "\":\"" + value + "\"},"
			} else {
				jsonData += "\"" + column_names[j] + "\":\"" + value + "\"}"
			}
		}
	}
	jsonData += "}"
	return jsonData
}

func convertCSVtoJSON(path string) {
	records, err0 := readCsv(path)
	Check(err0)
	if file_name == "speakers" {
		f, err1 := os.Create("output/speakers_converted.json")
		Check(err1)
		defer f.Close()
		_, err2 := f.WriteString(convertSpeakersToJSON(records))
		Check(err2)
	} else if file_name == "sessions" {
		f, err1 := os.Create("output/sessions_converted.json")
		Check(err1)
		defer f.Close()
		_, err2 := f.WriteString(convertSessionsToJSON(records))
		Check(err2)
	} else {
		fmt.Print("Provided input_file_* parameter in func writeToJson() does not contain right file path")
	}
}
