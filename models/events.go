package models

type Event struct {
	Id           int    `json: "id"`
	Name         string `json:"name"`
	Creator      User   `json: "creator"`
	Participants []User `json:"participants"`
	Items        []Item
}

var Events []Event
