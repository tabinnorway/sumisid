package diveclub

import (
	"context"
	"errors"
	"log"
)

var (
	ErrorNotImplemented = errors.New("Error method is not implemented")
	ErrorNotFound       = errors.New("Error item not found")
	ErrorNoAccess       = errors.New("Error user does not have access to item")
)

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

// DiveClubStore - this interface defines all methods that our service needs
// to manipulate the storage of dive clubs
type DiveClubStore interface {
	GetDiveClub(context.Context, int) (DiveClub, error)
	UpdateDiveClub(ctx context.Context, dc DiveClub) (bool, error)
	DeleteDiveClub(ctx context.Context, id int) (bool, error)
	CreateDiveClub(ctx context.Context, id int) (DiveClub, error)
}

type DiveClubService struct {
	DiveClubStore DiveClubStore
}

func NewDiveClubService(diveClubStore DiveClubStore) *DiveClubService {
	return &DiveClubService{
		DiveClubStore: diveClubStore,
	}
}

func (service *DiveClubService) GetDiveClub(ctx context.Context, id int) (DiveClub, error) {
	log.Println("Getting dive club")
	dc, err := service.GetDiveClub(ctx, id)
	if err != nil {
		log.Println(err)
		return DiveClub{}, ErrorNotFound
	}
	return dc, nil
}

func (service *DiveClubService) UpdateDiveClub(ctx context.Context, dc DiveClub) (bool, error) {
	return false, ErrorNotImplemented
}

func (service *DiveClubService) DeleteDiveClub(ctx context.Context, id int) (bool, error) {
	return false, ErrorNotImplemented
}

func (service *DiveClubService) CreateDiveClub(ctx context.Context, id int) (DiveClub, error) {
	return DiveClub{}, ErrorNotImplemented
}
