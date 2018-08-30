package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

type Product struct {
	gorm.Model
	Code string
	Name string
	Price uint
}

func main() {

	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed to connect database")
	}

	// Read
	var product Product
	db.First(&product, "Name = ?", "Apple") // find product with name
	fmt.Println(product)
}
