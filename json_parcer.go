package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createFile(file string) (*csv.Writer, *os.File) {
	//  Create a new file to store CSV data
	output_file, err := os.Create(file)
	if err != nil {
		Check(err)
	}
	//  Write the header of the CSV file and the successive rows by iterating through the JSON struct array
	writer := csv.NewWriter(output_file)
	return writer, output_file
}

func convertSessions(source_file *os.File) {
	session_headers := []string{"Session_id", "Description", "Tags", "Speakers", "Presentation", "Title", "Complexity", "Language"}
	session_csv_output := "output/sessions_converted.csv"
	var sessions map[int]Session
	if err := json.NewDecoder(source_file).Decode(&sessions); err != nil {
		Check(err)
	}

	writer, output := createFile(session_csv_output)

	defer output.Close()

	if err := writer.Write(session_headers); err != nil {
		Check(err)
	}

	for id, s := range sessions {
		tag := strings.Join(s.Tags, ";")
		speaker := strings.Join(s.Speakers, ";")
		var csv_row []string
		csv_row = append(csv_row, fmt.Sprint(id), s.Description, tag, speaker, s.Presentation, s.Title, s.Complexity, s.Language)
		if err := writer.Write(csv_row); err != nil {
			Check(err)
		}
	}
}

func convertSpeakers(source_file *os.File) {
	speakers_headers := []string{"Name", "Title", "ShortBio", "Photo", "Featured", "CompanyLogoUrl", "Country", "Pronouns", "Bio", "Order", "Socials", "Name", "PhotoUrl", "CompanyLogo", "Company", "Badges"}
	speakers_csv_ouput := "output/speakers_converted.csv"
	var speakers map[string]Speaker
	if err := json.NewDecoder(source_file).Decode(&speakers); err != nil {
		Check(err)
	}

	writer, output := createFile(speakers_csv_ouput)

	defer output.Close()

	if err := writer.Write(speakers_headers); err != nil {
		Check(err)
	}

	for name, s := range speakers {
		var socials []string
		var badges []string
		for _, social := range s.Socials {
			socials = append(socials, social.Icon, social.Link, social.Name)
		}
		for _, badge := range s.Badges {
			badges = append(badges, badge.Name, badge.Description, badge.Link)
		}
		badge := strings.Join(badges, ";")
		social := strings.Join(socials, ";")

		var csv_row []string
		csv_row = append(csv_row, name, s.Title, s.ShortBio, s.Photo, strconv.FormatBool(s.Featured), s.CompanyLogoURL, s.Country, s.Pronouns, s.Bio, fmt.Sprint(s.Order), social, s.Name, s.PhotoUrl, s.CompLogo, s.Company, badge)
		if err := writer.Write(csv_row); err != nil {
			Check(err)
		}
	}
}

func convertSchedule(source_file *os.File) {
	schedule_headers := []string{"Date", "Date Readable", "Title", "Timeslot Num", "Start time", "End time", "Items"}
	schedule_csv_ouput := "output/schedule_converted.csv"

	var schedules map[string]Schedule
	// Read the JSON file into the struct array
	if err := json.NewDecoder(source_file).Decode(&schedules); err != nil {
		Check(err)
	}

	writer, output := createFile(schedule_csv_ouput)

	defer output.Close()

	if err := writer.Write(schedule_headers); err != nil {
		Check(err)
	}
	for date, s := range schedules {
		var csv_row []string
		var tracks []string
		for _, track := range s.Tracks {
			tracks = append(tracks, track.Title)
		}
		track := strings.Join(tracks, ", ")
		csv_row = append(csv_row, date, s.DateReadable, track)
		if err := writer.Write(csv_row); err != nil {
			Check(err)
		}
		for i, timeslot := range s.Timeslots {
			var sessions []string
			defer writer.Flush()
			for _, session := range timeslot.Sessions {

				sessions = append(sessions, session.Items...)
			}
			session := strings.Join(sessions, ", ")
			csv_row = []string{}
			csv_row = append(csv_row, "", "", "", fmt.Sprint(i+1), timeslot.StartTime, timeslot.EndTime, session)
			if err := writer.Write(csv_row); err != nil {
				Check(err)
			}
		}
	}
}

func convertJSONtoCSV(path string) {
	source_file, err := os.Open(path)
	Check(err)
	defer source_file.Close()

	if strings.Contains(path, input_file_session) {
		convertSessions(source_file)
	} else if strings.Contains(path, input_file_speackers) {
		convertSpeakers(source_file)
	} else if strings.Contains(path, input_file_schedule) {
		convertSchedule(source_file)
	} else {
		fmt.Print("Provided input_file_* parameter in func convertJSONtoCSV() does not contain right file path")
	}
}
