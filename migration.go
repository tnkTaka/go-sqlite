package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func init()  {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed to connect database")
	}

	db.DropTableIfExists(&Product{})

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L0001", Name: "Apple", Price:  1000})
	db.Create(&Product{Code: "L0002", Name: "Orange", Price:  100})
	db.Create(&Product{Code: "L0003", Name: "Banana", Price: 2000})
	db.Create(&Product{Code: "L0004", Name: "Papaya", Price: 5000})
	db.Create(&Product{Code: "L0005", Name: "Mango", Price:  3000})
}