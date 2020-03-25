package database

import "github.com/jinzhu/gorm"

type Database struct {
	Db *gorm.DB
}

func NewDatabase() *Database {
	db, _ := gorm.Open("sqlite3", "database.db")
	return &Database{
		Db: db,
	}
}
