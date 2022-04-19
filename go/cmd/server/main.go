package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/mux"
	db "github.com/tabinnorway/sumisid/go/internal/database"
	"github.com/tabinnorway/sumisid/go/internal/diveclub"
	transportHttp "github.com/tabinnorway/sumisid/go/internal/transport/http"
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
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database: ", err)
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to migrate database")
		return err
	}

	dcService := diveclub.NewDiveClubService(db)

	port := ":8080"
	fmt.Print("Application is starting...")
	fmt.Println("listening on port ", port)
	now := time.Now()
	log.Println("Server started at: ", now.Local().Format(time.UnixDate))

	httpHandler := transportHttp.NewHandler(dcService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	// router := CreateRoutes()
	// http.ListenAndServe(port, router)
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Println(err)
	}
}
