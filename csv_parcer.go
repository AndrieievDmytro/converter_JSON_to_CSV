package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
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
		if path == "input/sessions.csv" {
			r.Comma = ';'
			r.LazyQuotes = true
		}
		records, err := r.ReadAll()
		Check(err)

		return records, nil
	}
	return nil, err
}

func convertSessionsToJSON(records [][]string) string {
	// Create json file to save records

	// Converting from array of string to JSON
	jsonData := ""
	paramsLength := len(records[0]) - 1
	/*
			Replace "\n" and "\t" in description with spaces.
			If you need to add other characters to replace just add new values in the function  NewReplacer() input.
		 	Odd values are what you want to replace. Even what do you want to receive on output.
	*/
	var replacer = strings.NewReplacer("\n", " ", "\t", " ")

	for i, record := range records {
		if i == 0 {
			continue
		}
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
			jsonData += "\"" + name + "\" : { \"description\": \"" + description + "\" ,  \"tags\": [ \"" +
				tags + "\" ],  \"speakers\": [ \"" + speakers + "\" ], \"presentation\": \"" +
				presentation + "\", \"title\": \"" + title + "\", \"complexity\": \"" + complexity + "\", \"language\": \"" + language + "\" },"
		}
	}
	jsonData = "{" + jsonData[:len(jsonData)-1] + "}"
	return jsonData
}

func convertSpeakersToJSON(records [][]string) string {
	jsonData := ""
	paramsLength := len(records[0]) - 1
	var replacer = strings.NewReplacer("\n", " ", "\t", " ")
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)
	for i, record := range records {
		if i == 0 {
			continue
		}
		var social_records []Social
		var badges_records []Badge
		wrongStr := false
		if len(record) < paramsLength || len(record) > paramsLength+1 {
			fmt.Println("Wrong parameters count")
			wrongStr = true
		}
		name_tag := record[0]
		title := record[1]
		short_bio := record[2]
		photo := record[3]
		featured := record[4]
		cmp_logo_url := record[5]
		country := record[6]
		pronouns := record[7]
		bio := replacer.Replace(record[8])
		order := record[9]
		socials := strings.Fields(strings.Replace(record[10], ";", " ", -1))
		name := record[11]
		photo_url := record[12]
		cmp_logo := record[13]
		company := record[14]
		badge := strings.Replace(record[15], " ", "", -1)
		badges := strings.Fields(strings.Replace(badge, ";", " ", -1))

		for i := 0; i < len(socials)-1; i += 3 {
			s := &Social{Link: socials[i+1], Icon: socials[i], Name: socials[i+2]}
			social_records = append(social_records, *s)
		}
		socials_json, err := json.Marshal(social_records)
		Check(err)

		for i := 0; i < len(badges)-1; i += 3 {
			var bages_description string
			submatchall := re.FindAllString(badges[i+1], -1)
			for _, element := range submatchall {
				bages_description = bages_description + " " + element
			}

			b := &Badge{Name: badges[0], Link: badges[2], Description: bages_description}
			badges_records = append(badges_records, *b)
		}
		badges_json, err1 := json.Marshal(badges_records)
		Check(err1)

		if !wrongStr {
			jsonData += "\"" + name_tag + "\" : { \"title\": \"" + title + "\" ,  \"socials\":" + string(socials_json) +
				",\"shortBio\": \"" + short_bio + "\" , \"photo\": \"" + photo + "\", \"featured\":" + featured + ", \"companyLogoUrl\": \"" +
				cmp_logo_url + "\", \"country\": \"" + country + "\" , \"pronouns\": \"" + pronouns + "\", \"bio\": \"" + bio + "\", \"order\": \"" +
				order + "\",  \"name\": \"" + name + "\", \"photoUrl\": \"" + photo_url + "\",  \"companyLogo\": \"" +
				cmp_logo + "\", \"company\": \"" + company + "\", \"badges\" :" + string(badges_json) + "},"
		}
	}
	jsonData = "{" + jsonData[:len(jsonData)-1] + "}"
	return jsonData
}

func convertCSVtoJSON(path string) {
	records, err0 := readCsv(path)
	Check(err0)
	if strings.Contains(path, "input/speakers.csv") {
		f, err1 := os.Create("output/speakers_converted.json")
		Check(err1)
		defer f.Close()
		_, err2 := f.WriteString(convertSpeakersToJSON(records))
		Check(err2)
	} else if strings.Contains(path, "input/sessions.csv") {
		f, err1 := os.Create("output/sessions_converted.json")
		Check(err1)
		defer f.Close()
		_, err2 := f.WriteString(convertSessionsToJSON(records))
		Check(err2)
	} else {
		fmt.Print("Provided input_file_* parameter in func writeToJson() does not contain right file path")
	}
}
