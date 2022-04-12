package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func CreateRoutes() *mux.Router {
	router := mux.NewRouter()

	// router.HandleFunc("/api/people", handlers.GetAllPeople).Methods(http.MethodGet)
	// router.HandleFunc("/api/people/{id}", handlers.GetPerson).Methods(http.MethodGet)
	// router.HandleFunc("/api/people/{id}", handlers.DeletePerson).Methods(http.MethodDelete)
	// router.HandleFunc("/api/people/{id}", handlers.UpdatePerson).Methods(http.MethodPut)
	// router.HandleFunc("/api/people", handlers.AddPerson).Methods(http.MethodPost)

	// router.HandleFunc("/api/diveclubs", handlers.GetAllDiveClubs).Methods(http.MethodGet)
	// router.HandleFunc("/api/diveclubs/{id}", handlers.GetDiveClub).Methods(http.MethodGet)
	// router.HandleFunc("/api/diveclubs/{id}", handlers.UpdateDiveClub).Methods(http.MethodPut)
	// router.HandleFunc("/api/diveclubs/{id}", handlers.DeleteDiveClub).Methods(http.MethodDelete)
	// router.HandleFunc("/api/diveclubs", handlers.AddDiveClub).Methods(http.MethodPost)

	return router
}

func Run() error {
	port := ":8080"
	fmt.Print("Starting application...")
	fmt.Println("listening on port ", port)
	now := time.Now()
	log.Println("Server started at: ", now.Local().Format(time.UnixDate))

	router := CreateRoutes()
	http.ListenAndServe(port, router)
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Println(err)
	}
}
