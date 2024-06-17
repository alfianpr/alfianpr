package database

import (
	"alfianpr/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func Connect(connectionString string) error {
	var err error
	DB, err = gorm.Open("sqlite3", connectionString)
	if err != nil {
		return err
	}
	models.Migrate(DB)
	return nil
}
