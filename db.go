package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	FirstName     string
	MiddleName    *string
	LastName      string
	DoB           string
	Gender        string
	Email         *string
	Address       string
	GuardianName  string
	GuardianPhone string
	GuardianEmail *string
	Grade         string
}

func db() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Student{})

	return db, nil
}
