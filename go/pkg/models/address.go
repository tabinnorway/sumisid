package models

type Address struct {
	Id           int    `json:"id"`
	Street       string `json:"street"`
	StreetNumber string `json:"streetNumber"`
	PostalCode   string `json:"postalCode"`
	City         string `json:"city"`
	CountryId    int    `json:"contryId"`
}
