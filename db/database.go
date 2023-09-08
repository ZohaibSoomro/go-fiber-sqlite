package db

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"github.com/zohaibsoomro/go-fiber-sqlite/models"
	"gorm.io/gorm"
)

const dbName = "employee.db"

var Db *gorm.DB

// connect initializes the connection to db
func Connect() error {

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error connecting to db %s", err)
	}
	Db = db
	Db.AutoMigrate(&models.Employee{})
	return nil
}

func GetDb() *gorm.DB {
	return Db
}
