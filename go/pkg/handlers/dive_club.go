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

func GetAllDiveClubs(w http.ResponseWriter, r *http.Request) {
	OK(w, mocks.DiveClubs)
}

func GetDiveClub(w http.ResponseWriter, r *http.Request) {
	// Get the dive club Id requested
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find the dive club with that id
	for _, diveClub := range mocks.DiveClubs {
		if diveClub.Id == id {
			OK(w, diveClub)
			return
		}
	}
	NotFound(w)
}

func AddDiveClub(w http.ResponseWriter, r *http.Request) {
	// Read dive club from request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var diveClub models.DiveClub
	json.Unmarshal(body, &diveClub)

	// Append dive club to list
	diveClub.Id = rand.Int()
	mocks.DiveClubs = append(mocks.DiveClubs, diveClub)
	Created(w, diveClub)
}

func UpdateDiveClub(w http.ResponseWriter, r *http.Request) {
	// get the id
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read dive club from request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var updatedDiveClub models.DiveClub
	json.Unmarshal(body, &updatedDiveClub)

	for index, diveClub := range mocks.DiveClubs {
		if diveClub.Id == id {
			// If we find the dive club, update it
			diveClub.Name = updatedDiveClub.Name
			diveClub.StreetAddress = updatedDiveClub.StreetAddress
			diveClub.StreetNumber = updatedDiveClub.StreetNumber
			diveClub.ZipCode = updatedDiveClub.ZipCode
			diveClub.PhoneNumber = updatedDiveClub.PhoneNumber
			diveClub.ContactPersonId = updatedDiveClub.ContactPersonId
			diveClub.ExtraInfo = updatedDiveClub.ExtraInfo
			mocks.DiveClubs[index] = diveClub

			OK(w, diveClub)
			return
		}
	}
	NotFound(w)
}

func DeleteDiveClub(w http.ResponseWriter, r *http.Request) {
	// Get the dive club Id requested
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find the dive club with that id
	for index, diveClub := range mocks.DiveClubs {
		if diveClub.Id == id {
			mocks.DiveClubs = append(mocks.DiveClubs[:index], mocks.DiveClubs[index+1:]...)
			// If we find the dive club, delete it
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	NotFound(w)
}
