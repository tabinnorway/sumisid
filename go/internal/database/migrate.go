package db

import (
	"fmt"
)

func (d *Database) MigrateDB() error {
	fmt.Println("Inital DB migration")

	driver, err = postgres.WithInstance(d.Client.DB, &postgres.Config{})
	return nil
}
