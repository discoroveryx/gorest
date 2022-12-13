package dbstorage

import (
	"app/user/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	Db *gorm.DB
}

func NewDB() DB {
	conn, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return DB{Db: conn}
}

func (d *DB) Migrate() {
	d.Db.AutoMigrate(&models.UserModel{})
}
