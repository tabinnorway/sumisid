package models

type ClubMember struct {
	Id       int `json:"id"`
	PersonId int `json:"personId"`
	ClubId   int `json:"clubId"`
}
