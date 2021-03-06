package main

import (
	"fmt"
	"log"
	"time"

	db "github.com/tabinnorway/sumisid/go/internal/database"
	"github.com/tabinnorway/sumisid/go/internal/services"
	club "github.com/tabinnorway/sumisid/go/internal/services"
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

	dcService := club.NewClubService(db)
	personService := services.NewPersonService(db)

	now := time.Now()
	log.Println(fmt.Sprintf("Server started at: %s", now.Local().Format(time.UnixDate)))

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
