package dbstorage

import (
	"app/user/models"

	"myconfig"

	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func (d *DB) Connect() *gorm.DB {
	myconf := myconfig.GetMyConfig()
	dbname := myconf.DBName
	dbengine := myconf.DBEngine

	if dbengine != "sqlite" {
	}

	conn := OpenSqlite(dbname)
	return conn
}

func (d *DB) Migrate() {
	d.db.AutoMigrate(&models.UserModel{})
}
