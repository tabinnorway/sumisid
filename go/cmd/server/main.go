package main

import (
	"fmt"
	"log"
	"time"

	db "github.com/tabinnorway/sumisid/go/internal/database"
	"github.com/tabinnorway/sumisid/go/internal/services"
	diveclub "github.com/tabinnorway/sumisid/go/internal/services"
	transportHttp "github.com/tabinnorway/sumisid/go/internal/transport/http"
)

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
	personService := services.NewPersonService(db)

	fmt.Print("Application is starting...")
	now := time.Now()
	log.Println("Server started at: ", now.Local().Format(time.UnixDate))

	httpHandler := transportHttp.NewHandler(dcService, personService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Println(err)
	}
}
