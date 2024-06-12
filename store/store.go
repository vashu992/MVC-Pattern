package store

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func (store *Postgres) NewStore() error {
	dsn := "host=localhost user=ashutosh password=password dbname=pmvpattern port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	} else {
		store.DB = db
	}
	fmt.Printf("server = %v\n", db)
	return nil

}

type StoreOperations interface {
	NewStore() error
}
