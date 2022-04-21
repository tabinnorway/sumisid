package diveclub

import (
	"context"
	"errors"
)

var (
	ErrorNotImplemented = errors.New("Error method is not implemented")
	ErrorNotFound       = errors.New("Error item not found")
	ErrorNoAccess       = errors.New("Error user does not have access to item")
)

type DiveClub struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	StreetAddress   string `json:"streetAddress"`
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
	UpdateDiveClub(ctx context.Context, id int, dc DiveClub) (DiveClub, error)
	DeleteDiveClub(ctx context.Context, id int) error
	CreateDiveClub(ctx context.Context, dc DiveClub) (DiveClub, error)
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
	dc, err := service.DiveClubStore.GetDiveClub(ctx, id)
	if err != nil {
		return DiveClub{}, err
	}
	return dc, nil
}

func (service *DiveClubService) UpdateDiveClub(ctx context.Context, id int, dc DiveClub) (DiveClub, error) {
	return service.DiveClubStore.UpdateDiveClub(ctx, id, dc)
}

func (service *DiveClubService) DeleteDiveClub(ctx context.Context, id int) error {
	return service.DiveClubStore.DeleteDiveClub(ctx, id)
}

func (service *DiveClubService) CreateDiveClub(ctx context.Context, dc DiveClub) (DiveClub, error) {
	dc, err := service.DiveClubStore.CreateDiveClub(ctx, dc)
	if err != nil {
		return DiveClub{}, err
	}
	return dc, nil
}
