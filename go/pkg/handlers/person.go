package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tabinnorway/sumisid/go/pkg/mocks"
	"github.com/tabinnorway/sumisid/go/pkg/models"
)

func GetAllPeople(w http.ResponseWriter, r *http.Request) {
	OK(w, mocks.People)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	// Get the club Id requested
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	// Find the club with that id
	for _, person := range mocks.People {
		if person.Id == id {
			OK(w, person)
			return
		}
	}
	NotFound(w)
}

func AddPerson(w http.ResponseWriter, r *http.Request) {
	// Read club from request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var person models.Person
	json.Unmarshal(body, &person)

	// Append club to list
	person.Id = rand.Int()
	mocks.People = append(mocks.People, person)
	Created(w, person)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	// Get the club Id requested
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find the club with that id
	for index, person := range mocks.People {
		if person.Id == id {
			mocks.People = append(mocks.People[:index], mocks.People[index+1:]...)
			// If we find the club, delete it
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	NotFound(w)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	// get the id
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read person from request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var updatedPerson models.Person
	json.Unmarshal(body, &updatedPerson)

	for index, person := range mocks.People {
		if person.Id == id {
			// If we find the person, update it
			person.FirstName = updatedPerson.FirstName
			person.MiddleName = updatedPerson.MiddleName
			person.LastName = updatedPerson.LastName
			person.BirthDate = updatedPerson.BirthDate
			person.IsAdmin = updatedPerson.IsAdmin
			person.PhoneNumber = updatedPerson.PhoneNumber
			person.MainClubId = updatedPerson.MainClubId
			mocks.People[index] = person

			OK(w, person)
			return
		}
	}
	NotFound(w)
}
