package database

import (
	"alfianpr/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(connectionString string) error {
	var err error
	DB, err = gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}
	models.Migrate(DB)
	return nil
}
