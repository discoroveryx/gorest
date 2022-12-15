package dbstorage

import (
	"app/user/models"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	Db *gorm.DB
}

func connDB(dbname string) *gorm.DB {
	conn, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return conn
}

func NewDB() DB {
	fmt.Println("NewDB")
	return DB{Db: connDB("main.db")}
}

func NewTestDB() DB {
	fmt.Println("NewTestDB")
	return DB{Db: connDB("test_1.db")}
}

func (d *DB) Migrate() {
	d.Db.AutoMigrate(&models.UserModel{})
}
