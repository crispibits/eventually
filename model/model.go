package model

import (
	"encoding/json"
	"os"
)

type Event struct {
	Id           int         `json:"id"`
	Title        string      `json:"title"`
	Content      string      `json:"contant"`
	Location     string      `json:"location"`
	Owner        string      `json:"owner"`
	Tags         []string    `json:"tags"`
	Restrictions []string    `json:"restrictions"`
	Costs        []Cost      `json:"costs"`
	Times        []Time      `json:"times"`
	TicketURLs   []TicketURL `json:"ticketUrls"`
}

type Cost struct {
	TicketType string  `json:"ticketType"`
	Amount     float32 `json:"amount"`
}

type Time struct {
	Doors string `json:"doors"`
	Start string `json:"start"`
	End   string `json:"end"`
}

type TicketURL struct {
	Provider string `json:"provider"`
	URL      string `json:"url"`
}

type Location struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Address     string `json:"address"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	URL         string `json:"url"`
}

func LoadEvents(file string) ([]Event, error) {
	var events []Event
	eventJson, err := os.ReadFile(file)
	if err != nil {
		return events, err
	}
	err = json.Unmarshal(eventJson, &events)
	if err != nil {
		return events, err
	}
	return events, nil
}

func LoadLocations(file string) ([]Location, error) {
	var locations []Location
	locationJson, err := os.ReadFile(file)
	if err != nil {
		return locations, err
	}
	err = json.Unmarshal(locationJson, &locations)
	if err != nil {
		return locations, err
	}
	return locations, nil
}
