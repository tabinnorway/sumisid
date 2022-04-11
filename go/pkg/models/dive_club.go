package models

type DiveClub struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	StreetAddress   string `json:"addressId"`
	StreetNumber    string `json:"streetNumber"`
	ZipCode         string `json:"zipCode"`
	PhoneNumber     string `json:"phoneNumber"`
	ContactPersonId int    `json:"contactPersonId"`
	ExtraInfo       string `json:"extraInfo"`
}
