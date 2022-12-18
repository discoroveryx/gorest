package dbstorage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenSqlite(dbname string) *gorm.DB {
	// myconf := myconfig.GetMyConfig()
	// dbname := myconf.DbName

	db, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
