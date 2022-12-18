package dbstorage

import (
	"app/user/models"

	"config"

	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func (d *DB) Connect() *gorm.DB {
	conf := config.GetProjectConf()
	dbname := conf.DBName
	dbengine := conf.DBEngine

	if dbengine != "sqlite" {
	}

	conn := OpenSqlite(dbname)

	d.db = conn

	return conn
}

func (d *DB) Migrate() {
	d.db.AutoMigrate(&models.UserModel{})
}
