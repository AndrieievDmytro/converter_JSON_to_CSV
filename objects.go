package main

type Session struct {
	Description  string   `json:"description"`
	Tags         []string `json:"tags"`
	Speakers     []string `json:"speakers"`
	Presentation string   `json:"presentation"`
	Title        string   `json:"title"`
	Complexity   string   `json:"complexity"`
	Language     string   `json:"language"`
}

type Speaker struct {
	Title          string   `json:"title"`
	ShortBio       string   `json:"shortBio"`
	Photo          string   `json:"photo"`
	Featured       bool     `json:"featured"`
	CompanyLogoURL string   `json:"companyLogoUrl"`
	Country        string   `json:"country"`
	Pronouns       string   `json:"pronouns"`
	Bio            string   `json:"bio"`
	Order          int      `json:"order"`
	Socials        []Social `json:"socials"`
	Name           string   `json:"name"`
	PhotoUrl       string   `json:"photoUrl"`
	CompLogo       string   `json:"companyLogo"`
	Company        string   `json:"company"`
	Badges         []Badge  `json:"badges"`
}

type Social struct {
	Link string `json:"link"`
	Icon string `json:"icon"`
	Name string `json:"name"`
}

type Badge struct {
	Link        string `json:"link"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type Schedule struct {
	Timeslots    []Timeslot `json:"timeslots"`
	DateReadable string     `json:"dateReadable"`
	Tracks       []Track    `json:"tracks"`
}

type Timeslot struct {
	StartTime string            `json:"startTime"`
	EndTime   string            `json:"endTime"`
	Sessions  []SessionSchedule `json:"sessions"`
}

type SessionSchedule struct {
	Items []string `json:"items"`
}

type Track struct {
	Title string `json:"title"`
}
