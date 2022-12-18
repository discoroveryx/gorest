package dbstorage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenSqlite(dbname string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
