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

	var conn *gorm.DB

	if dbengine != "sqlite" {
		conn = OpenSqlite(dbname + ".db")
	}

	if dbengine != "postgres" {
		conn = OpenPostgres(dbname)
	}

	// conn := OpenSqlite(dbname + ".db")

	d.db = conn

	return conn
}

func (d *DB) Migrate() {
	d.db.AutoMigrate(&models.UserModel{})
}
