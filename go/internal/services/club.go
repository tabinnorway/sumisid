package services

import (
	"context"
	"errors"
)

var (
	ErrorNotImplemented = errors.New("Error method is not implemented")
	ErrorNotFound       = errors.New("Error item not found")
	ErrorNoAccess       = errors.New("Error user does not have access to item")
)

type Club struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	StreetAddress   string `json:"streetAddress"`
	StreetNumber    string `json:"streetNumber"`
	ZipCode         string `json:"zipCode"`
	PhoneNumber     string `json:"phoneNumber"`
	ContactPersonId int    `json:"contactPersonId"`
	ExtraInfo       string `json:"extraInfo"`
}

// ClubStore - this interface defines all methods that our service needs
// to manipulate the storage of clubs
type ClubStore interface {
	GetAllClub(context.Context) ([]Club, error)
	GetClub(context.Context, int) (Club, error)
	UpdateClub(ctx context.Context, id int, dc Club) (Club, error)
	DeleteClub(ctx context.Context, id int) error
	CreateClub(ctx context.Context, dc Club) (Club, error)
}

type ClubService struct {
	ClubStore ClubStore
}

func NewClubService(clubStore ClubStore) *ClubService {
	return &ClubService{
		ClubStore: clubStore,
	}
}

func (service *ClubService) GetAllClub(ctx context.Context) ([]Club, error) {
	dcs, err := service.ClubStore.GetAllClub(ctx)
	if err != nil {
		return []Club{}, err
	}
	return dcs, nil
}

func (service *ClubService) GetClub(ctx context.Context, id int) (Club, error) {
	dc, err := service.ClubStore.GetClub(ctx, id)
	if err != nil {
		return Club{}, err
	}
	return dc, nil
}

func (service *ClubService) UpdateClub(ctx context.Context, id int, dc Club) (Club, error) {
	return service.ClubStore.UpdateClub(ctx, id, dc)
}

func (service *ClubService) DeleteClub(ctx context.Context, id int) error {
	return service.ClubStore.DeleteClub(ctx, id)
}

func (service *ClubService) CreateClub(ctx context.Context, dc Club) (Club, error) {
	dc, err := service.ClubStore.CreateClub(ctx, dc)
	if err != nil {
		return Club{}, err
	}
	return dc, nil
}
