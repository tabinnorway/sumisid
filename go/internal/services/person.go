package services

import (
	"context"
	"time"
)

type Person struct {
	Id          int       `json:"id"`
	Email       string    `json:"email"`
	FirstName   string    `json:"firstName"`
	MiddleName  string    `json:"middleName"`
	LastName    string    `json:"lastName"`
	BirthDate   time.Time `json:"birthDate"`
	IsAdmin     bool      `json:"isAdmin"`
	PhoneNumber string    `json:"phoneNumber"`
	MainClubId  int       `json:"mainClubId"`
	MainClub    DiveClub  `json:"mainClub"`
}

// PersonStore - this interface defines all methods that our service needs
// to manipulate the storage of people
type PersonStore interface {
	GetAllPerson(context.Context) ([]Person, error)
	GetPerson(context.Context, int) (Person, error)
	UpdatePerson(ctx context.Context, id int, p Person) (Person, error)
	DeletePerson(ctx context.Context, id int) error
	CreatePerson(ctx context.Context, p Person) (Person, error)
}

type PersonService struct {
	PersonStore PersonStore
}

func NewPersonService(personStore PersonStore) *PersonService {
	return &PersonService{
		PersonStore: personStore,
	}
}

func (service *PersonService) GetAllPerson(ctx context.Context) ([]Person, error) {
	ps, err := service.PersonStore.GetAllPerson(ctx)
	if err != nil {
		return []Person{}, err
	}
	return ps, nil
}

func (service *PersonService) GetPerson(ctx context.Context, id int) (Person, error) {
	dc, err := service.PersonStore.GetPerson(ctx, id)
	if err != nil {
		return Person{}, err
	}
	return dc, nil
}

func (service *PersonService) UpdatePerson(ctx context.Context, id int, p Person) (Person, error) {
	return service.PersonStore.UpdatePerson(ctx, id, p)
}

func (service *PersonService) DeletePerson(ctx context.Context, id int) error {
	return service.PersonStore.DeletePerson(ctx, id)
}

func (service *PersonService) CreatePerson(ctx context.Context, p Person) (Person, error) {
	p, err := service.PersonStore.CreatePerson(ctx, p)
	if err != nil {
		return Person{}, err
	}
	return p, nil
}
