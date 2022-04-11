package models

import "time"

type Person struct {
	Id          int       `json:"id"`
	FirstName   string    `json:"firstName"`
	MiddleName  string    `json:"middleName"`
	LastName    string    `json:"lastName"`
	BirthDate   time.Time `json:"birthDate"`
	IsAdmin     bool      `json:"isAdmin"`
	PhoneNumber string    `json:"phoneNumber"`
	MainClubId  int       `json:"diveClubId"`
}
