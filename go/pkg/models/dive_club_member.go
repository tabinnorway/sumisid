package models

type DiveClubMember struct {
	Id       int `json:"id"`
	PersonId int `json:"personId"`
	ClubId   int `json:"diveClubId"`
}
