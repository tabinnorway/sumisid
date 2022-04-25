package handlers

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"math/rand"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
// 	"github.com/tabinnorway/sumisid/go/pkg/mocks"
// 	"github.com/tabinnorway/sumisid/go/pkg/models"
// )

// func GetAllClubs(w http.ResponseWriter, r *http.Request) {
// 	OK(w, mocks.Clubs)
// }

// func GetClub(w http.ResponseWriter, r *http.Request) {
// 	// Get the dive club Id requested
// 	vars := mux.Vars(r)
// 	id, _ := strconv.Atoi(vars["id"])

// 	// Find the dive club with that id
// 	for _, diveClub := range mocks.Clubs {
// 		if diveClub.Id == id {
// 			OK(w, diveClub)
// 			return
// 		}
// 	}
// 	NotFound(w)
// }

// func AddClub(w http.ResponseWriter, r *http.Request) {
// 	// Read dive club from request body
// 	defer r.Body.Close()
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	var diveClub models.Club
// 	json.Unmarshal(body, &diveClub)

// 	// Append dive club to list
// 	diveClub.Id = rand.Int()
// 	mocks.Clubs = append(mocks.Clubs, diveClub)
// 	Created(w, diveClub)
// }

// func UpdateClub(w http.ResponseWriter, r *http.Request) {
// 	// get the id
// 	vars := mux.Vars(r)
// 	id, _ := strconv.Atoi(vars["id"])

// 	// Read dive club from request body
// 	defer r.Body.Close()
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	var updatedClub models.Club
// 	json.Unmarshal(body, &updatedClub)

// 	for index, diveClub := range mocks.Clubs {
// 		if diveClub.Id == id {
// 			// If we find the dive club, update it
// 			diveClub.Name = updatedClub.Name
// 			diveClub.StreetAddress = updatedClub.StreetAddress
// 			diveClub.StreetNumber = updatedClub.StreetNumber
// 			diveClub.ZipCode = updatedClub.ZipCode
// 			diveClub.PhoneNumber = updatedClub.PhoneNumber
// 			diveClub.ContactPersonId = updatedClub.ContactPersonId
// 			diveClub.ExtraInfo = updatedClub.ExtraInfo
// 			mocks.Clubs[index] = diveClub

// 			OK(w, diveClub)
// 			return
// 		}
// 	}
// 	NotFound(w)
// }

// func DeleteClub(w http.ResponseWriter, r *http.Request) {
// 	// Get the dive club Id requested
// 	vars := mux.Vars(r)
// 	id, _ := strconv.Atoi(vars["id"])

// 	// Find the dive club with that id
// 	for index, diveClub := range mocks.Clubs {
// 		if diveClub.Id == id {
// 			mocks.Clubs = append(mocks.Clubs[:index], mocks.Clubs[index+1:]...)
// 			// If we find the dive club, delete it
// 			w.WriteHeader(http.StatusNoContent)
// 			return
// 		}
// 	}
// 	NotFound(w)
// }
