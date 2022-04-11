package mocks

import (
	"time"

	"github.com/tabinnorway/sumisid/go/pkg/models"
)

var People = []models.Person{
	{
		Id:          1,
		FirstName:   "Terje",
		MiddleName:  "Anthon",
		LastName:    "Bergesen",
		BirthDate:   time.Date(1966, time.October, 29, 0, 0, 0, 0, time.Local),
		IsAdmin:     true,
		PhoneNumber: "900 12 465",
		MainClubId:  1,
	},
	{
		Id:          1,
		FirstName:   "Andrea",
		MiddleName:  "Færøyvik",
		LastName:    "Bergesen",
		BirthDate:   time.Date(2011, time.April, 11, 0, 0, 0, 0, time.Local),
		IsAdmin:     true,
		PhoneNumber: "900 12 465",
		MainClubId:  1,
	},
	{
		Id:          1,
		FirstName:   "Paul",
		MiddleName:  "Joachim",
		LastName:    "Thorsen",
		BirthDate:   time.Date(0, time.January, 0, 0, 0, 0, 0, time.Local),
		IsAdmin:     true,
		PhoneNumber: "999 99 999",
		MainClubId:  1,
	},
	{
		Id:          1,
		FirstName:   "Stavanger",
		MiddleName:  "",
		LastName:    "Person",
		BirthDate:   time.Date(0, time.January, 0, 0, 0, 0, 0, time.Local),
		IsAdmin:     false,
		PhoneNumber: "",
		MainClubId:  2,
	},
}
