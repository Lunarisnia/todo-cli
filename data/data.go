package data

import (
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/lunarisnia/todo-cli/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Todo struct {
	gorm.Model
	Title       string
	Description string
	Status      bool
}

func OpenDatabase() error {
	config.CheckDatabaseFile()
	var err error

	db, err = gorm.Open(sqlite.Open(config.DBPATH), &gorm.Config{})
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

func InsertTodo(title string, description string, status bool) {
	todo := Todo{Title: title, Description: description, Status: status}
	if err := db.Create(&todo).Error; err != nil {
		remindInit()
	}
}

func ReadAllTodos(everything bool) []Todo {
	todos := []Todo{}

	if everything {
		if err := db.Find(&todos).Error; err != nil {
			remindInit()
		}
		return todos
	}

	if err := db.Where("status = 0").Find(&todos).Error; err != nil {
		remindInit()
	}
	return todos
}

func FindOneTodo(todoID string) Todo {
	todo := Todo{}
	if err := db.Where("id = ?", todoID).First(&todo).Error; err != nil {
		remindInit()
	}
	return todo
}

func MarkTodoAsDone(todo *Todo) {
	if err := db.Model(todo).Update("status", true).Error; err != nil {
		remindInit()
	}
}

func MarkTodoAsNotDone(todo *Todo) {
	if err := db.Model(todo).Update("status", false).Error; err != nil {
		remindInit()
	}
}

func DeleteTodo(todo *Todo) {
	if err := db.Delete(todo).Error; err != nil {
		remindInit()
	}
}

func remindInit() {
	log.Fatalln(`
==========================
Have you ran "todo init"?
==========================
	`)
}
