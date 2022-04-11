package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/tabinnorway/sumisid/go/pkg/handlers"
)

func CreateRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/people", handlers.GetAllPeople).Methods(http.MethodGet)
	router.HandleFunc("/api/people/{id}", handlers.GetPerson).Methods(http.MethodGet)
	router.HandleFunc("/api/people/{id}", handlers.DeletePerson).Methods(http.MethodDelete)
	router.HandleFunc("/api/people/{id}", handlers.UpdatePerson).Methods(http.MethodPut)
	router.HandleFunc("/api/people", handlers.AddPerson).Methods(http.MethodPost)

	router.HandleFunc("/api/diveclubs", handlers.GetAllDiveClubs).Methods(http.MethodGet)
	router.HandleFunc("/api/diveclubs/{id}", handlers.GetDiveClub).Methods(http.MethodGet)
	router.HandleFunc("/api/diveclubs/{id}", handlers.UpdateDiveClub).Methods(http.MethodPut)
	router.HandleFunc("/api/diveclubs/{id}", handlers.DeleteDiveClub).Methods(http.MethodDelete)
	router.HandleFunc("/api/diveclubs", handlers.AddDiveClub).Methods(http.MethodPost)

	return router
}

func main() {
	router := CreateRoutes()
	now := time.Now()

	log.Println("Server started at: ", now.Local().Format(time.UnixDate))
	log.Println("Listening on port 4000")
	http.ListenAndServe(":4000", router)
}
