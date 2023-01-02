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

	var conn *gorm.DB

	if conf.DataBase.Engine == "sqlite" {
		conn = OpenSqlite(conf.DataBase.Name + ".db")
	}

	if conf.DataBase.Engine == "postgres" {
		conn = OpenPostgres(
			conf.DataBase.Host,
			conf.DataBase.Port,
			conf.DataBase.User,
			conf.DataBase.Name,
			conf.DataBase.TimeZone,
		)
	}

	d.db = conn

	return conn
}

func (d *DB) Migrate() {
	d.db.AutoMigrate(&models.UserModel{})
}
