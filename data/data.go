package data

import (
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type Todo struct {
	gorm.Model
	Title  string
	Status bool
}

func OpenDatabase() error {
	var err error

	db, err = gorm.Open(sqlite.Open("todo-db.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func MigrateDatabase() {
	err := db.AutoMigrate(&Todo{})
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	log.Println("Database initiated")
}